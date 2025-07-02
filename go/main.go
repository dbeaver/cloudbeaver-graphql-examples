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
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func run() error {
	// Instantiate a client
	envFlag := flag.String("env", "../.env", "Path to the .env file")
	operationsFlag := flag.String("operations", "../operations", "Path to the folder with GraphQL operations")
	flag.Parse()
	env, err := readEnv(*envFlag)
	if err != nil {
		return lib.WrapError("error while reading variables", err)
	}
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		// Invariant: the method that creates cookie jar with no options never returns non-nil err
		panic("encountered error while creating a cookie jar! " + err.Error())
	}
	graphQLClient := graphql.Client{HttpClient: &http.Client{Jar: cookieJar}}
	apiClient := api.Client{
		GraphQLClient:  graphQLClient,
		ServerURL:      env.serverURL,
		Endpoint:       env.serverURL + "/" + env.serviceURI + "/gql",
		OperationsPath: *operationsFlag,
	}

	// Run scenarios
	if err := scenario0(&apiClient, env.apiToken); err != nil {
		return err
	}
	return scenario1(&apiClient, env.licenseText)
}

type env struct {
	apiToken    string
	serverURL   string
	serviceURI  string
	licenseText string
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
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		before, after, found := strings.Cut(line, "=")
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
		case "license_text":
			env.licenseText = after
		default:
			slog.Warn(fmt.Sprintf("unknown env variable: %s", before))
		}
	}
	return env, err
}
