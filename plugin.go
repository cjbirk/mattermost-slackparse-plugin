package main

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// Plugin yadda yadda
type Plugin struct {
	plugin.MattermostPlugin
	Enabled bool
}

// OnActivate is invoked when the plugin is activated.
func (p *Plugin) OnActivate() error {

	return nil
}

// MessageWillBePosted is invoked when a message is posted by a user before it is commited
// to the database.
func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	model.SlackAttachment()
	return nil
}
