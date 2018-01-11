package huego

import(
  "encoding/json"
)

type Config struct {
  Name string `json:"name,omitempty"`
  SwUpdate *SwUpdate `json:"swupdate"`
  SwUpdate2 *SwUpdate2 `json:"swupdate2"`
  WhitelistMap map[string]Whitelist `json:"whitelist"`
  Whitelist []Whitelist `json:"-"`
  PortalState *PortalState `json:"portalstate"`
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
  DeviceTypes *DeviceTypes `json:"devicetypes"`
  UpdateState uint8 `json:"updatestate,omitempty"`
  Notify bool `json:"notify,omitempty"`
  Url string `json:"url,omitempty"`
  Text string `json:"text,omitempty"`
}

type DeviceTypes struct {
  Bridge bool `json:"bridge,omitempty"`
  Lights []Light `json:"lights,omitempty"`
  Sensors []Sensor `json:"sensors,omitempty"`
}

type SwUpdate2 struct {
  Bridge *BridgeConfig `json:"bridge"`
  CheckForUpdate bool `json:"checkforupdate,omitempty"`
  State string `json:"state,omitempty"`
  Install bool `json:"install,omitempty"`
  AutoInstall *AutoInstall `json:"autoinstall"`
  LastChange string `json:"lastchange,omitempty"`
  LastInstall string `json:"lastinstall,omitempty"`
}

type BridgeConfig struct {
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
  Lights []Light `json:"lights"`
  Groups []Group `json:"groups"`
  Config *Config `json:"config"`
  Schedules []Schedule `json:"schedules"`
  Scenes []Scene `json:"scenes"`
  Sensors []Sensor `json:"sensors"`
  Rules []Rule `json:"rules"`
}


// Returns the bridge configuration
func (b *Bridge) GetConfig() (*Config, error) {
  
  var config *Config

  url, err := b.getApiPath("/config/")
  res, err := get(url)
  if err != nil {
    return nil, err
  }

  err = json.Unmarshal(res, &config)
  if err != nil {
    return nil, err
  }

  wl := make([]Whitelist, 0, len(config.WhitelistMap))
  for k, v := range config.WhitelistMap {
    v.Username = k
    wl = append(wl, v)
  }

  config.Whitelist = wl

  return config, nil

}

// Creates a user by adding n to the list of whitelists in the bridge
func (b *Bridge) CreateUser(n string) (string, error) {

  var a []*ApiResponse

  body := struct {
    DeviceType string `json:"devicetype,omitempty"`
    GenerateClientKey bool `json:"generateclientkey,omitempty"`
  }{
    n,
    true,
  }

  url, err := b.getApiPath("/")
  if err != nil {
    return "", err
  }

  data, err := json.Marshal(&body)
  if err != nil {
    return "", err
  }

  res, err := post(url, data)
  if err != nil {
    return "", err
  }

  err = json.Unmarshal(res, &a)
  if err != nil {
    return "", err
  }

  resp, err := handleResponse(a)
  if err != nil {
    return "", err
  }

  return resp.Success["username"].(string), nil

}

// Returns a list of whitelists from the bridge
func (b *Bridge) GetUsers() ([]Whitelist, error) {
  c, err := b.GetConfig()
  if err != nil {
    return nil, err
  }
  return c.Whitelist, nil
}

// Updates the bridge configuration with c
func (b *Bridge) UpdateConfig(c *Config) (*Response, error) {

	var a []*ApiResponse

  url, err := b.getApiPath("/config/")
  if err != nil {
    return nil, err
  }

	data, err := json.Marshal(&c)
	if err != nil {
		return nil, err
	}

	res, err := put(url, data)
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

// Removes a whitelist item from whitelists on the bridge
func (b *Bridge) DeleteUser(n string) error {

  var a []*ApiResponse

  url, err := b.getApiPath("/config/whitelist/", n)
  if err != nil {
    return err
  }

  res, err := delete(url)
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
func (b *Bridge) GetFullState() (*Datastore, error) {

    var ds *Datastore

    url, err := b.getApiPath("/")
    if err != nil {
      return nil, err
    }

    res, err := get(url)
    if err != nil {
      return nil, err
    }

    err = json.Unmarshal(res, &ds)
    if err != nil {
      return nil, err
    }

    return ds, nil
}
