-------------------------- chapter04 数组

var array [5]int  // 长度5初始化为0
array := [5]int{10, 20, 30, 40, 50} // 使用字面量初始化数组
array := [...]int{10, 20, 30, 40, 50} // 自动计算数组长度
array := [5]int{1: 10, 2: 20} // 指定特定元素值

指针数组
array := [5]*int{0: new(int), 1: new(int)}
指针数组 赋值
*array[0] = 10
*array[1] = 20

// 数组值复制
var array1 [5]string
array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
array1 = array2 // 值复制 值有两份 互不影响

// 数组指针复制
var array1 [3]*string
array2 := [3]*string{new(string), new(string), new(string)}
*array2[0] = "Red"
*array2[1] = "Blue"
*array2[2] = "Green"
array1 = array2  // 指针复制 值一份 相互影响

// 多维数组
var array [4][2]int
array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
array := [4][2]int{1: {20, 21}, 3: {40, 41}}
array := [4][2]int{1: {0: 20}, 3: {1: 41}}
array[0][1] = 20
array[1][0] = 30
array[1][1] = 40

// 函数传递数组
函数间值传递变量， 拷贝大数组性能差，最好使用指针传递
var array [1e6]int
foo(&array)
func foo(array *[1e6]int) {}

-------------------------- chapter04 切片

切片
slice := make([]string, 5) // 长度和容量都是5的字符串切片
slice := make([]int, 3, 5) // 长度3 容量5 int切片
slice := []string{"Red", "Blue", "Green", "Yellow", "Pink"} // 长度和容量都是5的字符串切片
slice := []int{10, 20, 30} // 长度和容量都是3的int切片
slice := []string{99: ""} // 100个空字符串元素切片

nil切片
var slice []int // nil 整型切片，不做初始化。场景：返回切片，但发生异常时返回nil切片

空切片
slice := make([]int, 0) // 使用make创建空的整型切片
slice := []int{} // 使用字面量创建空的整型切片

计算切片长度和容量
对于底层数组容量是k的切片slice[i:j]来说：
长度：j - i
容量：k - i

基于同一个底层数组的不同切片，如果修改其中一个，同时会影响到下一个。
