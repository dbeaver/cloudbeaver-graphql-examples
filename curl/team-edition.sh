#!/usr/bin/env sh

set -e

script_dir="$(realpath "$(dirname "$0")")"

cookie_jar="$script_dir/cookie_jar.txt"
touch "$cookie_jar"

execute_gql() {
    data="{
    \"query\": \"$2\",
    \"variables\": { $3 }
}"
    curl \
    --request 'POST' \
    --header 'Content-Type: application/json' \
    --data "$data" \
    --cookie "$cookie_jar" \
    --cookie-jar "$cookie_jar" \
    "$1"
}

. "$script_dir"/../.env
# shellcheck disable=SC2154
# These variables are set in the .env file
gql_endpoint="$server_url/$service_uri/gql"

execute_cb_gql() {
    execute_gql "$gql_endpoint" "$1" "$2"
}

read_operation() {
    sed 's/"/\\"/g' < "$script_dir"/../operations/"$1".gql
}

auth() {
    # shellcheck disable=SC2154
    # api_token is set in the .env file
    execute_cb_gql "$(sed 's/"/\\"/g' < "$script_dir"/../operations/auth.gql)" "\"token\": \"$api_token\""
}

create_team() {
    execute_cb_gql "$(read_operation create_team)" "\"teamId\": \"$1\""
}

delete_team() {
    execute_cb_gql "$(read_operation delete_team)" "\"teamId\": \"$1\", \"force\": true"
}

auth
team_id="cloudbeaver-graphql-examples_curl-team"
create_team "$team_id"
delete_team "$team_id"
