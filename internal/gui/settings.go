package gui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// renderSettings - render settings menu
func renderSettings(ctx *context.Context) *fyne.Container {
	segmentsLabel := widget.NewLabel("Segments")

	segmentCountInput = widget.NewEntry()
	segmentCountInput.SetText((*ctx).Value("segments_count").(string))
	segmentCountInput.SetPlaceHolder("Enter segments count...")

	runnersLabel := widget.NewLabel("Runners")

	runnersCountInput = widget.NewEntry()
	runnersCountInput.SetText((*ctx).Value("runners_count").(string))
	runnersCountInput.SetPlaceHolder("Enter runners count...")

	serverPortLabel := widget.NewLabel("Server port")

	serverPortInput = widget.NewEntry()
	serverPortInput.SetText((*ctx).Value("server_port").(string))
	serverPortInput.SetPlaceHolder("Enter server port...")

	saveSettingsButton = widget.NewButton("Save settings", func() {
		*ctx = context.WithValue(*ctx, "segments_count", segmentCountInput.Text)
		*ctx = context.WithValue(*ctx, "runners_count", runnersCountInput.Text)
		*ctx = context.WithValue(*ctx, "server_port", serverPortInput.Text)
		*ctx = context.WithValue(*ctx, "changed", "1")
	})

	return container.NewVBox(
		segmentsLabel,
		segmentCountInput,
		runnersLabel,
		runnersCountInput,
		serverPortLabel,
		serverPortInput,
		saveSettingsButton,
	)
}

// block settings inputs
func blockSettings() {
	segmentCountInput.Disable()
	runnersCountInput.Disable()
	serverPortInput.Disable()
	saveSettingsButton.Disable()
}

// unlock settings inputs
func unlockSettings() {
	segmentCountInput.Enable()
	runnersCountInput.Enable()
	serverPortInput.Enable()
	saveSettingsButton.Enable()
}
