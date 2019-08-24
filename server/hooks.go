package main

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

//MessageWillBeUpdated handles checking pinned messages
func (p *Plugin) MessageWillBeUpdated(c *plugin.Context, newPost, oldPost *model.Post) (*model.Post, string) {
	var targetchannel *model.Channel
	var err2 *model.AppError
	targetchannel, err2 = p.API.GetChannel(oldPost.ChannelId)
	if err2 != nil {
		return nil, "An error has occured determining if the file could be uploaded or not. (Error determing target channel)"
	}

	var targetsession *model.Session
	var err *model.AppError

	targetsession, err = p.API.GetSession(c.SessionId)
	if err != nil {
		p.API.SendEphemeralPost(newPost.UserId, &model.Post{
			ChannelId: newPost.ChannelId,
			Message:   "An error has occured determining if the message could be pinned (GetSession:84)",
		})
		return nil, "stop"
	}

	if !p.getUserPermission(targetsession.UserId, "pin", targetchannel.TeamId) {
		if newPost.IsPinned != oldPost.IsPinned {
			p.API.SendEphemeralPost(targetsession.UserId, &model.Post{
				ChannelId: newPost.ChannelId,
				Message:   "You don't have permission to pin messages.",
			})
			return nil, "plugin.message_will_be_posted.dismiss_post"
		}
	}

	return newPost, ""
}

//MessageWillBePosted handles the message will be posted event.
func (p *Plugin) MessageWillBePosted(c *plugin.Context, post *model.Post) (*model.Post, string) {
	var targetchannel *model.Channel
	var err2 *model.AppError
	targetchannel, err2 = p.API.GetChannel(post.ChannelId)
	if err2 != nil {
		return nil, "An error has occured determining if the file could be uploaded or not. (Error determing target channel)"
	}
	var targetuser *model.User
	targetuser, err2 = p.API.GetUser(post.UserId)
	if err2 != nil {
		return nil, "An error has occured determining if the file could be uploaded or not. (Error determing target user)"
	}

	if !p.getUserPermission(post.UserId, "send", targetchannel.TeamId) {
		p.API.SendEphemeralPost(post.UserId, &model.Post{
			ChannelId: post.ChannelId,
			Message:   "You don't have permission to send messages.",
		})
		return nil, "plugin.message_will_be_posted.dismiss_post"
	}

	if !p.getUserPermission(post.UserId, "files", targetchannel.TeamId) {
		if len(post.FileIds) != 0 {
			p.API.SendEphemeralPost(post.UserId, &model.Post{
				ChannelId: post.ChannelId,
				Message:   "You don't have permission to upload files.",
			})
			return nil, "plugin.message_will_be_posted.dismiss_post"
		}
	}

	muteStatus, err3 := p.API.KVGet("lb_roles:" + targetchannel.TeamId + ":" + targetuser.Username + ":mute")

	if err3 != nil {
		return nil, "An error has occured determining if the message could be sent or not. (Error determining mute status)"
	}

	if string(muteStatus) == "true" {
		if len(post.FileIds) != 0 {
			p.API.SendEphemeralPost(post.UserId, &model.Post{
				ChannelId: post.ChannelId,
				Message:   "You've been muted by a team admin.",
			})
			return nil, "plugin.message_will_be_posted.dismiss_post"
		}
	}

	return post, ""
}
