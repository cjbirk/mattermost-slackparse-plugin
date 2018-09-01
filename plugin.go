package main

import (
	"fmt"
	"net/http"

	"github.com/mattermost/mattermost-server/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
	Enabled bool
}

// OnActivate is invoked when the plugin is activated.
func (p *Plugin) OnActivate() error {

	return nil
}

// Test
type HelloWorldPlugin struct {
	plugin.MattermostPlugin
}

func (p *HelloWorldPlugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

// This example demonstrates a plugin that handles HTTP requests which respond by greeting the
// world.
func main() {
	plugin.ClientMain(&HelloWorldPlugin{})
}
