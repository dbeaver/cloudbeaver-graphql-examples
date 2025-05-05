package api

import (
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/graphql"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

const authQuery = `
query authLogin($token: String!) {
    authLogin(provider: "token", credentials: { token: $token }) {
        userTokens {
            userId
        }
        authStatus
    }
}
`

type Client struct {
	GraphQLClient graphql.Client
	Endpoint      string
}

func (client Client) sendRequest(request graphql.Request) ([]byte, error) {
	response, err := client.GraphQLClient.Execute(client.Endpoint, request)
	if err != nil {
		return nil, lib.WrapError("error while sending an api request", err)
	}
	return response, nil
}

func (client Client) Auth(token string) ([]byte, error) {
	variables := map[string]any{
		"token": token,
	}
	request := graphql.Request{Query: authQuery, Variables: variables}
	return client.sendRequest(request)
}
