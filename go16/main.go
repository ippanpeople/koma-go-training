// 接口實現多態性
package main

import (
	"fmt"
	"math"
	"reflect"
)

// 定義一個幾何接口
type geometry interface {
	area() float64
	perim() float64
}

// 定義一個矩形結構體
type rect struct {
	w, h float64
}

// 定義一個圓形結構體
type circle struct {
	radius float64
}

// 矩形計算面積
func (r rect) area() float64 {
	return r.w * r.h
}

// 矩形計算周長
func (r rect) perim() float64 {
	return 2 * (r.w + r.h)
}

// 圓形計算面積
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 圓形計算周長
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 計算函數, 參數為幾何接口類型
func measure(g geometry) {
	fmt.Println(reflect.TypeOf(g), g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	//	聲明一個矩形
	r := rect{w: 4, h: 5}
	//	聲明一個圓形
	c := circle{radius: 10}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	fmt.Println("area:", c.area())
	fmt.Println("perim:", c.perim())

	measure(r)
	measure(c)

}
