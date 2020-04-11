package main

import (
	"fmt"

	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/controller"
	"github.com/AkashTyagi-SD/Webservicesgolang/github.com/models"
)

func main() {
	fmt.Println("Hello custome workspace")
	fmt.Println(models.Morning)
	mul := controller.Multiple(2, 3)
	fmt.Println("mul", mul)
}
