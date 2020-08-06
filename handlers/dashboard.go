package handlers

import (
	"math/rand"
	"net/http"

	"github.com/tomaszkoziara/mvc-go/views"
)

// IDashboard is the interface for dashboard handler.
type IDashboard interface {
	Show(w http.ResponseWriter, r *http.Request)
}

// Dashboard is the dashboard handler for all methods.
type Dashboard struct {
	view views.IDashboard
}

// Show shows the dashboard.
func (d *Dashboard) Show(w http.ResponseWriter, r *http.Request) {
	var err error
	if rand.Int()%2 == 0 {
		err = d.view.ShowPerson(w, "Mario", "Rossi", 42)
	} else {
		err = d.view.ShowCat(w, "Corazon", []string{"fish", "milk", "birds"})
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// NewDashboardHandler return a new dashboard handler.
func NewDashboardHandler(view views.IDashboard) IDashboard {
	return &Dashboard{
		view: view,
	}
}
