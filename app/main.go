package main

import (
	"fmt"

	"github.com/Mt-Lampert/slick_anthGG/slick"
)

func main() {
	app := slick.New()

	app.Get(`/profile`, HandleProfile)

	fmt.Println(`Running Slick framework at port :3000`)
	app.Start(`:3000`)
}

// vim: ts=4 sw=4 fdm=indent
