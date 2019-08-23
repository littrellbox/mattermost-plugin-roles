package main

import (
	"net/http"
)

type roleRequestWithAuth struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	UserID   string `json:"user_id"` //used to verify session token
}

type roleRequestNoAuth struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
}

type roleRequestWithAuthWithTarget struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	UserID   string `json:"user_id"` //used to verify session token
	TargetID string `json:"target_id"`
}

type userRequest struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

func (p *Plugin) handleUserGetPermission(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the permission
}

func (p *Plugin) handleUserGetAllPermissions(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the permissions
}

func (p *Plugin) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the role names
}

func (p *Plugin) handleGetFullPermissions(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the role names
}

func (p *Plugin) handleRoleGetAllPermissions(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the permissions
}

func (p *Plugin) handleRoleGetPermission(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the permission
}

func (p *Plugin) handleRoleSetPermission(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle setting the permission
}

func (p *Plugin) handleCreateRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}

func (p *Plugin) handleDestroyRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}

func (p *Plugin) handleAddUserToRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}

func (p *Plugin) handleRemoveUserFromRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}
