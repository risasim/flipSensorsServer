package database

type DataEntry struct {
	TimeStamp   string  `json:"timestamp"`
	Temperature float64 `json:"temperature"`
	GasLevel    float64 `json:"gasLevel"`
}
