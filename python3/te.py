#!/usr/bin/env python3

import asyncio
import os
from typing import Optional, Dict, Any

import aiohttp
from dotenv import load_dotenv
from gql import gql, Client
from gql.transport.aiohttp import AIOHTTPTransport


async def main():
    # Load important variables from the .env file. Among them there is an API token.
    # Consider using a more robust and secure approach when using this in production!
    script_dir = os.path.dirname(os.path.abspath(__file__))
    parent_dir = os.path.dirname(script_dir)
    dotenv_path = os.path.join(parent_dir, ".env")
    load_dotenv(dotenv_path=dotenv_path)

    server_url = non_none_env('server_url')
    service_uri = non_none_env('service_uri')
    api_client = APIClient(f"{server_url}/{service_uri}/gql")

    await api_client.execute_query("auth", {
        "token": non_none_env('api_token'),
    })
    team_id = "cloudbeaver-graphql-examples_python3-team"
    await api_client.execute_query("create_team", {
        "teamId": team_id,
    })
    await api_client.execute_query("delete_team", {
        "teamId": team_id,
        "force": True,
    })

def non_none_env(var_name: str) -> str:
    value = os.getenv(var_name)
    if value is None:
        raise ValueError(f"'{var_name}' cannot be None")
    return value

class APIClient:
    _client: Client

    def __init__(self, endpoint: str):
        transport = AIOHTTPTransport(url=endpoint, ssl=True, client_session_args={'cookie_jar': aiohttp.CookieJar()})
        self._client = Client(transport=transport, fetch_schema_from_transport=True)

    async def execute_query(self, operation_name: str, variables: Optional[Dict[str, Any]] = None):
        with open("../operations/" + operation_name + ".gql", encoding="utf-8") as file:
            result = await self._client.execute_async(document=gql(file.read()), variable_values=variables)
            print(result)

asyncio.run(main())
