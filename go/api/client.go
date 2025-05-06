package api

import (
	"encoding/json"
	"errors"

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

func (client Client) sendRequest(query string, variables map[string]any) (json.RawMessage, error) {
	request := graphql.Request{Query: query, Variables: variables}
	response, err := client.GraphQLClient.Execute(client.Endpoint, request)
	if err != nil {
		return json.RawMessage{}, lib.WrapError("error while sending an api request", err)
	}
	if response.Errors != nil {
		err = errors.New("error recieved when executing an API call: \n" + string(response.Errors))
	}
	return response.Data, err
}

func (client Client) Auth(token string) (json.RawMessage, error) {
	variables := map[string]any{
		"token": token,
	}
	return client.sendRequest(authQuery, variables)
}

func (client Client) CreateTeam(teamId string) (json.RawMessage, error) {
	variables := map[string]any{
		"teamId": teamId,
	}
	return client.sendRequest(createTeamQuery, variables)
}

func (client Client) DeleteTeam(teamId string, force bool) (json.RawMessage, error) {
	variables := map[string]any{
		"teamId": teamId,
		"force":  force,
	}
	return client.sendRequest(deleteTeamQuery, variables)
}
