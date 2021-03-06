package main

import (
	"GoDesignPattern/Behavioral/iterator"
	"fmt"
	"strconv"
)

func main() {
	container := iterator.NewContainer()
	container.Append(1)
	container.Append(2)
	container.Append(3)
	container.Append(4)
	container.Append(5)
	container.Append(6)

	for container.HasNext(){
		obj := container.Next()
		fmt.Println(strconv.Itoa(obj))
	}
}





