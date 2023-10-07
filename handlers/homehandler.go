package handlers

import (
	"net/http"

	"github.com/cdouglas757/spacegame/components"
	"github.com/cdouglas757/spacegame/db"
)

func NewHomeHandler(store *db.SpaceMenStore) *HomeHandler {
	return &HomeHandler{spacemenDb: store}
}

type HomeHandler struct {
	spacemenDb *db.SpaceMenStore
}

func (handler *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sdto := handler.spacemenDb.GetSpaceManDetails(1)

	handler.Get(w, r, sdto)
}

func (h *HomeHandler) Get(w http.ResponseWriter, r *http.Request, sdto *db.SpaceManDTO) {

	home := components.Home()

	components.Layout("Home", home, sdto).Render(r.Context(), w)
}
