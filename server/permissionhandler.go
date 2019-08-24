package main

import "strings"

func (p *Plugin) getUserPermission(userID string, permission string, teamID string) (hasPermission bool) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID + ":" + userID + ":roles")
	if err != nil {
		p.API.LogError("error adding user to role", "userID", userID)
		return false
	}
	if rolesString == nil || string(rolesString) == "" {
		return p.getDefaultPermission(permission)
	}
	rolesStrings := strings.Split(strings.TrimSpace(string(rolesString)), ",")
	permissionToFind := false
	for i := 0; i < len(rolesStrings); i++ {
		permissionCheck := p.getRolePermission(rolesStrings[i], teamID, permission)
		if permissionCheck == true {
			permissionToFind = true
			break //the user has the permission, stop
		}
	}
	return permissionToFind
}

func (p *Plugin) getUserPermissions(userID string, teamID string) (permissions map[string]bool) {
	m := make(map[string]bool)
	m["send"] = p.getUserPermission(userID, "send", teamID)
	m["files"] = p.getUserPermission(userID, "files", teamID)
	m["mute"] = p.getUserPermission(userID, "mute", teamID)
	m["kick"] = p.getUserPermission(userID, "kick", teamID)
	m["ban"] = p.getUserPermission(userID, "ban", teamID)
	m["pin"] = p.getUserPermission(userID, "pin", teamID)
	return m
}

func (p *Plugin) getDefaultPermission(permission string) (hasPermission bool) {
	switch permission {
	case "send":
		return true
	case "files":
		return true
	default:
		return false
	}
}

//TODO: Add team role user list variable

func (p *Plugin) getRolePermission(roleName string, teamID string, permission string) (hasPermission bool) {
	permissionString, err := p.API.KVGet("lbroles_:" + teamID + ":" + roleName + ":" + permission)
	if err != nil {
		p.API.LogError("error getting rolePermission", "roleName", roleName)
		return false
	}
	if permissionString == nil {
		return p.getDefaultPermission(permission)
	}
	if string(permissionString) == "true" {
		return true
	}
	return false
}

func (p *Plugin) getRolePermissions(roleName string, teamID string) (permissions map[string]bool) {
	m := make(map[string]bool)
	m["send"] = p.getRolePermission(roleName, teamID, "send")
	m["files"] = p.getRolePermission(roleName, teamID, "files")
	m["mute"] = p.getRolePermission(roleName, teamID, "mute")
	m["kick"] = p.getRolePermission(roleName, teamID, "kick")
	m["ban"] = p.getRolePermission(roleName, teamID, "ban")
	m["pin"] = p.getRolePermission(roleName, teamID, "pin")
	return m
}

func (p *Plugin) setRolePermission(roleName string, teamID string, permission string, value bool) {
	//* you may think there is a better way to do this, but there is not.
	//* why in the world i can't just convert 1, 0, int8(1), int8(0), true or false
	//* to binary i will never understand.
	stringToUse := "false"

	if value {
		stringToUse = "true"
	}

	p.API.KVSet("lbroles_:"+teamID+":"+roleName+":"+permission, []byte(stringToUse))
}

func (p *Plugin) addUserToRole(roleName string, teamID string, userID string) (statusCode int) {
	//TODO: Add team role user list variable
	rolesString, err := p.API.KVGet("lbroles_:" + teamID + ":" + userID + ":roles")
	if err != nil {
		p.API.LogError("error adding user to role", "userID", userID)
		return 1
	}
	newRolesString := roleName
	if rolesString != nil {
		newRolesString = string(rolesString) + "," + roleName
	}
	p.API.KVSet("lbroles_:"+teamID+":"+userID+":roles", []byte(newRolesString))
	return 0
}

func (p *Plugin) removeUserFromRole(roleName string, teamID string, userID string) (statusCode int) {
	//TODO: Add team role user list variable
	rolesString, err := p.API.KVGet("lbroles_:" + teamID + ":" + userID + ":roles")
	if err != nil {
		p.API.LogError("error adding user to role", "userID", userID)
		return 1
	}
	if rolesString == nil {
		return 3
	}
	rolesStrings := strings.Split(strings.TrimSpace(string(rolesString)), ",")
	if stringInSlice(roleName, rolesStrings) {
		rolesStrings = remove(rolesStrings, roleName)
	} else {
		return 2
	}

	p.API.KVSet("lbroles_:"+teamID+":"+userID+":roles", []byte(strings.Join(rolesStrings, ",")))
	return 0
}

func (p *Plugin) createRole(roleName string, teamID string) (statusCode int) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID + ":roles")
	if err != nil {
		p.API.LogError("error creating role", "teamID", teamID)
		return 1
	}
	newRolesString := roleName
	if rolesString != nil {
		newRolesString = string(rolesString) + "," + roleName
	}
	p.API.KVSet("lbroles_:"+teamID+":roles", []byte(newRolesString))
	return 0
}

func (p *Plugin) destroyRole(roleName string, teamID string) (statusCode int) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID + ":roles")
	if err != nil {
		p.API.LogError("error destroying role", "teamID", teamID)
		return 1
	}
	rolesStrings := strings.Split(strings.TrimSpace(string(rolesString)), ",")
	if stringInSlice(roleName, rolesStrings) {
		rolesStrings = remove(rolesStrings, roleName)
	} else {
		return 2
	}

	p.API.KVSet("lbroles_:"+teamID+":roles", []byte(strings.Join(rolesStrings, ",")))
	return 0
}
