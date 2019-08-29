package main

import (
	"encoding/json"
	"net/http"

	"github.com/mattermost/mattermost-server/model"
)

type userGetPermissionRequest struct {
	TeamID     string `json:"team_id"`
	UserID     string `json:"user_id"`
	Permission string `json:"permission"`
}

func (p *Plugin) handleUserGetPermission(w http.ResponseWriter, r *http.Request) {
	var req userGetPermissionRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	reqinfo, err := json.Marshal(p.getUserPermission(req.UserID, req.Permission, req.TeamID))

	if err != nil {
		p.API.LogError("failed to convert to json (handleUserGetPermission)", "err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err2 := w.Write(reqinfo); err2 != nil {
		p.API.LogError("failed to write response to handleUserGetAllPermissions", "err", err2.Error())
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
}

func (p *Plugin) handleGetRoles(w http.ResponseWriter, r *http.Request) {
	var req userGetRolesRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	reqinfo, err := json.Marshal(p.getTeamRoles(req.TeamID))

	if err != nil {
		p.API.LogError("failed to convert to json (handleGetRoles)", "err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err2 := w.Write(reqinfo); err2 != nil {
		p.API.LogError("failed to write response to handleGetRoles", "err", err2.Error())
	}
}

type roleGetUsersRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
}

func (p *Plugin) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	var req roleGetUsersRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	reqinfo, err := json.Marshal(p.getRoleMembers(req.RoleName, req.TeamID))

	if err != nil {
		p.API.LogError("failed to convert to json (handleGetUsers)", "err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err2 := w.Write(reqinfo); err2 != nil {
		p.API.LogError("failed to write response to handleGetUsers", "err", err2.Error())
	}
}

type roleGetAllPermissionsRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
}

func (p *Plugin) handleRoleGetAllPermissions(w http.ResponseWriter, r *http.Request) {
	var req roleGetAllPermissionsRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	reqinfo, err := json.Marshal(p.getRolePermissions(req.RoleName, req.TeamID))

	if err != nil {
		p.API.LogError("failed to convert to json (handleRoleGetAllPermissions)", "err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err2 := w.Write(reqinfo); err2 != nil {
		p.API.LogError("failed to write response to handleRoleGetAllPermissions", "err", err2.Error())
	}
}

type roleGetPermissionRequest struct {
	TeamID     string `json:"team_id"`
	RoleName   string `json:"role_name"`
	Permission string `json:"permission"`
}

func (p *Plugin) handleRoleGetPermission(w http.ResponseWriter, r *http.Request) {
	var req roleGetPermissionRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	reqinfo, err := json.Marshal(p.getRolePermission(req.RoleName, req.TeamID, req.Permission))

	if err != nil {
		p.API.LogError("failed to convert to json (handleRoleGetPermission)", "err", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err2 := w.Write(reqinfo); err2 != nil {
		p.API.LogError("failed to write response to handleRoleGetPermission", "err", err2.Error())
	}
}

type roleSetPermissionRequest struct {
	TeamID     string `json:"team_id"`
	RoleName   string `json:"role_name"`
	Permission string `json:"permission"`
	Value      bool   `json:"value"`
}

func (p *Plugin) handleRoleSetPermission(w http.ResponseWriter, r *http.Request) {
	var req roleSetPermissionRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	teamMember, err := p.API.GetTeamMember(req.TeamID, userID)

	if err != nil {
		p.API.LogError("failed to convert to json (handleRoleGetPermission)", "err", err.Error())
	}

	if !stringInSlice(model.TEAM_ADMIN_ROLE_ID, teamMember.GetRoles()) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	p.setRolePermission(req.RoleName, req.TeamID, req.Permission, req.Value)

	if _, err2 := w.Write([]byte("done")); err2 != nil {
		p.API.LogError("failed to write response to handleUserGetAllPermissions", "err", err2.Error())
	}
}

type roleCreateRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
}

func (p *Plugin) handleCreateRole(w http.ResponseWriter, r *http.Request) {
	var req roleSetPermissionRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	teamMember, err := p.API.GetTeamMember(req.TeamID, userID)

	if err != nil {
		p.API.LogError("failed to get team member", "err", err.Error())
	}

	if !stringInSlice(model.TEAM_ADMIN_ROLE_ID, teamMember.GetRoles()) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	intcode := p.createRole(req.RoleName, req.TeamID)

	if intcode == 1 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if _, err2 := w.Write([]byte("done")); err2 != nil {
		p.API.LogError("failed to write response to handleUserGetAllPermissions", "err", err2.Error())
	}
}

type roleDestroyRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
}

func (p *Plugin) handleDestroyRole(w http.ResponseWriter, r *http.Request) {
	var req roleSetPermissionRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	teamMember, err := p.API.GetTeamMember(req.TeamID, userID)

	if err != nil {
		p.API.LogError("failed to get team member", "err", err.Error())
	}

	if !stringInSlice(model.TEAM_ADMIN_ROLE_ID, teamMember.GetRoles()) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	intcode := p.destroyRole(req.RoleName, req.TeamID)

	if intcode == 1 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if intcode == 2 {
		http.Error(w, "Non-existant role", http.StatusBadRequest)
	}

	if _, err2 := w.Write([]byte("done")); err2 != nil {
		p.API.LogError("failed to write response to handleUserGetAllPermissions", "err", err2.Error())
	}
}

type roleAddUserToRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	TargetID string `json:"target_id"`
}

func (p *Plugin) handleAddUserToRole(w http.ResponseWriter, r *http.Request) {
	var req roleAddUserToRoleRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	teamMember, err := p.API.GetTeamMember(req.TeamID, userID)

	if err != nil {
		p.API.LogError("failed to get team member", "err", err.Error())
	}

	if !stringInSlice(model.TEAM_ADMIN_ROLE_ID, teamMember.GetRoles()) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	intcode := p.addUserToRole(req.RoleName, req.TeamID, req.TargetID)

	if intcode == 1 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if intcode == 2 {
		http.Error(w, "Non-existant role", http.StatusBadRequest)
	}

	if intcode == 3 {
		http.Error(w, "Non-existant role", http.StatusBadRequest)
	}

	if _, err2 := w.Write([]byte("done")); err2 != nil {
		p.API.LogError("failed to write response to handleUserGetAllPermissions", "err", err2.Error())
	}
}

type roleRemoveUserFromRoleRequest struct {
	TeamID   string `json:"team_id"`
	RoleName string `json:"role_name"`
	TargetID string `json:"target_id"`
}

func (p *Plugin) handleRemoveUserFromRole(w http.ResponseWriter, r *http.Request) {
	var req roleAddUserToRoleRequest

	//body is empty
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	userID := r.Header.Get("Mattermost-User-Id")

	if userID == "" {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	teamMember, err := p.API.GetTeamMember(req.TeamID, userID)

	if err != nil {
		p.API.LogError("failed to get team member", "err", err.Error())
	}

	if !stringInSlice(model.TEAM_ADMIN_ROLE_ID, teamMember.GetRoles()) {
		http.Error(w, "Not authorized", http.StatusUnauthorized)
		return
	}

	intcode := p.addUserToRole(req.RoleName, req.TeamID, req.TargetID)

	if intcode == 1 {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	if intcode == 2 {
		http.Error(w, "Non-existant role", http.StatusBadRequest)
	}

	if intcode == 3 {
		http.Error(w, "Non-existant role", http.StatusBadRequest)
	}

	if _, err2 := w.Write([]byte("done")); err2 != nil {
		p.API.LogError("failed to write response to handleUserGetAllPermissions", "err", err2.Error())
	}
}
