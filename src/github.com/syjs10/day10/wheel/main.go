package main

import "fmt"

// Point 点
type Point struct {
	X, Y int
}

// Circle 圆
type Circle struct {
	Center Point
	Radius int
}

// Wheel 轮子
type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
	w = Wheel{Circle{Point{8, 8}, 5}, 20}
	w = Wheel{
		Circle: Circle{
			Center: Point{
				X: 8,
				Y: 8,
			},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v \n", w)
}
