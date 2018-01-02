package huego

import(
  "encoding/json"
)

type Config struct {
  Name string `json:"name,omitempty"`
  SwUpdate *SwUpdate `json:"swupdate,omitempty"`
  SwUpdate2 *SwUpdate2 `json:"swupdate2,omitempty"`
  WhitelistMap map[string]*Whitelist `json:"whitelist"`
  Whitelist []*Whitelist `json:"-"`
  PortalState *PortalState `json:"portalstate,omitempty"`
  ApiVersion string `json:"apiversion,omitempty"`
  SwVersion string `json:"swversion,omitempty"`
  ProxyAddress string `json:"proxyaddress,omitempty"`
  ProxyPort uint16 `json:"proxyport,omitempty"`
  LinkButton bool `json:"linkbutton,omitempty"`
  IpAddress string `json:"ipaddress,omitempty"`
  Mac string `json:"mac,omitempty"`
  NetMask string `json:"netmask,omitempty"`
  Gateway string `json:"gateway,omitempty"`
  Dhcp bool `json:"dhcp,omitempty"`
  PortalServices bool `json:"portalservices,omitempty"`
  UTC string `json:"UTC,omitempty"`
  LocalTime string `json:"localtime,omitempty"`
  TimeZone string `json:"timezone,omitempty"`
  ZigbeeChannel uint8 `json:"zigbeechannel,omitempty"`
  ModelId string `json:"modelid,omitempty"`
  BridgeId string `json:"bridgeid,omitempty"`
  FactoryNew bool `json:"factorynew,omitempty"`
  ReplacesBridgeId string `json:"replacesbridgeid,omitempty"`
  DatastoreVersion string `json:"datastoreversion,omitempty"`
  StarterKitId string `json:"starterkitid,omitempty"`
}

type SwUpdate struct {
  CheckForUpdate bool `json:"checkforupdate,omitempty"`
  DeviceTypes *DeviceTypes `json:"devicetypes,omitempty"`
  UpdateState uint8 `json:"updatestate,omitempty"`
  Notify bool `json:"notify,omitempty"`
  Url string `json:"url,omitempty"`
  Text string `json:"text,omitempty"`
}

type DeviceTypes struct {
  Bridge bool `json:"bridge,omitempty"`
  Lights []*Light `json:"lights,omitempty"`
  Sensors []*Sensor `json:"sensors,omitempty"`
}

type SwUpdate2 struct {
  Bridge *Bridge `json:"bridge,omitempty"`
  CheckForUpdate bool `json:"checkforupdate,omitempty"`
  State string `json:"state,omitempty"`
  Install bool `json:"install,omitempty"`
  AutoInstall *AutoInstall `json:"autoinstall,omitempty"`
  LastChange string `json:"lastchange,omitempty"` // Should be of type Time
  LastInstall string `json:"lastinstall,omitempty"` // SHould be of type Time
}

type Bridge struct {
  State string `json:"state,omitempty"`
  LastInstall string `json:"lastinstall,omitempty"`
}

type AutoInstall struct {
  On bool `json:"on,omitempty"`
  UpdateTime string `json:"updatetime,omitempty"`
}

type InternetService struct {
  Internet string `json:"internet,omitempty"`
  RemoteAccess string `json:"remoteaccess,omitempty"`
  Time string `json:"time,omitempty"`
  SwUpdate string `json:"swupdate,omitempty"`
}

type Backup struct {
  Status string `json:"backup,omitempty"`
  ErrorCode int `json:"errorcode,omitempty"`
}

type Whitelist struct {
  Name string `json:"name"`
  Username string
  CreateDate string `json:"create date"`
  LastUseDate string `json:"last use date"`
}

type PortalState struct {
  SignedOn bool `json:"signedon,omitempty"`
  Incoming bool `json:"incoming,omitempty"`
  Outgoing bool `json:"outgoing,omitempty"`
  Communication string `json:"communication,omitempty"`
}

type Datastore struct {
  Lights []*Light `json:"lights,omitempty"`
  Groups []*Group `json:"groups,omitempty"`
  Config *Config `json:"config,omitempty"`
  Schedules []*Schedule `json:"schedules,omitempty"`
  Scenes []*Scene `json:"scenes,omitempty"`
  Sensors []*Sensor `json:"sensors,omitempty"`
  Rules []*Rule `json:"rules,omitempty"`
}


// Get configuration
func (h *Hue) GetConfig() (*Config, error) {
  
  var config *Config

  url := h.GetApiUrl("/config/")
  res, err := h.GetResource(url)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &config)
  if err != nil {
    return nil, err
  }

  wl := make([]*Whitelist, 0, len(config.WhitelistMap))
  for k, v := range config.WhitelistMap {
    v.Username = k
    wl = append(wl, v)
  }

  config.Whitelist = wl

  return config, nil

}

// Create a user
func (h *Hue) CreateUser(n string) (*Response, error) {

  var a []*ApiResponse

  body := struct {
    DeviceType string `json:"devicetype,omitempty"`
    GenerateClientKey bool `json:"generateclientkey,omitempty"`
  }{
    n,
    true,
  }

  url := h.GetApiUrl("/")
  data, err := json.Marshal(&body)
  if err != nil {
    return nil, err
  }

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

// Returns Whitelist 
func (h *Hue) GetUsers() ([]*Whitelist, error) {
  c, err := h.GetConfig()
  if err != nil {
    return nil, err
  }
  return c.Whitelist, nil
}

// Update configuration
func (h *Hue) UpdateConfig(c *Config) (*Response, error) {

	var a []*ApiResponse

	url := h.GetApiUrl("/config/")

	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

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

// Delete a user from configuration
func (h *Hue) DeleteUser(n string) error {

  var a []*ApiResponse

  url := h.GetApiUrl("/config/whitelist/", n)

  res, err := h.DeleteResource(url)
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

// Get full state (datastore)
func (h *Hue) GetFullState() (*Datastore, error) {

    var ds *Datastore

    url := h.GetApiUrl("/")
    res, err := h.GetResource(url)
    if err != nil {
      return nil, err
    }

    err = json.Unmarshal(res, &ds)
    if err != nil {
      return nil, err
    }

    return ds, nil
}
