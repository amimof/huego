package huego

import "context"

// Group represents a bridge group https://developers.meethue.com/documentation/groups-api
type Group struct {
	Name       string               `json:"name,omitempty"`
	Lights     []string             `json:"lights,omitempty"`
	Type       string               `json:"type,omitempty"`
	GroupState *GroupState          `json:"state,omitempty"`
	Recycle    bool                 `json:"recycle,omitempty"`
	Class      string               `json:"class,omitempty"`
	Stream     *Stream              `json:"stream,omitempty"`
	Locations  map[string][]float64 `json:"locations,omitempty"`
	State      *State               `json:"action,omitempty"`
	ID         int                  `json:"-"`
	bridge     *Bridge
}

// GroupState defines the state on a group.
// Can be used to control the state of all lights in a group rather than controlling them individually
type GroupState struct {
	AllOn bool `json:"all_on,omitempty"`
	AnyOn bool `json:"any_on,omitempty"`
}

// Stream define the stream status of a group
type Stream struct {
	ProxyMode string  `json:"proxymode,omitempty"`
	ProxyNode string  `json:"proxynode,omitempty"`
	ActiveRaw *bool   `json:"active,omitempty"`
	OwnerRaw  *string `json:"owner,omitempty"`
}

// Active returns the stream active state, and will return false if ActiveRaw is nil
func (s *Stream) Active() bool {
	if s.ActiveRaw == nil {
		return false
	}

	return *s.ActiveRaw
}

// Owner returns the stream Owner, and will return an empty string if OwnerRaw is nil
func (s *Stream) Owner() string {
	if s.OwnerRaw == nil {
		return ""
	}

	return *s.OwnerRaw
}

// SetState sets the state of the group to s.
func (g *Group) SetState(s State) error {
	return g.SetStateContext(context.Background(), s)
}

// SetStateContext sets the state of the group to s.
func (g *Group) SetStateContext(ctx context.Context, s State) error {
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, s)
	if err != nil {
		return err
	}
	g.State = &s
	return nil
}

// Rename sets the name property of the group
func (g *Group) Rename(new string) error {
	return g.RenameContext(context.Background(), new)
}

// RenameContext sets the name property of the group
func (g *Group) RenameContext(ctx context.Context, new string) error {
	update := Group{Name: new}
	_, err := g.bridge.UpdateGroupContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.Name = new
	return nil
}

// Off sets the On state of one group to false, turning all lights in the group off
func (g *Group) Off() error {
	return g.OffContext(context.Background())
}

// OffContext sets the On state of one group to false, turning all lights in the group off
func (g *Group) OffContext(ctx context.Context) error {
	state := State{On: false}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, state)
	if err != nil {
		return err
	}
	g.State.On = false
	return nil
}

// On sets the On state of one group to true, turning all lights in the group on
func (g *Group) On() error {
	return g.OnContext(context.Background())
}

// OnContext sets the On state of one group to true, turning all lights in the group on
func (g *Group) OnContext(ctx context.Context) error {
	state := State{On: true}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, state)
	if err != nil {
		return err
	}
	g.State.On = true
	return nil
}

// IsOn returns true if light state On property is true
func (g *Group) IsOn() bool {
	return g.State.On
}

// Bri sets the light brightness state property
func (g *Group) Bri(new uint8) error {
	return g.BriContext(context.Background(), new)
}

