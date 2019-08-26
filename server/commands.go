package main

import (
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	argumentArray := strings.Split(args.Command, " ")

	if argumentArray[0] == "muteuser" {
		if p.getUserPermission(args.UserId, "mute", args.TeamId) {
			if len(argumentArray) == 2 {
				var targetuser *model.User
				var err2 *model.AppError
				targetuser, err2 = p.API.GetUserByUsername(argumentArray[2])

				if err2 != nil {
					return &model.CommandResponse{
						ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
						Text:         "An error has occured finding that user. (Try checking the username)",
					}, nil
				}

				p.API.KVSet("lbrm:"+args.TeamId[0:9]+":"+targetuser.Id[0:15]+":muted", []byte("true"))
				return &model.CommandResponse{
					ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
					Text:         "Muted user " + argumentArray[1] + ".",
				}, nil
			}
			return &model.CommandResponse{
				ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
				Text:         "You need to specify a username.",
			}, nil
		}

		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         "You don't have permission to do that.",
		}, nil
	}

	if argumentArray[0] == "kick" {
		if p.getUserPermission(args.UserId, "kick", args.TeamId) {
			if len(argumentArray) == 2 {
				var targetuser *model.User
				var err2 *model.AppError
				targetuser, err2 = p.API.GetUserByUsername(argumentArray[2])

				if err2 != nil {
					return &model.CommandResponse{
						ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
						Text:         "An error has occured finding that user. (Try checking the username)",
					}, nil
				}

				p.API.DeleteTeamMember(args.TeamId, targetuser.Id, args.UserId)
				return &model.CommandResponse{
					ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
					Text:         "Kicked user " + argumentArray[1] + ".",
				}, nil
			}
			return &model.CommandResponse{
				ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
				Text:         "You need to specify a username.",
			}, nil
		}

		return &model.CommandResponse{
			ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
			Text:         "You don't have permission to do that.",
		}, nil
	}

	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         "That command doesn't exist.",
	}, nil
}
