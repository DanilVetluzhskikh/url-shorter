package main

import (
	"fmt"
	"learn/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
