package main

import (
	"fmt"
	ellipse "github.com/sasonogumelar/arith"
)

func main() {
	line := ellipse.Init{2, 3}
	fmt.Println(line.Add())
}
