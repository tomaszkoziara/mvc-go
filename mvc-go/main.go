package main

import (
	"github.com/tomaszkoziara/mvc-go/html"

	"github.com/tomaszkoziara/mvc-go/handlers"
	"github.com/tomaszkoziara/mvc-go/models"
	"github.com/tomaszkoziara/mvc-go/server"
	"github.com/tomaszkoziara/mvc-go/views"
)

func main() {

	tmpl := html.NewTemplate()
	err := tmpl.Load("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}

	userStore := models.NewUserStore()

	dashboardView := views.NewDashboardView(tmpl)

	userHandler := handlers.NewUserHandler(userStore)
	dashboardHandler := handlers.NewDashboardHandler(dashboardView)

	srv := server.New(userHandler, dashboardHandler)
	srv.Serve()

}
