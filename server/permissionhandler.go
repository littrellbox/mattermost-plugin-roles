package main

func (p *Plugin) getUserPermission(userid string, permission string, teamID string) (hasPermission bool) {
	//TODO: implement getting user perms
	return p.getDefaultPermssion(permission)
}

//only gets user perms, not role perms
func (p *Plugin) getUserPermissionUO(userID string, permission string, teamID string) {

}

func (p *Plugin) getUserPermissions(userID string, teamID string) {

}

func (p *Plugin) getUserPermissionsUO(userID string, teamID string) {

}

func (p *Plugin) setUserPermission(userID string, teamID string, permission string, value bool) {

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

func (p *Plugin) getRolePermission(roleName string, teamID string, permission string) {

}

func (p *Plugin) setRolePermission(roleName string, teamID string, permission string, value bool) {

}

func (p *Plugin) addUserToRole(roleName string, teamID string, userID string) {

}

func (p *Plugin) removeUserFromRole(roleName string, teamID string, userID string) {

}

func (p *Plugin) createRole(roleName string, teamID string) {

}

func (p *Plugin) destroyRole(roleName string, teamID string) {

}
