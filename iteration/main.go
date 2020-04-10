package main

import "fmt"

//Function to repeat string character N times
func Repeat(a string, n int) string {
	var r string
	if n >= 0 {
		for i := 0; i < n; i++ {
			r += a
		}
	} else {
		r = "N can not be less than 1"
	}
	return r
}
func main() {
	a := "a"
	fmt.Println(Repeat(a, 2))
}
