# 场景题
## 对两个slice进行append操作，以下内容是否可以编译通过
```go
func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := []int{5, 6, 7, 8}
	s1 = append(s1, s2)
	//正确用法
	//s1 = append(s1, s2...)
	fmt.Println(s1)
}
```

## 结构体对比，以下结构体是否可以编译通过
```go
func main() {
    sn1 = struct {
        name string
        age  int
    }{name: "test", age: 18}
    sn2 = struct {
        name string
        age  int
    }{name: "test", age: 18}
    if sn1 == sn2 {
        fmt.Println("sn1 == sn2")
    }
	
    //匿名结构体
    sm1 := struct {
        age  int
        name map[string]string
    }{age: 1, name: map[string]string{"a": "1"}}
    //匿名结构体
    sm2 := struct {
        age  int
        name map[string]string
    }{age: 1, name: map[string]string{"a": "1"}}
    
    //对比结构体,报错,因为只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
    //但是结构体属性中有不可以比较的类型，如map,slice。 如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。
    //这里就属于不可比较的类型,map
    if sm1 == sm2 {
        fmt.Println("sm1 == sm2")
    }
    /**
      输出结果
      invalid operation: sm1 == sm2 (struct containing map[string]string cannot be compared)
    */
    //如果非要比较,可以使用reflect.DeepEqual()来进行比较,该方式可以对比map,slice类型,还有属性顺序不一样的,不过属性顺序不一样的就会不相等.
    if reflect.DeepEqual(sm1, sm2) {
        fmt.Println("sm1 == sm2")
    } else {
        fmt.Println("sm1 != sm2")
    }
    /**
      输出结果
      sm1 == sm2
    */
    
    //下面这个就和sn1和sn2的属性顺序不同,sn3与sn1就不是相同的结构体了，不能直接用==进行比较。只能用reflect.DeepEqual反射来比较,但是结果还是sn1 !=sn3
    sn3 := struct {
        name string
        age  int
    }{age: 11, name: "qq"}
    if reflect.DeepEqual(sn1, sn3) {
        fmt.Println("sn1 ==sn3")
    } else {
        fmt.Println("sn1 !=sn3")
    }
    /**
      输出结果
      sn1 !=sn3
    */
}

```