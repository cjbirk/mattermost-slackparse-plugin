package main

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// Plugin yadda yadda
type Plugin struct {
	plugin.MattermostPlugin
	api     plugin.API
	Enabled bool
}

// OnActivate is invoked when the plugin is activated.
func (p *Plugin) OnActivate(api plugin.API) error {
	p.api = api
	p.API.LogDebug("I'M JUST A POOR BOY")

	if err := p.OnConfigurationChange(); err != nil {
		return err
	}

	return nil
}

// FilterPost our post to process via ParseSlackAttachment
func (p *Plugin) FilterPost(post *model.Post) (*model.Post, string) {
	p.API.LogDebug("NOBODY LOVES ME")
	if post.Type == model.POST_SLACK_ATTACHMENT {
		p.API.LogDebug("HE'S JUST A POOR BOY")
		attachments, _ := post.Props["attachments"].([]*model.SlackAttachment)
		//var newAttachment *model.SlackAttachment
		//newAttachment = attachments[0]
		//for attachment := range attachments {
		//	if attachment
		//}
		p.API.LogDebug("FROM A POOR", "FAMILY::::", attachments)
		model.ParseSlackAttachment(post, []*model.SlackAttachment{attachments})
	}

	return post, ""
}

// MessageWillBePosted is invoked when a message is posted by a user before it is commited
// to the database.
func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	p.API.LogDebug("SPARE HIM HIS LIFE OF THIS MONSTROSITY")
	return p.FilterPost(post)
}

func main() {
	plugin.ClientMain(&Plugin{})
}
