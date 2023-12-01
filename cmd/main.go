package main

import (
	"Magical-Index-with-GPT/services/servicesimpl"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Supply(fx.Annotate(":8080", fx.ResultTags(`name:"hostPort"`))),
		servicesimpl.Module,
	)
	app.Run()
}
