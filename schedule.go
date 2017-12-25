package huego

import(
  "encoding/json"
  "strconv"
)

type Schedule struct {
  Name	string `json:"name,omitempty"`
  Description	string `json:"description,omitempty"`
  Command	*Command  `json:"command,omitempty"`
  Time	string	`json:"time,omitempty"` // This should be of time Date
  LocalTime	string	`json:"localtime,omitempty"`
  StartTime	string	`json:"starttime,omitempty"`
  Status	string `json:"status,omitempty"`
  AutoDelete  bool	`json:"autodelete,omitempty"`
  Id  int `json:,omitempty`
}

type Command struct {
  Address string `json:"address,omitempty"`
  Method string `json:"method,omitempty"`
  Body interface{} `json:"body,omitempty"`
}

// Get all schedules
func (h *Hue) GetSchedules() ([]*Schedule, error) {

  var r map[string]Schedule

  res, err := h.GetResource("/schedules/")
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
func (h *Hue) CreateSchedule(s *Schedule) ([]*Response, error) {

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
func (h *Hue) UpdateSchedule(i int, schedule *Schedule) ([]*Response, error) {
  
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
func (h *Hue) DeleteSchedule(i int) ([]*Response, error) {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/schedules/", id)

	res, err := h.DeleteResource(url)
	if err != nil {
		return nil, err
	}

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
