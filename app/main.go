package main

import (
	"fmt"

	"github.com/Mt-Lampert/slick_anthGG/model"
	"github.com/Mt-Lampert/slick_anthGG/slick"
	"github.com/Mt-Lampert/slick_anthGG/views/profile"
)

func main() {
	app := slick.New()

	app.Get(`/profile`, HandleProfile)

	fmt.Println(`Running Slick framework at port :3000`)
	app.Start(`:3000`)
}

func HandleProfile(ctx *slick.Context) error {
	user := &model.User{
		FirstName: `Matthias`,
		LastName:  `Langbart`,
		Email:     `perrin.blog@web.de`,
	}
	return ctx.Render(profile.Index(user))
}

// vim: ts=4 sw=4 fdm=indent
