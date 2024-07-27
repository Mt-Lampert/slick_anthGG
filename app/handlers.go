package main

import (
	"github.com/Mt-Lampert/slick_anthGG/model"
	"github.com/Mt-Lampert/slick_anthGG/slick"
	"github.com/Mt-Lampert/slick_anthGG/views/dashboard"
	"github.com/Mt-Lampert/slick_anthGG/views/profile"
)

func HandleDashboard(ctx *slick.Context) error {
	return ctx.Render(dashboard.Index())
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
