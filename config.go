package huego

// Config holds the bridge hardware configuration
type Config struct {
	Name             string               `json:"name,omitempty"`
	SwUpdate         SwUpdate             `json:"swupdate"`
	SwUpdate2        SwUpdate2            `json:"swupdate2"`
	WhitelistMap     map[string]Whitelist `json:"whitelist"`
	Whitelist        []Whitelist          `json:"-"`
	PortalState      PortalState          `json:"portalstate"`
	APIVersion       string               `json:"apiversion,omitempty"`
	SwVersion        string               `json:"swversion,omitempty"`
	ProxyAddress     string               `json:"proxyaddress,omitempty"`
	ProxyPort        uint16               `json:"proxyport,omitempty"`
	LinkButton       bool                 `json:"linkbutton,omitempty"`
	IPAddress        string               `json:"ipaddress,omitempty"`
	Mac              string               `json:"mac,omitempty"`
	NetMask          string               `json:"netmask,omitempty"`
	Gateway          string               `json:"gateway,omitempty"`
	Dhcp             bool                 `json:"dhcp,omitempty"`
	PortalServices   bool                 `json:"portalservices,omitempty"`
	UTC              string               `json:"UTC,omitempty"`
	LocalTime        string               `json:"localtime,omitempty"`
	TimeZone         string               `json:"timezone,omitempty"`
	ZigbeeChannel    uint8                `json:"zigbeechannel,omitempty"`
	ModelID          string               `json:"modelid,omitempty"`
	BridgeID         string               `json:"bridgeid,omitempty"`
	FactoryNew       bool                 `json:"factorynew,omitempty"`
	ReplacesBridgeID string               `json:"replacesbridgeid,omitempty"`
	DatastoreVersion string               `json:"datastoreversion,omitempty"`
	StarterKitID     string               `json:"starterkitid,omitempty"`
	InternetService  InternetService      `json:"internetservices,omitempty"`
}

// SwUpdate contains information related to software updates. Deprecated in 1.20
type SwUpdate struct {
	CheckForUpdate bool        `json:"checkforupdate,omitempty"`
	DeviceTypes    DeviceTypes `json:"devicetypes"`
	UpdateState    uint8       `json:"updatestate,omitempty"`
	Notify         bool        `json:"notify,omitempty"`
	URL            string      `json:"url,omitempty"`
	Text           string      `json:"text,omitempty"`
}

// SwUpdate2 contains information related to software updates
type SwUpdate2 struct {
	Bridge         BridgeConfig `json:"bridge"`
	CheckForUpdate bool         `json:"checkforupdate,omitempty"`
	State          string       `json:"state,omitempty"`
	Install        bool         `json:"install,omitempty"`
	AutoInstall    AutoInstall  `json:"autoinstall"`
	LastChange     string       `json:"lastchange,omitempty"`
	LastInstall    string       `json:"lastinstall,omitempty"`
}

// DeviceTypes details the type of updates available
type DeviceTypes struct {
	Bridge  bool     `json:"bridge,omitempty"`
	Lights  []Light  `json:"lights,omitempty"`
	Sensors []Sensor `json:"sensors,omitempty"`
}

// BridgeConfig holds information about software updates
type BridgeConfig struct {
	State       string `json:"state,omitempty"`
	LastInstall string `json:"lastinstall,omitempty"`
}

// AutoInstall holds automatic update configuration
type AutoInstall struct {
	On         bool   `json:"on,omitempty"`
	UpdateTime string `json:"updatetime,omitempty"`
}

// InternetService stores information about the internet connectivity to the bridge
type InternetService struct {
	Internet     string `json:"internet,omitempty"`
	RemoteAccess string `json:"remoteaccess,omitempty"`
	Time         string `json:"time,omitempty"`
	SwUpdate     string `json:"swupdate,omitempty"`
}

// Backup holds configuration backup status information
type Backup struct {
	Status    string `json:"backup,omitempty"`
	ErrorCode int    `json:"errorcode,omitempty"`
}

// Whitelist represents a whitelist user ID in the bridge
type Whitelist struct {
	Name        string `json:"name"`
	Username    string
	CreateDate  string `json:"create date"`
	LastUseDate string `json:"last use date"`
}

// PortalState is a struct representing the portal state
type PortalState struct {
	SignedOn      bool   `json:"signedon,omitempty"`
	Incoming      bool   `json:"incoming,omitempty"`
	Outgoing      bool   `json:"outgoing,omitempty"`
	Communication string `json:"communication,omitempty"`
}

// Datastore combines all bridge resources into one struct
type Datastore struct {
	Lights    []Light    `json:"lights"`
	Groups    []Group    `json:"groups"`
	Config    Config     `json:"config"`
	Schedules []Schedule `json:"schedules"`
	Scenes    []Scene    `json:"scenes"`
	Sensors   []Sensor   `json:"sensors"`
	Rules     []Rule     `json:"rules"`
}
