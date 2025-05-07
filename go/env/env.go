package env

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

type Env struct {
	ServerInfo ServerInfo `json:"server"`
	Token      string     `json:"apiToken"`
}

type ServerInfo struct {
	ServerURL  string `json:"serverURL"`
	ServiceURI string `json:"serviceURI"`
}

func Read() (Env, error) {
	env := Env{}
	bytes, err := os.ReadFile("env/env.json")
	if err != nil {
		return env, err
	}
	err = json.Unmarshal(bytes, &env)
	if err != nil {
		err = lib.WrapError("error while unmarshalling the env file", err)
	}
	return env, err
}

func (env *Env) GraphqlEndpoint() string {
	return env.ServerInfo.ServerURL + "/" + env.ServerInfo.ServiceURI + "/gql"
}

func (env *Env) PurgeToken() {
	env.Token = ""
	runtime.GC()
}
