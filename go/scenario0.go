package main

import (
	"log/slog"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/api"
)

func scenario0(apiClient *api.Client, token string) error {
	// Auth w/ API token
	err := apiClient.Auth(token)
	if err != nil {
		return err
	}

	// Creation / deletion of a team
	err = apiClient.CreateTeam(teamId)
	if err != nil {
		return err
	}
	defer cleanup("delete team "+teamId, func() error {
		return apiClient.DeleteTeam(teamId, true)
	})

	// Creation of a project
	projectId, err := apiClient.CreateProject(projectName)
	if err != nil {
		return err
	}
	defer cleanup("delete project "+projectId, func() error {
		return apiClient.DeleteProject(projectId)
	})

	// Grant access
	return apiClient.AddProjectAccess(projectId, teamId)
}

func cleanup(callDescription string, apiCall func() error) {
	err := apiCall()
	if err != nil {
		slog.Warn("unable to " + callDescription)
		slog.Warn(err.Error())
	}
}
