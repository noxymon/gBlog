package main

import (
	"noxymon.web.id/internal"
)

func main() {
	e := internal.New()
	e.Logger.Fatal(e.Start(":8080"))
}
