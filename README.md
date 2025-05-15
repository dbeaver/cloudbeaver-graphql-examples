# CloudBeaver GraphQL examples

This repo contains examples of using GraphQL API for [CloudBeaver Enterprise](https://dbeaver.com/cloudbeaver-enterprise/), [CloudBeaver AWS](https://aws.amazon.com/marketplace/pp/prodview-kijugxnqada5i), and [DBeaver Team Edition](https://dbeaver.com/dbeaver-team-edition/).

CloudBeaver and Team Edition communicate with their web browser frontends by exposing a GraphQL API. In some cases, users may want to leverage the same API to enable advanced use cases by automating certain operations usually done using UI. To explore the API, you can [use an embedded GraphiQL console](https://github.com/dbeaver/cloudbeaver/wiki/Server-API-explorer). You can look at the real examples of using the API in this repository.
 
## The repo layout

- The [curl](curl) folder contains examples using `curl` command line tool.
- The [go](go) folder contains examples for the Go programming language.
- The [operations](operations) folder contains raw examples. They are used by projects from other folders.
- The [python3](python3) includes examples for Python 3.
