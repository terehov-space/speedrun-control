package out

type RunnerList map[string]Runner

type Runner struct {
	Name    string  `json:"name"`
	Segment float64 `json:"segment"`
	Time    string  `json:"time"`
}
