package main

import (
	"fmt"

	"github.com/Junx27/raiden/feature"
	"github.com/Junx27/raiden/hello"
)

func main() {
	hello.SayHello()
	res := feature.Add(1, 2)
	fmt.Println(res)
}
