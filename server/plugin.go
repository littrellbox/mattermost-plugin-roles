package main

import (
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path; path {
	case "/api/v1/roles/user/getpermission":
		p.handleUserGetPermission(w, r)
	case "/api/v1/roles/user/setpermission":
		p.handleUserSetPermission(w, r)
	case "/api/v1/roles/user/getallpermissions":
		p.handleUserGetAllPermissions(w, r)
	case "/api/v1/roles/user/getroles":
		p.handleGetRoles(w, r)
	case "/api/v1/roles/role/getallpermissions":
		p.handleRoleGetAllPermissions(w, r)
	case "/api/v1/roles/role/getpermission":
		p.handleRoleGetPermission(w, r)
	case "/api/v1/roles/role/setpermission":
		p.handleRoleSetPermission(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	//var targetchannel *model.Channel
	//var err2 *model.AppError
	//targetchannel, err2 = p.API.GetChannel(post.ChannelId)
	//if err2 != nil {
	//	return nil, "An error has occured determining if the file could be uploaded or not. (Error determing target channel)"
	//}
	//var targetuser *model.User
	//targetuser, err2 = p.API.GetUser(post.UserId)
	//if err2 != nil {
	//	return nil, "An error has occured determining if the file could be uploaded or not. (Error determing target user)"
	//}
	//targetchannel.TeamId

	return post, ""
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
