package main

import (
	"io"
	"log"
	"os"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

var (
	Info *log.Logger
)

// Plugin yadda yadda
type Plugin struct {
	plugin.MattermostPlugin
	Enabled bool
}

// Init logging anus
func Init(infoHandle io.Writer) {

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// OnActivate is invoked when the plugin is activated.
func (p *Plugin) OnActivate() error {
	return nil
}

// FilterPost our post to process via ParseSlackAttachment
func (p *Plugin) FilterPost(post *model.Post) (*model.Post, string) {
	Info.Println("THIS IS THE FILTERPOST FUNCTION U DINGUS")
	if post.Type == model.POST_SLACK_ATTACHMENT {
		// 	if err := json.NewDecoder(r.Body).Decode(&webhook); err != nil {
		// 		http.Error(w, err.Error(), http.StatusBadRequest)
		// 	} else if attachment, err := webhook.SlackAttachment(); err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	} else if attachment == nil {
		// 		return
		// 	}
		// post.Props is the field that we're going to parse to get the attach-
		// ment thing - lamer
		attachments, _ := post.Props["attachments"].([]*model.SlackAttachment)
		// do something
		var newAttachment *model.SlackAttachment
		newAttachment = attachments[0]

		model.ParseSlackAttachment(post, []*model.SlackAttachment{newAttachment})
	}

	return post, ""
}

// MessageWillBePosted is invoked when a message is posted by a user before it is commited
// to the database.
func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	Info.Println("HEY LOOK AT THESE LOGS IN THE MESSAGEWILLBEPOSTED GABBAGE")
	return p.FilterPost(post)
}

func main() {
	Init(os.Stdout)
	plugin.ClientMain(&Plugin{})
}
