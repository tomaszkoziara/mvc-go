package views

import (
	"net/http"

	"github.com/tomaszkoziara/mvc-go/html"
)

// IDashboard is the interface for dashboard view.
type IDashboard interface {
	ShowPerson(w http.ResponseWriter, name string, lastName string, magicNumber int) error
	ShowCat(w http.ResponseWriter, name string, favouriteFoods []string) error
}

// Dashboard is the dashboard view.
type Dashboard struct {
	tmpl html.ITemplate
}

// ShowPerson shows dashboard for a person.
func (d *Dashboard) ShowPerson(w http.ResponseWriter, name string, lastName string, magicNumber int) error {
	return d.tmpl.Apply(w, "person-dashboard", map[string]interface{}{
		"name":        name,
		"lastName":    lastName,
		"magicNumber": magicNumber,
	})
}

// ShowCat shows dashboard for a cat.
func (d *Dashboard) ShowCat(w http.ResponseWriter, name string, favouriteFoods []string) error {
	return d.tmpl.Apply(w, "cat-dashboard", map[string]interface{}{
		"name":           name,
		"favouriteFoods": favouriteFoods,
	})
}

// NewDashboardView return a new dashboard view.
func NewDashboardView(tmpl html.ITemplate) *Dashboard {
	return &Dashboard{
		tmpl: tmpl,
	}
}
