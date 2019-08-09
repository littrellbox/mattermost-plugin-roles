package main

func (p *Plugin) getUserPermission(userid string, permission string, teamid string) (hasPermission bool) {
	//todo implement getting user perms
	return p.getDefaultPermssion(permission)
}

//only gets user perms, not role perms
func (p *Plugin) getUserPermissionUO(userid string, permission string, teamid string) {

}

func (p *Plugin) getDefaultPermssion(permission string) (hasPermission bool) {
	switch permission {
	case "send":
		return true
	case "files":
		return true
	default:
		return false
	}
}
