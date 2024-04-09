package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"sort"
)

/**
在大数据处理场景中，我们经常需要对大文件进行排序操作。由于大文件无法一次性读入内存，因此我们需要将大文件切分为多个小文件，然后对每个小文件进行排序操作，最后将排序后的小文件合并为一个有序的大文件。

整个流程可以分为以下几个步骤：
步骤一：切分	将大文件切分为多个小文件，以便可以一次性读入内存并进行排序操作。
步骤二：排序	对每个小文件进行排序操作，可以使用快速排序或外部排序算法。
步骤三：合并	将排序后的小文件按照规定的顺序合并为一个有序的大文件。
*/

/*
*
切分大文件
@param inputFile 输入文件
@param outputFilePattern 输出文件模式
@param chunkSize 切分大小
*/
func splitFile(inputFile string, outputFilePattern string, chunkSize int) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := stat.Size()
	chunkNum := int(fileSize) / chunkSize

	buf := make([]byte, chunkSize)
	for i := 0; i < chunkNum; i++ {
		outputFile := fmt.Sprintf(outputFilePattern, i+1)
		outFile, err := os.Create(outputFile)
		if err != nil {
			return err
		}

		_, err = file.Read(buf)
		if err != nil {
			return err
		}

		_, err = outFile.Write(buf)
		if err != nil {
			return err
		}

		outFile.Close()
	}

	return nil
}

/*
*
快速排序算法
@param inputFile 输入文件
@param outputFile 输出文件
@param bufferSize 缓冲区大小
*/
func quickSortFile(inputFile string, outputFile string, bufferSize int) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := stat.Size()

	buf := make([]byte, bufferSize)
	recordSize := binary.Size(int64(0))

	records := make([]int64, 0)
	for i := 0; i < int(fileSize)/bufferSize; i++ {
		_, err := file.Read(buf)
		if err != nil {
			return err
		}

		for j := 0; j < bufferSize/recordSize; j++ {
			record := binary.LittleEndian.Uint64(buf[j*recordSize : (j+1)*recordSize])
			records = append(records, int64(record))
		}
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i] < records[j]
	})

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for _, record := range records {
		binary.Write(outFile, binary.LittleEndian, record)
	}

	return nil
}
