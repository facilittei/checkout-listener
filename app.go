package main

import (
	"github.com/facilittei/checkout-listener/adapters/runners"
)

// App entrypoint
type App struct {
	Runner runners.Runner
}

// NewApp creates a new instance of the application
func NewApp(runner runners.Runner) *App {
	return &App{
		Runner: runner,
	}
}

// Run application
func (app *App) Run() {
	app.Runner.Run()
}
