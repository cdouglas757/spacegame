package handlers

import (
	"net/http"
	"time"

	"github.com/cdouglas757/spacegame/components"
	"github.com/cdouglas757/spacegame/db"
	"github.com/cdouglas757/spacegame/spaceman"
)

func NewJunkyardHandler(state *spaceman.SpaceManState, store *db.SpaceMenStore) *JunkyardHandler {
	return &JunkyardHandler{playerState: state, spacemenDb: store}
}

type JunkyardHandler struct {
	playerState *spaceman.SpaceManState
	spacemenDb  *db.SpaceMenStore
}

func (handler *JunkyardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sdto := handler.spacemenDb.GetSpaceManDetails(1)
	if r.Method == http.MethodPost {
		handler.Post(w, r, sdto)
		return
	}
	handler.Get(w, r, sdto)
}

func (h *JunkyardHandler) Get(w http.ResponseWriter, r *http.Request, sdto *db.SpaceManDTO) {
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

	components.Layout("Junkyard", sfs, sdto).Render(r.Context(), w)
}

func (h *JunkyardHandler) Post(w http.ResponseWriter, r *http.Request, sdto *db.SpaceManDTO) {

	h.playerState.NextAction = time.Now().Add(time.Second * 30)

	message := "You found some scrap!"

	components.SearchForScrap(false, message).Render(r.Context(), w)
}
