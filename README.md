# CloudBeaver GraphQL examples

This repo contains examples of using GraphQL API for [CloudBeaver Enterprise](https://dbeaver.com/cloudbeaver-enterprise/), [CloudBeaver AWS](https://aws.amazon.com/marketplace/pp/prodview-kijugxnqada5i), and [DBeaver Team Edition](https://dbeaver.com/dbeaver-team-edition/).

# The repo layout

- The [curl](curl) folder contains examples using `curl` command line tool.
- The [go](go) folder contains examples for the Go programming language.
- The [operations](operations) folder contains raw examples. They are used by projects from other folders.
- The [python3](python3) includes examples for Python 3.

## GraphQL API and prerequsites

- You have to have CloudBeaver or Team Edition server running somewhere in you network
- In order to use examples you need to create API token in you user profile
- Configure environment with the server endpoint and the API token above
- To explore GraphQL API you can use GraphQL console which comes with CloudBeaver server at the endpoint `https://SERVER-ADDRESS/api/gql/console`
