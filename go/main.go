package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/api"
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
	envFlag := flag.String("env", "../.env", "Path to the .env file")
	operationsFlag := flag.String("operations", "../operations", "Path to the folder with GraphQL operations")
	flag.Parse()
	env, err := readEnv(*envFlag)
	if err != nil {
		return lib.WrapError("error while reading variables", err)
	}
	apiClient := initClient(env.serverURL+"/"+env.serviceURI+"/gql", *operationsFlag)

	// Auth
	err = apiClient.Auth(env.apiToken)
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

type env struct {
	apiToken   string
	serverURL  string
	serviceURI string
}

func readEnv(envFilePath string) (env, error) {
	env := env{}
	file, err := os.Open(envFilePath)
	if err != nil {
		return env, lib.WrapError("error while opening the env file", err)
	}
	defer lib.CloseOrWarn(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		before, after, found := strings.Cut(scanner.Text(), "=")
		if !found {
			continue
		}
		before = strings.TrimSpace(before)
		after = strings.TrimSpace(after)
		switch before {
		case "api_token":
			env.apiToken = after
		case "server_url":
			env.serverURL = after
		case "service_uri":
			env.serviceURI = after
		default:
			slog.Warn(fmt.Sprintf("unknown env variable: %s", before))
		}
	}
	return env, err
}

func initClient(endpoint string, operationsPath string) api.Client {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		// Invariant: the method that creates cookie jar with no options never returns non-nil err
		panic("encountered error while creating a cookie jar! " + err.Error())
	}
	graphQLClient := graphql.Client{HttpClient: &http.Client{Jar: cookieJar}}
	return api.Client{GraphQLClient: graphQLClient, Endpoint: endpoint, OperationsPath: operationsPath}
}

func cleanup(callDescription string, apiCall func() error) {
	err := apiCall()
	if err != nil {
		slog.Warn("unable to " + callDescription)
		slog.Warn(err.Error())
	}
}
