package huego

// Schedule represents a bridge schedule https://developers.meethue.com/documentation/schedules-api-0
type Schedule struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Command     *Command `json:"command"`
	Time        string   `json:"time,omitempty"`
	LocalTime   string   `json:"localtime"`
	StartTime   string   `json:"starttime,omitempty"`
	Status      string   `json:"status,omitempty"`
	AutoDelete  bool     `json:"autodelete"`
	ID          int      `json:"-"`
}

// Command defines the request to be made when the schedule occurs
type Command struct {
	Address string      `json:"address"`
	Method  string      `json:"method"`
	Body    interface{} `json:"body"`
}
