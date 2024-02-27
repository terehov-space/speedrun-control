package gui

import (
	"encoding/json"
	"fmt"
	"go-speedrun-control/internal/models/out"
	"go-speedrun-control/internal/models/syberia"
	"io"
	"math"
	"net/http"
	"strconv"
)

func StartServer() {
	server = &http.Server{Addr: fmt.Sprintf(":%v", (*ctx).Value("server_port").(string))}

	// start server
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			fmt.Println("Error", err)
		}
	}()
}

func StopServer() {
	_ = server.Shutdown(*ctx)
}

func setHttpHandlers() {
	http.HandleFunc("/set", setRunnerData)

	http.HandleFunc("/get", getRunnersData)

	http.HandleFunc("/segments", getRunnerSegments)

	http.HandleFunc("/times", getRunnerTimes)
}

func setRunnerData(w http.ResponseWriter, r *http.Request) {
	requestData, _ := io.ReadAll(r.Body)
	runner := syberia.Runner{}

	_ = json.Unmarshal(requestData, &runner)

	_, exists := runners[runner.Runner]

	if exists {
		runnerData := out.Runner{
			Name:    runner.Runner,
			Time:    runner.GTimeStr,
			Segment: runner.SplitIndex,
		}

		setPlayerData(runnerData)

		runners[runner.Runner] = runnerData
	}

	_, _ = io.WriteString(w, string(requestData))
}

func getRunnersData(w http.ResponseWriter, r *http.Request) {
	responseData, _ := json.Marshal(runners)
	_, _ = io.WriteString(w, string(responseData))
}

func getRunnerSegments(w http.ResponseWriter, r *http.Request) {
	segmentCount, _ := strconv.ParseFloat((*ctx).Value("segments_count").(string), 64)
	responseResource := map[string]float64{
		"segments": segmentCount,
	}

	for _, runner := range runners {
		fmt.Println(roundFloat(runner.Segment/segmentCount, 2))
		responseResource[runner.Name] = roundFloat(runner.Segment/segmentCount, 2)
	}

	responseData, _ := json.Marshal(responseResource)

	_, _ = io.WriteString(w, string(responseData))
}

func getRunnerTimes(w http.ResponseWriter, r *http.Request) {
	responseData := map[string]string{}

	for _, runner := range runners {
		responseData[runner.Name] = runner.Time
	}

	response, _ := json.Marshal(responseData)

	_, _ = io.WriteString(w, string(response))
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
