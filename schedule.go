package huego

import(
  "encoding/json"
  "strconv"
)

// https://developers.meethue.com/documentation/schedules-api-0
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

// Returns all scehdules known to the bridge
func (b *Bridge) GetSchedules() ([]*Schedule, error) {

  var r map[string]Schedule

  url, err := b.getApiPath("/schedules/")
  if err != nil {
    return nil, err
  }

  res, err := b.getResource(url)
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

// Returns one schedule by id defined in i
func (b *Bridge) GetSchedule(i int) (*Schedule, error) {

	var schedule *Schedule

  url, err := b.getApiPath("/schedules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.getResource(url)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &schedule)
	if err != nil {
		return nil, err
	}

	return schedule, nil

}

// Creates one schedule and sets its attributes defined in s
func (b *Bridge) CreateSchedule(s *Schedule) (*Response, error) {

  var a []*ApiResponse

  data, err := json.Marshal(&s)
  if err != nil {
    return nil, err
  }

  url, err := b.getApiPath("/schedules/")
	if err != nil {
		return nil, err
	}

  res, err := b.postResource(url, data)
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

// Updates one schedule by its id of i and attributes by schedule
func (b *Bridge) UpdateSchedule(i int, schedule *Schedule) (*Response, error) {
  
  var a []*ApiResponse

	data, err := json.Marshal(&schedule)
	if err != nil {
		return nil, err
	}

  url, err := b.getApiPath("/schedules/", strconv.Itoa(i))
	if err != nil {
		return nil, err
	}

	res, err := b.putResource(url, data)
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

// Deletes one schedule from the bridge by its id of i
func (b *Bridge) DeleteSchedule(i int) error {
  
  var a []*ApiResponse

	id := strconv.Itoa(i)
  url, err := b.getApiPath("/schedules/", id)
	if err != nil {
		return err
	}

	res, err := b.deleteResource(url)
	if err != nil {
		return err
	}

  _ = json.Unmarshal(res, &a)
  
  _, err = handleResponse(a)
  if err != nil {
    return err
  }

	return nil
}
