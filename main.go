package main

import (
	"gin-app/routes"
)

func main() {

	r := routes.SetupRouter()

	r.Run(":3000")
}
