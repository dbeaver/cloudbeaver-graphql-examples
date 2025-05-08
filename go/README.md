# Usage of DBeaver Team Edition GraphQL API with Go

How to start:
1. Change the current directory to the directory with this README file.
2. Create the `env.json` file from the `env.json.template` and place it in the same directory:
```sh
cp env/env.json.template env/env.json
```
3. Fill the `env/env.json` file
   * The fields in the `server` object correspond to the fields of the same object from the [cloudbeaver.conf](https://dbeaver.com/docs/cloudbeaver/Server-configuration/#main-server-configuration).
   * `apiToken` is the [API token](https://github.com/dbeaver/cloudbeaver/wiki/Generate-API-access-token).
5. You are ready to Go!
```sh
go build && ./go
```
or
```sh
go run .
```
