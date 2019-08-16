package main

func (p *Plugin) getUserPermission(userid string, permission string, teamID string) (hasPermission bool) {
	//TODO: implement getting user perms
	return p.getDefaultPermssion(permission)
}

//only gets user perms, not role perms
func (p *Plugin) getUserPermissionUO(userID string, permission string, teamID string) (hasPermission bool) {
	//TODO: implement getting user perms
	return p.getDefaultPermssion(permission)
}

func (p *Plugin) getUserPermissions(userID string, teamID string) (permissions map[string]bool) {
	//TODO: implement getting user perms
	m := make(map[string]bool)
	m["send"] = p.getDefaultPermssion("send")
	m["files"] = p.getDefaultPermssion("files")
	m["mute"] = p.getDefaultPermssion("mute")
	m["kick"] = p.getDefaultPermssion("kick")
	m["ban"] = p.getDefaultPermssion("ban")
	return m
}

func (p *Plugin) getUserPermissionsUO(userID string, teamID string) (permissions map[string]bool) {
	//TODO: implement getting user perms
	m := make(map[string]bool)
	m["send"] = p.getDefaultPermssion("send")
	m["files"] = p.getDefaultPermssion("files")
	m["mute"] = p.getDefaultPermssion("mute")
	m["kick"] = p.getDefaultPermssion("kick")
	m["ban"] = p.getDefaultPermssion("ban")
	return m
}

func (p *Plugin) setUserPermission(userID string, teamID string, permission string, value bool) {
	//TODO: implement setting user perms
}

func (p *Plugin) getDefaultPermssion(permission string) (hasPermission bool) {
	//TODO: make less hacky
	switch permission {
	case "send":
		return true
	case "files":
		return true
	default:
		return false
	}
}

func (p *Plugin) getRolePermission(roleName string, teamID string, permission string) (hasPermission bool) {
	//TODO: implement getting role perms
	return p.getDefaultPermssion(permission)
}

func (p *Plugin) getRolePermissions(roleName string, teamID string) (permissions map[string]bool) {
	//TODO: implement getting role perms
	m := make(map[string]bool)
	m["send"] = p.getDefaultPermssion("send")
	m["files"] = p.getDefaultPermssion("files")
	m["mute"] = p.getDefaultPermssion("mute")
	m["kick"] = p.getDefaultPermssion("kick")
	m["ban"] = p.getDefaultPermssion("ban")
	return m
}

func (p *Plugin) setRolePermission(roleName string, teamID string, permission string, value bool) {
	//TODO: implement setting role perms
}

func (p *Plugin) addUserToRole(roleName string, teamID string, userID string) {
	//TODO: implement adding user to role
}

func (p *Plugin) removeUserFromRole(roleName string, teamID string, userID string) {
	//TODO: implement removing user from role
}

func (p *Plugin) createRole(roleName string, teamID string) {
	//TODO: implement creating a role
}

func (p *Plugin) destroyRole(roleName string, teamID string) {
	//TODO: implement destroying roles
}
