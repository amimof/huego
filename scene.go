package huego

import "context"

// Scene represents a bridge scene https://developers.meethue.com/documentation/scenes-api
type Scene struct {
	Name            string        `json:"name,omitempty"`
	Type            string        `json:"type,omitempty"`
	Group           string        `json:"group,omitempty"`
	Lights          []string      `json:"lights,omitempty"`
	Owner           string        `json:"owner,omitempty"`
	Recycle         bool          `json:"recycle"`
	Locked          bool          `json:"locked,omitempty"`
	AppData         interface{}   `json:"appdata,omitempty"`
	Picture         string        `json:"picture,omitempty"`
	LastUpdated     string        `json:"lastupdated,omitempty"`
	Version         int           `json:"version,omitempty"`
	StoreSceneState bool          `json:"storescenestate,omitempty"`
	LightStates     map[int]State `json:"lightstates,omitempty"`
	ID              string        `json:"-"`
	bridge          *Bridge
}

// Recall will recall the scene in the group identified by id
func (s *Scene) Recall(id int) error {
	return s.RecallContext(context.Background(), id)
}

// RecallContext will recall the scene in the group identified by id
func (s *Scene) RecallContext(ctx context.Context, id int) error {
	_, err := s.bridge.RecallSceneContext(ctx, s.ID, id)
	if err != nil {
		return err
	}
	return nil
}
