package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/graphql"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

type Client struct {
	GraphQLClient  graphql.Client
	Endpoint       string
	OperationsPath string
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

func (client Client) readOperationText(operationName string) (string, error) {
	path := client.OperationsPath + "/" + operationName + ".gql"
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", lib.WrapError(fmt.Sprintf("unable to read operation file %s", path), err)
	}
	return string(bytes), nil
}

func (client Client) Auth(token string) error {
	query, err := client.readOperationText("auth")
	if err != nil {
		return err
	}
	variables := map[string]any{
		"token": token,
	}
	return client.sendRequestDiscardingData("auth", query, variables)
}

func (client Client) CreateTeam(teamId string) error {
	query, err := client.readOperationText("create_team")
	if err != nil {
		return err
	}
	variables := map[string]any{
		"teamId": teamId,
	}
	return client.sendRequestDiscardingData(
		fmt.Sprintf("create team '%s'", teamId),
		query,
		variables,
	)
}

func (client Client) DeleteTeam(teamId string, force bool) error {
	query, err := client.readOperationText("delete_team")
	if err != nil {
		return err
	}
	variables := map[string]any{
		"teamId": teamId,
		"force":  force,
	}
	return client.sendRequestDiscardingData(
		fmt.Sprintf("delete team '%s'", teamId),
		query,
		variables,
	)
}

func (client Client) CreateProject(projectName string) (id string, err error) {
	query, err := client.readOperationText("create_project")
	if err != nil {
		return "", err
	}
	variables := map[string]any{
		"projectName": projectName,
	}
	rawData, err := client.sendRequest(
		fmt.Sprintf("create project with name '%s'", projectName),
		query,
		variables,
	)
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
	query, err := client.readOperationText("delete_project")
	if err != nil {
		return err
	}
	variables := map[string]any{
		"projectId": projectId,
	}
	return client.sendRequestDiscardingData(
		fmt.Sprintf("delete project with id '%s'", projectId),
		query,
		variables,
	)
}

// subjectIds: Ids of teams or individual users
func (client Client) AddProjectAccess(projectId string, subjectIds ...string) error {
	query, err := client.readOperationText("add_project_permissions")
	if err != nil {
		return err
	}
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
		query,
		variables,
	)
}
