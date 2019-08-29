package main

import (
	"strings"
)

//max role name length must be 20
func (p *Plugin) getUserPermission(userID string, permission string, teamID string) (hasPermission bool) {
	rolesString, err := p.API.KVGet("lbroles_:" + "teamID[0:9]" + ":" + "userID[0:9]" + ":roles")

	if err != nil {
		p.API.LogError("error adding user to role", "userID", userID)
		return false
	}

	if rolesString == nil {
		p.API.LogInfo("plswork", "rolesString", nil)
		return p.getDefaultPermission(permission)
	}

	if string(rolesString) == "" {
		p.API.LogInfo("plswork", "rolesString", nil)
		return p.getDefaultPermission(permission)
	}

	p.API.LogInfo("plswork", "b", nil)
	rolesStrings := strings.Split(strings.TrimSpace(string(rolesString)), ",")
	permissionToFind := false
	for i := 0; i < len(rolesStrings); i++ {
		permissionCheck := p.getRolePermission(rolesStrings[i], teamID[0:9], permission)
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

func (p *Plugin) getRoleMembers(roleName string, teamID string) (members []string) {
	usersString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":" + roleName[0:19] + ":users")
	if err != nil {
		p.API.LogError("error getting members", "err", err.Error())
		return nil
	}
	return strings.Split(string(usersString), ",")
}

func (p *Plugin) getRolePermission(roleName string, teamID string, permission string) (hasPermission bool) {
	permissionString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":" + roleName[0:19] + ":" + permission)
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

func (p *Plugin) getUserRoles(teamID string, userID string) (roles []string) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":" + userID[0:9] + ":roles")
	if err != nil {
		p.API.LogError("error adding user to role", "err", err.Error())
		return nil
	}
	rolesStrings := strings.Split(strings.TrimSpace(string(rolesString)), ",")
	return rolesStrings

}

func (p *Plugin) getTeamRoles(teamID string) (roles []string) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":roles")
	if err != nil {
		p.API.LogError("error adding user to role", "err", err.Error())
		return nil
	}
	rolesStrings := strings.Split(strings.TrimSpace(string(rolesString)), ",")
	return rolesStrings
}

func (p *Plugin) setRolePermission(roleName string, teamID string, permission string, value bool) {
	//* you may think there is a better way to do this, but there is not.
	//* why in the world i can't just convert 1, 0, int8(1), int8(0), true or false
	//* to binary i will never understand.
	stringToUse := "false"

	if value {
		stringToUse = "true"
	}

	p.API.KVSet("lbroles_:"+teamID[0:9]+":"+roleName[0:19]+":"+permission, []byte(stringToUse))
}

func (p *Plugin) addUserToRole(roleName string, teamID string, userID string) (statusCode int) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":" + userID[0:9] + ":roles")
	rolesUsersString, err2 := p.API.KVGet("lbroles_:" + teamID[0:9] + ":" + roleName + ":users")
	if err != nil {
		p.API.LogError("error adding user to role", "err", err.Error())
		return 1
	}
	if err2 != nil {
		p.API.LogError("error adding user to role", "err2", err2.Error())
		return 1
	}
	newRolesString := roleName
	if rolesString != nil {
		newRolesString = string(rolesString) + "," + roleName
	}
	newRolesString2 := userID
	if rolesUsersString != nil {
		newRolesString2 = string(rolesUsersString) + "," + roleName
	}
	p.API.KVSet("lbroles_:"+teamID[0:9]+":"+roleName[0:19]+":users", []byte(newRolesString2))
	p.API.KVSet("lbroles_:"+teamID[0:9]+":"+userID[0:9]+":roles", []byte(newRolesString))
	return 0
}

func (p *Plugin) removeUserFromRole(roleName string, teamID string, userID string) (statusCode int) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":" + userID[0:9] + ":roles")
	rolesUsersString, err2 := p.API.KVGet("lbroles_:" + teamID[0:9] + ":" + roleName + ":users")
	if err != nil {
		p.API.LogError("error adding user to role", "userID", userID)
		return 1
	}
	if err2 != nil {
		p.API.LogError("error adding user to role", "err2", err2.Error())
		return 1
	}
	if rolesString == nil {
		return 3
	}
	if rolesUsersString == nil {
		return 3
	}
	rolesStrings := strings.Split(strings.TrimSpace(string(rolesString)), ",")
	if stringInSlice(roleName, rolesStrings) {
		rolesStrings = remove(rolesStrings, roleName)
	} else {
		return 2
	}
	rolesUsersStrings := strings.Split(strings.TrimSpace(string(rolesUsersString)), ",")
	if stringInSlice(userID, rolesStrings) {
		rolesUsersStrings = remove(rolesUsersStrings, userID)
	} else {
		return 2
	}

	p.API.KVSet("lbroles_:"+teamID[0:9]+":"+roleName[0:19]+":users", []byte(strings.Join(rolesUsersStrings, ",")))
	p.API.KVSet("lbroles_:"+teamID[0:9]+":"+userID[0:9]+":roles", []byte(strings.Join(rolesStrings, ",")))
	return 0
}

func (p *Plugin) createRole(roleName string, teamID string) (statusCode int) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":roles")
	if err != nil {
		p.API.LogError("error creating role", "teamID", teamID)
		return 1
	}
	newRolesString := roleName[0:19]
	if rolesString != nil {
		newRolesString = string(rolesString) + "," + roleName[0:19]
	}
	p.API.KVSet("lbroles_:"+teamID[0:9]+":roles", []byte(newRolesString))
	return 0
}

func (p *Plugin) destroyRole(roleName string, teamID string) (statusCode int) {
	rolesString, err := p.API.KVGet("lbroles_:" + teamID[0:9] + ":roles")
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

	p.API.KVSet("lbroles_:"+teamID[0:9]+":roles", []byte(strings.Join(rolesStrings, ",")))
	return 0
}
