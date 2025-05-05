package graphql

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
}

func (client Client) Execute(url string, request Request) ([]byte, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	response, err := client.HttpClient.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer closeOrWarn(response.Body)
	return io.ReadAll(response.Body)
}

func closeOrWarn(closer io.Closer) {
	if err := closer.Close(); err != nil {
		slog.Warn("error while closing a Closer: " + err.Error())
	}
}
