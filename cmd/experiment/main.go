package main

import (
	"fmt"

	"github.com/mr-panta/wire-experiment/internal/factory"
)

func main() {
	manager, err := factory.NewDefaultManager(1, 2)
	fmt.Println(err)
	fmt.Println(manager)
}
