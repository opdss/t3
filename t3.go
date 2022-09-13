package t3

import "fmt"

var T31 int

func init() {
	fmt.Println("T3 init()")
	T31 = 10
}

func SetT3(val int) int {
	fmt.Println("T3 SetT3(): t3 => %d", val)
	T31 = val
	return T31
}

func SetT3P100() int {
	fmt.Println("T3 SetT3(): t3 => %d")
	T31 += 100
	return T31
}

func GetT3() int {
	fmt.Println("T3 GetT3(): t3 => %d", T31)
	return T31
}
