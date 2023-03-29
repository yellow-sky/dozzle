package web

import (
	"net/http"
	"time"
)

func (h *handler) restartContainer(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	container, err := h.clientFromRequest(r).FindContainer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	timeout := 3 //TODO: move to config or args
	h.clientFromRequest(r).ContainerRestart(r.Context(), container.ID, &timeout)
	w.WriteHeader(200)
}

func (h *handler) pullContainerImage(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	container, err := h.clientFromRequest(r).FindContainer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timeout := 3 * time.Second //TODO: move to config or args
	logBytes, err := h.clientFromRequest(r).ContainerImagePull(r.Context(), container.ID, &timeout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(200)
		w.Write(logBytes)
		return
	}
}