// BriContext sets the light brightness state property
func (g *Group) BriContext(ctx context.Context, new uint8) error {
	update := State{On: true, Bri: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Bri = new
	return nil
}

// Hue sets the light hue state property (0-65535)
func (g *Group) Hue(new uint16) error {
	return g.HueContext(context.Background(), new)
}

// HueContext sets the light hue state property (0-65535)
func (g *Group) HueContext(ctx context.Context, new uint16) error {
	update := State{On: true, Hue: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Hue = new
	return nil
}

// Sat sets the light saturation state property (0-254)
func (g *Group) Sat(new uint8) error {
	return g.SatContext(context.Background(), new)
}

// SatContext sets the light saturation state property (0-254)
func (g *Group) SatContext(ctx context.Context, new uint8) error {
	update := State{On: true, Sat: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Sat = new
	return nil
}

// Xy sets the x and y coordinates of a color in CIE color space. (0-1 per value)
func (g *Group) Xy(new []float32) error {
	return g.XyContext(context.Background(), new)
}

// XyContext sets the x and y coordinates of a color in CIE color space. (0-1 per value)
func (g *Group) XyContext(ctx context.Context, new []float32) error {
	update := State{On: true, Xy: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Xy = new
	return nil
}

// Ct sets the light color temperature state property
func (g *Group) Ct(new uint16) error {
	return g.CtContext(context.Background(), new)
}

// CtContext sets the light color temperature state property
func (g *Group) CtContext(ctx context.Context, new uint16) error {
	update := State{On: true, Ct: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Ct = new
	return nil
}

// Scene sets the scene by it's identifier of the scene you wish to recall
func (g *Group) Scene(scene string) error {
	return g.SceneContext(context.Background(), scene)
}

// SceneContext sets the scene by it's identifier of the scene you wish to recall
func (g *Group) SceneContext(ctx context.Context, scene string) error {
	update := State{On: true, Scene: scene}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Scene = scene
	return nil
}

// TransitionTime sets the duration of the transition from the light’s current state to the new state
func (g *Group) TransitionTime(new uint16) error {
	return g.TransitionTimeContext(context.Background(), new)
}

// TransitionTimeContext sets the duration of the transition from the light’s current state to the new state
func (g *Group) TransitionTimeContext(ctx context.Context, new uint16) error {
	update := State{On: g.State.On, TransitionTime: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.TransitionTime = new
	return nil
}

// Effect the dynamic effect of the lights in the group, currently “none” and “colorloop” are supported
func (g *Group) Effect(new string) error {
	return g.EffectContext(context.Background(), new)
}

// EffectContext the dynamic effect of the lights in the group, currently “none” and “colorloop” are supported
func (g *Group) EffectContext(ctx context.Context, new string) error {
	update := State{On: true, Effect: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Effect = new
	return nil
}

// Alert makes the lights in the group blink in its current color. Supported values are:
// “none” – The light is not performing an alert effect.
// “select” – The light is performing one breathe cycle.
// “lselect” – The light is performing breathe cycles for 15 seconds or until alert is set to "none".
func (g *Group) Alert(new string) error {
	return g.AlertContext(context.Background(), new)
}

// AlertContext makes the lights in the group blink in its current color. Supported values are:
// “none” – The light is not performing an alert effect.
// “select” – The light is performing one breathe cycle.
// “lselect” – The light is performing breathe cycles for 15 seconds or until alert is set to "none".
func (g *Group) AlertContext(ctx context.Context, new string) error {
	update := State{On: true, Alert: new}
	_, err := g.bridge.SetGroupStateContext(ctx, g.ID, update)
	if err != nil {
		return err
	}
	g.State.Effect = new
	return nil
}

// EnableStreaming enables streaming for the group by setting the Stream Active property to true
func (g *Group) EnableStreaming() error {
	return g.EnableStreamingContext(context.Background())
}

// EnableStreamingContext enables streaming for the group by setting the Stream Active property to true
func (g *Group) EnableStreamingContext(ctx context.Context) error {
	active := true
	update := Group{
		Stream: &Stream{
			ActiveRaw: &active,
		},
	}
	_, err := g.bridge.UpdateGroupContext(ctx, g.ID, update)
	if err != nil {
		return err
	}

	if g.Stream != nil {
		g.Stream.ActiveRaw = &active
		g.Stream.OwnerRaw = &g.bridge.User
	} else {
		g.Stream = &Stream{
			ActiveRaw: &active,
			OwnerRaw:  &g.bridge.User,
		}
	}

	return nil
}

// DisableStreaming disabled streaming for the group by setting the Stream Active property to false
func (g *Group) DisableStreaming() error {
	return g.DisableStreamingContext(context.Background())
}

// DisableStreamingContext disabled streaming for the group by setting the Stream Active property to false
func (g *Group) DisableStreamingContext(ctx context.Context) error {
	active := false
	update := Group{
		Stream: &Stream{
			ActiveRaw: &active,
		},
	}
	_, err := g.bridge.UpdateGroupContext(ctx, g.ID, update)
	if err != nil {
		return err
	}

	if g.Stream != nil {
		g.Stream.ActiveRaw = &active
		g.Stream.OwnerRaw = nil
	}

	return nil
}
