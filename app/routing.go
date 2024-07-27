package main

import "github.com/Mt-Lampert/slick_anthGG/slick"

// handles routing for the app
func routes(app *slick.Slick) {
	app.Get(`/profile`, HandleProfile)
}

// vim: ts=4 sw=4 fdm=indent
