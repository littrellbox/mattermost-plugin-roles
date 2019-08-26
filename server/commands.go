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
