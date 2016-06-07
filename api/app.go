package api

import (
	"fmt"

	"github.com/plimble/ace"
)

//App is a struct that represents a Khan API Application
type App struct {
	Port int
	Host string
	App  *ace.Ace
}

//GetDefaultApp returns a new Khan API Application bound to 0.0.0.0:8888
func GetDefaultApp() *App {
	return GetApp("0.0.0.0", 8888)
}

//GetApp returns a new Khan API Application
func GetApp(host string, port int) *App {
	app := &App{
		Host: host,
		Port: port,
	}
	app.Configure()
	return app
}

//Configure instantiates the required dependencies for Khan Api Application
func (app *App) Configure() {
	app.configureApplication()
}

func (app *App) configureApplication() {
	app.App = ace.New()
}

//Url specifies a triple of method, path and request handler
type Url struct {
	Method  string
	Path    string
	Handler ace.HandlerFunc
}

//AddHandlers adds the specified handlers to the route
func (app *App) AddHandlers(urls ...Url) {
	for _, currURL := range urls {
		urls := []ace.HandlerFunc{currURL.Handler}
		app.App.Handle(currURL.Method, currURL.Path, urls)
	}
}

//Start starts listening for web requests at specified host and port
func (app *App) Start() {
	app.App.Run(fmt.Sprintf("%s:%d", app.Host, app.Port))
}
