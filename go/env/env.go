package env

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

type Env struct {
	APIToken   string
	serverURL  string
	serviceURI string
}

func Read(path string) (Env, error) {
	env := Env{}
	file, err := os.Open(path)
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
			env.APIToken = after
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

func (env *Env) GraphqlEndpoint() string {
	return env.serverURL + "/" + env.serviceURI + "/gql"
}

func (env *Env) PurgeToken() {
	env.APIToken = ""
	runtime.GC()
}
