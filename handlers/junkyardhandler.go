package handlers

import (
	"net/http"
	"time"

	"github.com/cdouglas757/spacegame/components"
	"github.com/cdouglas757/spacegame/spaceman"
)

func NewJunkyardHandler(state *spaceman.SpaceManState) *JunkyardHandler {
	return &JunkyardHandler{playerState: state}
}

type JunkyardHandler struct {
	playerState *spaceman.SpaceManState
}

func (handler *JunkyardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handler.Post(w, r)
		return
	}
	handler.Get(w, r)
}

func (h *JunkyardHandler) Get(w http.ResponseWriter, r *http.Request) {
	canSearch := false

	if h.playerState.NextAction.IsZero() || time.Now().After(h.playerState.NextAction) {
		canSearch = true
	}

	message := ""

	if canSearch {
		message = "Would you like to search the junkyard?"
	} else {
		minutesLeft := time.Until(h.playerState.NextAction)
		minutesLeft -= minutesLeft % time.Second
		message = "You cannot search again yet! You can search again in " + minutesLeft.String()
	}

	sfs := components.SearchForScrap(canSearch, message)

	components.Layout("Junkyard", sfs).Render(r.Context(), w)
}

func (h *JunkyardHandler) Post(w http.ResponseWriter, r *http.Request) {

	h.playerState.NextAction = time.Now().Add(time.Second * 30)

	message := "You found some scrap!"

	components.SearchForScrap(false, message).Render(r.Context(), w)
}
