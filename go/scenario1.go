package main

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/dbeaver/cloudbeaver-graphql-examples/go/api"
)

// This scenario is used to perform initial configuration of the server, such as setting up the license and local admin.
func scenario1(apiClient *api.Client, licenseText string) error {
	username := "teadmin"
	password := "VeryStr0ngPassw0rd"

	// USE WITH CAUTION: this disables TLS certt verification, which is not recommended for production use
	if false {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	rawResponse, err := apiClient.InitialConfiguration(username, password, objectPrefix+"server-name")
	if err != nil {
		return err
	}
	fmt.Println(string(rawResponse))

	if err := apiClient.LocalAuth(username, password); err != nil {
		return err
	}

	rawResponse, err = apiClient.ImportLicense(licenseText)
	if err != nil {
		return err
	}
	fmt.Println(string(rawResponse))

	return nil
}
