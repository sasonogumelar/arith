package main

import (
	"fmt"
	ellipse "test1"
)

func main() {
	line := ellipse.Init{2, 3}
	fmt.Println(line.Add())
}
