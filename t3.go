package t3

import "fmt"

var T3 int

func init() {
	fmt.Println("T3 init()")
	T3 = 10
}

func SetT3(val int) int {
	fmt.Println("T3 SetT3(): t3 => %d", val)
	T3 = val
	return T3
}

func SetT3P100() int {
	fmt.Println("T3 SetT3(): t3 => %d")
	T3 += 100
	return T3
}

func GetT3() int {
	fmt.Println("T3 GetT3(): t3 => %d", T3)
	return T3
}
