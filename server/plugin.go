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

//OnActivate Activates the plugin
func (p *Plugin) OnActivate() error {
	p.API.RegisterCommand(&model.Command{
		Trigger:          "muteuser",
		AutoComplete:     false,
		AutoCompleteDesc: "Moderation tool. If you can see this, something's broken.",
		DisplayName:      "muteuser test",
	})

	return nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path; path {
	case "/api/v1/roles/user/getpermission":
		p.handleUserGetPermission(w, r)
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
	case "/api/v1/roles/role/create":
		p.handleCreateRole(w, r)
	case "/api/v1/roles/role/destroy":
		p.handleDestroyRole(w, r)
	case "/api/v1/roles/role/adduser":
		p.handleAddUserToRole(w, r)
	case "/api/v1/roles/role/removeuser":
		p.handleRemoveUserFromRole(w, r)
	default:
		http.NotFound(w, r)
	}
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
