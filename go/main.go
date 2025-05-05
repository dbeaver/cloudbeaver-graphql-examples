package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/api"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/graphql"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
	"github.com/dbeaver/cloudbeaver-graphql-examples/go/variables"
)

func main() {
	if err := main0(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func main0() error {
	variables, err := variables.Read()
	if err != nil {
		return lib.WrapError("error while reading variables", err)
	}
	httpClient, err := graphql.StandardHttpClient()
	if err != nil {
		return err
	}
	graphQLClient := graphql.Client{HttpClient: httpClient}
	apiClient := api.Client{GraphQLClient: graphQLClient, Endpoint: variables.GraphqlEndpoint()}
	response, err := apiClient.Auth(variables.Token)
	variables.PurgeToken()
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
