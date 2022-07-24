package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < 16; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(n int) (res string) {
	f := newFB(n)
	res = f.val
	return
}

type fb struct {
	number int
	val    string
}

var fbMap = map[int]string{
	15: "FizzBuzz",
	5:  "Buzz",
	3:  "Fizz",
}

func newFB(num int) (resFB *fb) {
	resFB = &fb{number: num}
	for k, v := range fbMap {
		if num%k == 0 {
			resFB.val = v
			return
		}
	}
	resFB.val = strconv.Itoa(num)
	return
}
