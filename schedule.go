package huego

import(
  "encoding/json"
  "strconv"
  "fmt"
)

type Schedule struct {
  Name	string `json:"name"`
  Description	string `json:"description"`
  Command	*Command  `json:"command"`
  Time	string	`json:"time,omitempty"`
  LocalTime	string	`json:"localtime"`
  StartTime	string	`json:"starttime,omitempty"`
  Status	string `json:"status,omitempty"`
  AutoDelete  bool	`json:"autodelete"`
  Id  int `json:"-"`
}

type Command struct {
  Address string `json:"address"`
  Method string `json:"method"`
  Body interface{} `json:"body"`
}

// Get all schedules
func (h *Hue) GetSchedules() ([]*Schedule, error) {

  var r map[string]Schedule

  res, err := h.GetResource(h.GetApiUrl("/schedules/"))
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  schedules := make([]*Schedule, 0, len(r))

  for i, s := range r {
    s.Id, err = strconv.Atoi(i)
    if err != nil {
      return nil, err
    }
    schedules = append(schedules, &s)
  }

  return schedules, nil

}

// Get one schedule
func (h *Hue) GetSchedule(i int) (*Schedule, error) {

	var schedule *Schedule

	res, err := h.GetResource(h.GetApiUrl("/schedules/", strconv.Itoa(i)))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil

}

// Create a schedule
func (h *Hue) CreateSchedule(s *Schedule) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url := h.GetApiUrl("/schedules/")
  res, err := h.PostResource(url, data)
  if err != nil {
    return nil, err
  }

  fmt.Println(string(res))

  err = json.Unmarshal(res, &a)
  if err != nil {
    return nil, err
  }

  resp, err := handleResponse(a)
  if err != nil {
    return nil, err
  }

  return resp, nil

}

// Update a schedule
func (h *Hue) UpdateSchedule(i int, schedule *Schedule) (*Response, error) {
  
  var a []*ApiResponse

	data, err := json.Marshal(&schedule)
	if err != nil {
		return nil, err
	}

	url := h.GetApiUrl("/schedules/", strconv.Itoa(i))
	res, err := h.PutResource(url, data)
	if err != nil {
		return nil, err
  }

  fmt.Println(string(res))

	err = json.Unmarshal(res, &a)
	if err != nil {
		return nil, err
  }
  
  resp, err := handleResponse(a)
  if err != nil {
    return nil, err
  }

	return resp, nil
}

// Delete a schedule
func (h *Hue) DeleteSchedule(i int) (*Response, error) {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/schedules/", id)

	res, err := h.DeleteResource(url)
	if err != nil {
		return nil, err
	}

  fmt.Println(string(res))

  err = json.Unmarshal(res, &a)
  if err != nil {
    return nil, err
  }
  
  resp, err := handleResponse(a)
  if err != nil {
    return nil, err
  }

	return resp, nil
}
