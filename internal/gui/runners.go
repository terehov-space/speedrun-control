package gui

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"go-speedrun-control/internal/models/out"
	"strconv"
)

func renderRunners(ctx *context.Context) *fyne.Container {
	runnersCount, _ := strconv.Atoi((*ctx).Value("runners_count").(string))

	// reset basic init
	runnerScreenContainer := []fyne.CanvasObject{}
	runnersNameInputs = map[string]*widget.Entry{}
	runnersTimeInputs = map[string]*widget.Entry{}
	runnersSegmentInputs = map[string]*widget.Entry{}
	runnerProgressBar = map[string]*widget.ProgressBar{}

	// init runners rows
	makeRunnerRows(&runnerScreenContainer, runnersCount)

	saveRunnersButton = widget.NewButton("Save", saveRunners)

	runnerScreenContainer = append(runnerScreenContainer, container.NewVBox(
		saveRunnersButton,
	))

	startServerButton = widget.NewButton("Start", startServer)
	stopServerButton = widget.NewButton("Stop", stopServer)
	resetServerButton = widget.NewButton("Reset", resetRunners)

	saveRunnersButton.Enable()
	stopServerButton.Disable()
	resetServerButton.Disable()

	runnerScreenContainer = append(runnerScreenContainer,
		container.New(
			layout.NewGridLayout(3),
			startServerButton,
			stopServerButton,
			resetServerButton,
		),
	)

	runnersDataWrapper := container.NewVBox(runnerScreenContainer...)

	return runnersDataWrapper
}

func setPlayerData(runner out.Runner) {
	for runnerName, input := range runnersSegmentInputs {
		if runnerName == runner.Name {
			input.SetText(fmt.Sprintf("%v", runner.Segment))
		}
	}

	for runnerName, input := range runnersTimeInputs {
		if runnerName == runner.Name {
			input.SetText(runner.Time)
		}
	}

	for runnerName, progress := range runnerProgressBar {
		if runnerName == runner.Name {
			totalSegments, _ := strconv.ParseFloat((*ctx).Value("segments_count").(string), 64)

			progress.SetValue(roundFloat(runner.Segment/totalSegments, 2))
		}
	}
}

func makeRunnerRows(screenContainer *[]fyne.CanvasObject, count int) {
	runners = map[string]out.Runner{}

	for i := 0; i < count; i++ {
		runnerName := fmt.Sprintf("Default_%v", i)

		runnerNameInput := widget.NewEntry()
		runnerNameInput.SetPlaceHolder("Enter runner name...")
		runnerNameInput.SetText(runnerName)

		runnerPercInput := widget.NewEntry()
		runnerPercInput.SetPlaceHolder("Enter runner segment...")
		runnerPercInput.SetText("0")

		runnerTimeInput := widget.NewEntry()
		runnerTimeInput.SetPlaceHolder("Enter runner segment...")
		runnerTimeInput.SetText("00:00:00.00")

		runnersNameInputs[runnerName] = runnerNameInput
		runnersTimeInputs[runnerName] = runnerTimeInput
		runnersSegmentInputs[runnerName] = runnerPercInput

		runners[runnerName] = out.Runner{
			Name:    runnerName,
			Time:    "00:00:00.00",
			Segment: 0.0,
		}

		runnerContainer := container.New(
			layout.NewGridLayout(3),
			runnerNameInput,
			runnerPercInput,
			runnerTimeInput,
		)

		progress := widget.NewProgressBar()
		runnerProgressBar[runnerName] = progress

		*screenContainer = append(*screenContainer, runnerContainer, progress)
	}
}

func disableRunnerNameInputs() {
	for _, input := range runnersNameInputs {
		input.Disable()
	}
}

func enableRunnerNameInputs() {
	for _, input := range runnersNameInputs {
		input.Enable()
	}
}

func saveRunners() {
	runners = out.RunnerList{}

	for runnerName, _ := range runnersNameInputs {
		segments, _ := strconv.ParseFloat(runnersSegmentInputs[runnerName].Text, 64)
		runnerData := out.Runner{
			Name:    runnersNameInputs[runnerName].Text,
			Time:    runnersTimeInputs[runnerName].Text,
			Segment: segments,
		}

		runners[runnerName] = runnerData
		setPlayerData(runnerData)
	}

	fmt.Println(runners)

	saveRunnersButton.Enable()
	resetServerButton.Enable()
}

func resetRunners() {
	runners = out.RunnerList{}
	index := 0

	for _, input := range runnersNameInputs {
		runnerName := fmt.Sprintf("Default_%v", index)
		input.SetText(runnerName)
		index++
		runnerData := out.Runner{
			Name:    runnerName,
			Segment: 0.0,
			Time:    "00:00:00.00",
		}

		setPlayerData(runnerData)

		runners[runnerName] = runnerData
	}

	saveRunnersButton.Enable()
	resetServerButton.Enable()
}

func startServer() {
	blockSettings()
	disableRunnerNameInputs()
	StartServer()

	saveRunnersButton.Enable()
	startServerButton.Disable()
	stopServerButton.Enable()
	resetServerButton.Disable()
}

func stopServer() {
	unlockSettings()
	enableRunnerNameInputs()
	StopServer()

	saveRunnersButton.Enable()
	startServerButton.Enable()
	stopServerButton.Disable()
	resetServerButton.Enable()
}
