package main

import (
	"fmt"

	"github.com/wagaodev/Go-Opportunities/config"
	"github.com/wagaodev/Go-Opportunities/router"
)

func main() {
	err := config.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	router.Initialize()
}
