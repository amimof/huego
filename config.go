package huego

import(
  "encoding/json"
  "strconv"
)

type Config struct {
  Name string `json:"name,omitempty"`
  SwUpdate *SwUpdate `json:"swupdate,omitempty"`
  SwUpdate2 *SwUpdate2 `json:"swupdate2,omitempty"`
  Whitelist []*Whitelist `json:"whitelist,omitempty"`
  PortalState *PortalState `json:"portalstate,omitempty"`
  ApiVersion string `json:"apiversion,omitempty"`
  SwVersion string `json:"swversion,omitempty"`
  ProxyAddress string `json:"proxyaddress,omitempty"`
  ProxyPort int `json:"proxyport,omitempty"`
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
  ZigbeeChannel string `json:"zigbeechannel,omitempty"`
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
  UpdateState int `json:"updatestate,omitempty"`
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
  LastChange string `json:"lastchange,omitempty"`
  LastInstall string `json:"lastinstall,omitempty"`
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
  Name string `json:"name,omitempty"`
  CreateDate string `json:"create date,omitempty"`
  LastUseDate string `json:"last use date,omitempty"`
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

  return config, nil

}

// Create a user
func (h *Hue) CreateUser(n string) ([]*Response, error) {

  var r []*Response

  url := h.GetApiUrl("/config/whitelist/")
  data, err := json.Marshal([]byte{`{"devicetype": "fixit"}`})
  if err != nil {
    return nil, err
  }

  res, err := h.PostResource(url, data)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  return r, nil

}

// Update configuration
func (h *Hue) UpdateConfig(i int, c *Config) ([]*Response, error) {
	var r []*Response

	id := strconv.Itoa(i)
	url := h.GetApiUrl("/config/")

	data, err := json.Marshal(&c)
	if err != nil {
		return r, err
	}

	res, err := h.PutResource(url, data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

// Delete a user from configuration
func (h *Hue) DeeteUser(n string) ([]*Response, error) {

  var r []*Response

  url := h.GetApiUrl("/config/whitelist/", n)

  res, err := h.DeleteResource(url)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &r)
  if err != nil {
    return nil, err
  }

  return r, nil

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
