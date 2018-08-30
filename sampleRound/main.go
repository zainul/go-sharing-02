package sampleRound

import (
	"fmt"
)

func main() {
	var f float64
	f = 12.3
	fmt.Println(int64(f)) // 12
	f = 12.6
	fmt.Println(int64(f)) // 12

	f = 12.3
	fmt.Println(int64(f + 0.5)) // 12
	f = 12.6
	fmt.Println(int64(f + 0.5)) // 13

	f = 12.31
	fmt.Println(float64(int64(f*10+0.5)) / 10) // 12.3
	f = 12.66
	fmt.Println(float64(int64(f*10+0.5)) / 10) // 12.7

	f = 12.31
	fmt.Println(float64(int64(f*20+0.5)) / 20) // 12.3
	f = 12.66
	fmt.Println(float64(int64(f*20+0.5)) / 20) // 12.65

	fmt.Println(Round(0.363636, 0.05)) // 0.35
	fmt.Println(Round(3.232, 0.05))    // 3.25
	fmt.Println(Round(0.4888, 0.05))   // 0.5

	fmt.Println(Round(-0.363636, 0.05)) // -0.3
	fmt.Println(Round(-3.232, 0.05))    // -3.2
	fmt.Println(Round(-0.4888, 0.05))   // -0.45

	fmt.Println(Round2(-0.363636, 0.05)) // -0.35
	fmt.Println(Round2(-3.232, 0.05))    // -3.25
	fmt.Println(Round2(-0.4888, 0.05))   // -0.5
}

func Round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}

func Round2(x, unit float64) float64 {
	if x > 0 {
		return float64(int64(x/unit+0.5)) * unit
	}
	return float64(int64(x/unit-0.5)) * unit
}
