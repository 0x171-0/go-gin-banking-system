package main

import (
	"go-gin-template/api"
)

func main() {
	r := api.InitRouter()
	r.Run(":3002")
}
