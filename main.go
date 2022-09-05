package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

var afstand int
var uren int
var kmph int

func init() {
	flag.IntVar(&afstand, "a", 150000000, "afstand")
	flag.IntVar(&uren, "u", 8766, "uren")
	flag.Parse()
}

func main() {
	if afstand < 1 || uren < 1 {
		fmt.Println("afstand en uren mogen niet kleiner dan 1 zijn")
		os.Exit(1)
	}
	diameter := afstand * 2
	pi := math.Pi
	omtrek := float64(diameter) * pi
	kmph := omtrek / float64(uren)
	fmt.Println(int(kmph), "Kilometer per uur")

}
