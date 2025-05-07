package main

import (
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/api"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/env"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/graphql"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

func main() {
	if err := main0(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func main0() error {
	// Instantiate a client
	env, err := env.Read()
	if err != nil {
		return lib.WrapError("error while reading variables", err)
	}
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return lib.WrapError("unable to create a cookie jar", err)
	}
	graphQLClient := graphql.Client{HttpClient: &http.Client{Jar: cookieJar}}
	apiClient := api.Client{GraphQLClient: graphQLClient, Endpoint: env.GraphqlEndpoint()}

	// Auth
	err = apiClient.Auth(env.Token)
	env.PurgeToken()
	if err != nil {
		return err
	}

	// Create a team
	teamId := "exampleTeamId"
	err = apiClient.CreateTeam(teamId)
	if err != nil {
		return err
	}

	// Delete a team
	err = apiClient.DeleteTeam(teamId, true)
	if err != nil {
		return err
	}

	return nil
}
