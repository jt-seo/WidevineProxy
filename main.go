package main

import (
	"com.jtworld.wvproxy/router"
)

func main() {
	e := router.New()

	e.Start(":8080")
}
