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

const (
	objectPrefix = "cloudbeaver-graqhql-examples-go-"
	teamId       = objectPrefix + "team"
	projectName  = objectPrefix + "project"
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
	apiClient := initClient(env.GraphqlEndpoint())

	// Auth
	err = apiClient.Auth(env.Token)
	env.PurgeToken()
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
	err = apiClient.AddProjectAccess(projectId, teamId)
	if err != nil {
		return err
	}

	return nil
}

func initClient(endpoint string) api.Client {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		// Invariant: the method that creates cookie jar with no options never returns non-nil err
		panic("encountered error while creating a cookie jar! " + err.Error())
	}
	graphQLClient := graphql.Client{HttpClient: &http.Client{Jar: cookieJar}}
	return api.Client{GraphQLClient: graphQLClient, Endpoint: endpoint}
}

func cleanup(callDescription string, apiCall func() error) {
	err := apiCall()
	if err != nil {
		slog.Warn("unable to " + callDescription)
		slog.Warn(err.Error())
	}
}
