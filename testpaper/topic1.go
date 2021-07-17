package main

import "fmt"

type Compare interface {
	max(nums []int) int
	maxTwo(num1, num2 int) int
	maxThree(num1, num2, num3 int) int
	maxFive(num1, num2, num3, num4, num5 int) int
}

type Comparer struct{}

func (c *Comparer) max(nums []int) int {
	return 0
}

func (c *Comparer) maxTwo(num1, num2 int) int {

	if num1 >= num2 {
		return num1
	} else {
		return num2
	}

}

func (c *Comparer) maxThree(num1, num2, num3 int) int {
	if num1 > num2 || num1 == num2 {
		if num1 > num3 || num1 == num3 {
			return num1
		} else {
			return num3
		}
	} else {
		if num2 > num3 || num2 == num3 {
			return num2
		} else {
			return num3
		}
	}
}

func main() {
	c1 := Comparer{}
	max1 := c1.maxTwo(1, 2)
	fmt.Println(max1)

	maxThree := c1.maxThree(3, 3, 3)
	fmt.Println(maxThree)

}
