import os
from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport
from dotenv import load_dotenv

# Load important variables from the .env file. Among them there is an API token.
# Consider using a more robust and secure approach when using this in production!
# Also, see `get_api_token` function
load_dotenv()

def non_none_env(var_name: str) -> str:
    value = os.getenv(var_name)
    if value is None:
        raise ValueError(f"'{var_name}' cannot be None")
    return value

def get_api_token() -> str:
    return non_none_env('api_token')

def get_endpoint() -> str:
    serverURL = non_none_env('server_url')
    serviceURI = non_none_env('service_uri')
    return f"{serverURL}/{serviceURI}/gql"

token = get_api_token()
transport = AIOHTTPTransport(url=get_endpoint(), ssl=True)
client = Client(transport=transport, fetch_schema_from_transport=True)

def execute_gql(raw_gql: str):
    result = client.execute(gql(raw_gql))
    print(result)

execute_gql(
    """
    query authLogin {
        authLogin(
            provider: "token",
            credentials: {
                token: "%s"
            }
        ) {
            userTokens {
                userId
            }
            authStatus
        }
    }
    """ % token
)
