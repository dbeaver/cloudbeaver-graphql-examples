package api

import (
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
)

type Client struct {
	GraphQLClient graphql.Client
	Endpoint      string
}

func (client Client) sendRequest(operationName, query string, variables map[string]any) error {
	request := graphql.Request{Query: query, Variables: variables}
	slog.Info("--> GraphQL call [" + operationName + "]")
	start := time.Now()
	defer func() {
		duration := time.Since(start)
		slog.Info(fmt.Sprintf("<-- Call [%s] finished (%s)", operationName, duration))
	}()
	response, err := client.GraphQLClient.Execute(client.Endpoint, request)
	if err != nil {
		return lib.WrapError("error while sending an api request", err)
	}
	if response.Errors != nil {
		return errors.New("error recieved when executing an API call: \n" + string(response.Errors))
	}
	return nil
}

func (client Client) Auth(token string) error {
	variables := map[string]any{
		"token": token,
	}
	return client.sendRequest("auth", authQuery, variables)
}

func (client Client) CreateTeam(teamId string) error {
	variables := map[string]any{
		"teamId": teamId,
	}
	return client.sendRequest(fmt.Sprintf("create team '%s'", teamId), createTeamQuery, variables)
}

func (client Client) DeleteTeam(teamId string, force bool) error {
	variables := map[string]any{
		"teamId": teamId,
		"force":  force,
	}
	return client.sendRequest(fmt.Sprintf("delete team '%s'", teamId), deleteTeamQuery, variables)
}
