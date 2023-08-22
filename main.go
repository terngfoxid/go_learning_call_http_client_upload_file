package main

import (
	_newRouter "go-back/routers"
)

func main() {

	r := _newRouter.SetupRouter()
	//running
	r.Run(":8081")
}
