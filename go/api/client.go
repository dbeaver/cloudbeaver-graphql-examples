package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/graphql"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

// Queries and mutations
const (
	authQuery = `
query authLogin($token: String!) {
    authLogin(provider: "token", credentials: { token: $token }) {
        userTokens {
            userId
        }
        authStatus
    }
}
`

	createTeamQuery = `
query createTeam($teamId: ID!) {
  createTeam(teamId: $teamId) {
    teamId
  }
}
`

	deleteTeamQuery = `
query deleteTeam($teamId: ID!, $force: Boolean) {
  deleteTeam(teamId: $teamId, force: $force)
}
`

	createProjectMutation = `
mutation RmCreateProject($projectName: String!) {
    rmCreateProject(projectName: $projectName) {
        id
    }
}
`

	deleteProjectMutation = `
mutation RmDeleteProject($projectId: ID!) {
    rmDeleteProject(projectId: $projectId)
}
`
	addProjectPermissionsMutation = `
mutation addProjectsPermissions($projectIds: [ID!]!, $subjectIds: [ID!]!, $permissions: [String!]!) {
    rmAddProjectsPermissions(
        projectIds: $projectIds
        subjectIds: $subjectIds
        permissions: $permissions
    )
}
`
)

type Client struct {
	GraphQLClient graphql.Client
	Endpoint      string
}

func (client Client) sendRequest(operationName, query string, variables graphql.Object) (json.RawMessage, error) {
	request := graphql.Request{Query: query, Variables: variables}
	slog.Info("--> GraphQL call [" + operationName + "]")
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		slog.Info(fmt.Sprintf("<-- Call [%s] finished (%s)", operationName, duration))
	}()
	response, err := client.GraphQLClient.Execute(client.Endpoint, request)
	if err != nil {
		return json.RawMessage{}, lib.WrapError("error while sending an api request", err)
	}
	if response.Errors != nil {
		err = errors.New("error recieved when executing an API call: \n" + string(response.Errors))
	}
	return response.Data, err
}

func (client Client) sendRequestDiscardingData(operationName, query string, variables graphql.Object) error {
	_, err := client.sendRequest(operationName, query, variables)
	return err
}

func (client Client) Auth(token string) error {
	variables := map[string]any{
		"token": token,
	}
	return client.sendRequestDiscardingData("auth", authQuery, variables)
}

func (client Client) CreateTeam(teamId string) error {
	variables := map[string]any{
		"teamId": teamId,
	}
	return client.sendRequestDiscardingData(
		fmt.Sprintf("create team '%s'", teamId),
		createTeamQuery,
		variables)
}

func (client Client) DeleteTeam(teamId string, force bool) error {
	variables := map[string]any{
		"teamId": teamId,
		"force":  force,
	}
	return client.sendRequestDiscardingData(
		fmt.Sprintf("delete team '%s'", teamId),
		deleteTeamQuery,
		variables)
}

func (client Client) CreateProject(projectName string) (id string, err error) {
	variables := map[string]any{
		"projectName": projectName,
	}
	rawData, err := client.sendRequest(
		fmt.Sprintf("create project with name '%s'", projectName),
		createProjectMutation,
		variables)
	if err != nil {
		return id, err
	}
	var rmCreateProjectResponse RmCreateProjectResponse
	err = json.Unmarshal(rawData, &rmCreateProjectResponse)
	if err != nil {
		return id, lib.WrapError("unable to unmarshal rmCreateProjectResponse", err)
	}
	return rmCreateProjectResponse.RmCreateProject.Id, err
}

func (client Client) DeleteProject(projectId string) error {
	variables := map[string]any{
		"projectId": projectId,
	}
	rawData, err := client.sendRequest(
		fmt.Sprintf("delete project with id '%s'", projectId),
		deleteProjectMutation,
		variables)
	if err != nil {
		return err
	}
	slog.Debug(string(rawData))
	return nil
}

// subjectIds: Ids of teams or individual users
func (client Client) AddProjectAccess(projectId string, subjectIds ...string) error {
	variables := map[string]any{
		"projectIds": [1]string{projectId},
		"subjectIds": subjectIds,
		"permissions": [2]string{
			"project-datasource-view",
			"project-resource-view",
		},
	}
	return client.sendRequestDiscardingData(
		fmt.Sprintf("grant subjects %s access to project with id '%s'", subjectIds, projectId),
		addProjectPermissionsMutation,
		variables,
	)
}
