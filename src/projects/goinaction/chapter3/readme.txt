1. godoc fmt 看包文档
2. go get 导入URL指定的包
3. 重名包通过命名导入来导入
import (
    "fmt"
    myfmt "mylib/fmt"
)
4. 导入未使用的包将引起编译错误
5. 需要导入，但不使用包时，使用空白标识符_ 来重命名这个导入
6. 在包中可以包含任意多个init 函数，编译器在main函数运行之前运行init函数
    init用于初始化变量
go build hello.go
go clean hello.go
go vet main.go
go fmt 格式化代码