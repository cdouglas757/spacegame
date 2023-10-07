package spaceman

import "time"

type SpaceManState struct {
	NextAction time.Time
}

func NewSpaceManState() *SpaceManState {
	return &SpaceManState{NextAction: time.Time{}}
}
