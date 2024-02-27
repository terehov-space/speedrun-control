package gui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

// Show - open control window
func Show(c *context.Context) {
	ctx = c

	// set default http handlers
	setHttpHandlers()

	speedRunApp := app.New()
	window := speedRunApp.NewWindow("SpeedRun")

	renderApp(window)

	go func() {
		for {
			renderApp(window)
		}
	}()

	window.Resize(fyne.Size{Width: 800, Height: 600})

	window.ShowAndRun()
}

// renderApp - render main tabs
func renderApp(window fyne.Window) {
	if (*ctx).Value("changed") == "1" {
		tabs := container.NewAppTabs(

			container.NewTabItem("Сервер", renderSettings(ctx)),
			container.NewTabItem("Игроки", renderRunners(ctx)),
		)

		tabs.SetTabLocation(container.TabLocationLeading)

		window.SetContent(tabs)
		*ctx = context.WithValue(*ctx, "changed", "0")
	}
}
