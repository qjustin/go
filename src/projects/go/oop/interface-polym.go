package main

import "fmt"

/**
当一个类型定义了接口中的所有方法，我们称它实现了该接口
*/

// 1. 定义一个接口
type Phone interface {
	call()
}

// 2. 实现一个接口
type Nokia struct {
	name string
}

func (nokia *Nokia) call() {
	fmt.Println(nokia.name)
	nokia.name = "tony"
}

// 3. 多态的例子
// 3.1 定义一个接口
type Light interface {
	discount() int
	lightInfo()
}

// 定义两个类型
type WallLight struct {
	brand   string
	weight int
	price  int
}

type FlashLight struct {
	brand   string
	weight int
	price  int
}

// 两个类型分别实现接口
func (walllight *WallLight) discount() int {
	walllight.price -= 100
	return walllight.price
}

func (walllight *WallLight) lightInfo() {
	fmt.Println(walllight.brand)
	fmt.Println(walllight.weight)
	fmt.Println(walllight.price)
}

func (flashlight *FlashLight) discount() int {
	flashlight.price -= 200
	return flashlight.price
}

func (flashlight *FlashLight) lightInfo() {
	fmt.Println(flashlight.brand)
	fmt.Println(flashlight.weight)
	fmt.Println(flashlight.price)
}

func calTotalPrice(lights []Light) int {
	totalPrice := 0

	for _, light := range lights {
		light.lightInfo()
		totalPrice += light.discount()
	}

	return totalPrice
}
func main() {
	nokia := Nokia{name: "sam"}
	nokia.call()
	fmt.Println(nokia.name)
	walllight := WallLight{weight: 100, price: 500, brand: "WallLight-AE"}
	flashlight := FlashLight{weight: 200, price: 300, brand: "FlashLight-CCC"}

	goods := []Light{&walllight, &flashlight}
	totalPrice := calTotalPrice(goods)
	fmt.Println(totalPrice)
}
