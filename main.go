package main

import (
	"fmt"
	"github.com/chrisdcarpenter/bookTracking/internal/author"
)

func main() {
	a := author.Author{
		Name: "Tao Wong",
	}
	fmt.Printf("Hello %s", a.Name)
}
