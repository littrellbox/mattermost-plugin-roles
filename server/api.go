package main

import (
	"encoding/json"
	"net/http"
)

type userGetPermissionRequest struct {
	TeamID     string `json:"team_id"`
	UserID     string `json:"user_id"`
	Permission string `json:"permission"`
}

func (p *Plugin) handleUserGetPermission(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the permission
}

type userGetAllPermissionsRequest struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

func (p *Plugin) handleUserGetAllPermissions(w http.ResponseWriter, r *http.Request) {
	var req userGetAllPermissionsRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	http.Error(w, req.UserID+" "+req.TeamID, http.StatusInternalServerError)

	reqinfo, err := json.Marshal(p.getUserPermissions(req.UserID, req.TeamID))

	p.API.LogDebug("roles string json getallpermissions", "reqinfo", reqinfo)

	if err != nil {
		p.API.LogError("failed to convert to json (handleUserGetAllPermissions)", "err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err2 := w.Write(reqinfo); err2 != nil {
		p.API.LogError("failed to write response to handleUserGetAllPermissions", "err", err2.Error())
	}
}

type userGetRolesRequest struct {
	TeamID string `json:"team_id"`
	UserID string `json:"user_id"`
}

func (p *Plugin) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the role names
}

type roleGetAllPermissionsRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
}

func (p *Plugin) handleRoleGetAllPermissions(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the permissions
}

type roleGetPermissionRequest struct {
	TeamID     string `json:"team_id"`
	RoleName   string `json:"role_name"`
	Permission string `json:"permission"`
}

func (p *Plugin) handleRoleGetPermission(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle sending the permission
}

type roleSetPermissionRequest struct {
	TeamID     string `json:"team_id"`
	RoleName   string `json:"role_name"`
	UserID     string `json:"user_id"` //used to verify session token
	Permission string `json:"permission"`
	Value      bool   `json:"value"`
}

func (p *Plugin) handleRoleSetPermission(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle setting the permission
}

type roleCreateRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	UserID   string `json:"user_id"` //used to verify session token
}

func (p *Plugin) handleCreateRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}

type roleDestroyRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	UserID   string `json:"user_id"` //used to verify session token
}

func (p *Plugin) handleDestroyRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}

type roleAddUserToRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	UserID   string `json:"user_id"` //used to verify session token
	TargetID string `json:"target_id"`
}

func (p *Plugin) handleAddUserToRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}

type roleRemoveUserFromRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	UserID   string `json:"user_id"` //used to verify session token
	TargetID string `json:"target_id"`
}

func (p *Plugin) handleRemoveUserFromRole(w http.ResponseWriter, r *http.Request) {
	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	//handle creating the role
}
