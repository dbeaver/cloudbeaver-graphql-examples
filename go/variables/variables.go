package variables

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

type Variables struct {
	ServerInfo ServerInfo `json:"server"`
	Token      string     `json:"apiToken"`
}

type ServerInfo struct {
	ServerURL  string `json:"serverURL"`
	ServiceURI string `json:"serviceURI"`
}

func Read() (Variables, error) {
	variables := Variables{}
	bytes, err := os.ReadFile("variables/variables.json")
	if err != nil {
		return variables, err
	}
	err = json.Unmarshal(bytes, &variables)
	if err != nil {
		err = lib.WrapError("error while unmarshalling the variables file", err)
	}
	return variables, err
}

func (variables *Variables) GraphqlEndpoint() string {
	return variables.ServerInfo.ServerURL + "/" + variables.ServerInfo.ServiceURI + "/gql"
}

func (variables *Variables) PurgeToken() {
	variables.Token = ""
	runtime.GC()
}
