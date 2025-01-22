package main

import (
	"fmt"

	"example.com/m/config"
	//"golang.org/x/tools/go/cfg"
	//"github.com/sattorovshoxrux3009/rest-api/config"
)

func main(){
	cfg := config.Load(".")
	fmt.Println(cfg)
}