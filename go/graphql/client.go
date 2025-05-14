package graphql

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/lib"
)

type Client struct {
	HttpClient *http.Client
}

func (client Client) Execute(endpoint string, request Request) (Response, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return Response{}, lib.WrapError("unable to marshal a GraphQL request", err)
	}
	httpResponse, err := client.HttpClient.Post(endpoint, "application/json", bytes.NewReader(payload))
	if err != nil {
		return Response{}, lib.WrapError("unable to execute POST request", err)
	}
	defer lib.CloseOrWarn(httpResponse.Body)
	rawResponseBody, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return Response{}, lib.WrapError("unable to read GraphQL response body", err)
	}
	var response Response
	if err = json.Unmarshal(rawResponseBody, &response); err != nil {
		err = lib.WrapError("error while unmarshalling GraphQL response body into json", err)
	}
	return response, err
}
