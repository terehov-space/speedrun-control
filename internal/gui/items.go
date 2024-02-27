package gui

import (
	"context"
	"fyne.io/fyne/v2/widget"
	"go-speedrun-control/internal/models/out"
	"net/http"
)

// context
var ctx *context.Context

// models
var runners out.RunnerList

// inputs
var runnersNameInputs map[string]*widget.Entry
var runnersSegmentInputs map[string]*widget.Entry
var runnersTimeInputs map[string]*widget.Entry

var segmentCountInput *widget.Entry
var runnersCountInput *widget.Entry
var serverPortInput *widget.Entry

// progress
var runnerProgressBar map[string]*widget.ProgressBar

// control buttons
var saveRunnersButton *widget.Button
var startServerButton *widget.Button
var stopServerButton *widget.Button
var resetServerButton *widget.Button

var saveSettingsButton *widget.Button

// server
var server *http.Server
