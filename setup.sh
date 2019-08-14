#!/usr/bin/env bash

# Read in new module name, title and description
echo 'Enter the name for your module (e.g. github.com/ORG/NAME):'
read -r MODULE_NAME

echo 'Enter a title for your project (e.g. My Cool Project) NOTE: you can change this later in the swagger.yml file:'
read -r TITLE

echo 'Enter a description for your project (e.g. My project is cool and does some things) NOTE: you can change this later in the swagger.yml file:'
read -r DESCRIPTION

# Change yaml file
LC_ALL=C find . -type f -name 'swagger.yml' -exec sed -i '' "s%A boilerplate for a Go Api%$DESCRIPTION%g" {} +
LC_ALL=C find . -type f -name 'swagger.yml' -exec sed -i '' "s%Boilerplate%$TITLE%g" {} +



# Change imports and variables
LC_ALL=C find . -type f -name '*.go' -exec sed -i '' "s%github.com/adeptmind/adept-go-postgres-api-boilerplate%$MODULE_NAME%g" {} +
LC_ALL=C find . -type f -name '*.mod' -exec sed -i '' "s%github.com/adeptmind/adept-go-postgres-api-boilerplate%$MODULE_NAME%g" {} +
LC_ALL=C find . -type f -name 'Dockerfile' -exec sed -i '' "s%github.com/adeptmind/adept-go-postgres-api-boilerplate%$MODULE_NAME%g" {} +

OBJECT_NAME=$(echo "$TITLE" | python3 -c "import sys; print(sys.stdin.read().title())" | tr -d '[:space:]')
LC_ALL=C find . -type f -name '*.go' -exec sed -i '' "s%Boilerplate%$OBJECT_NAME%g" {} +

# Generate the new files
mkdir -p ./gen && swagger generate server -A "$(echo "$TITLE" | tr '[:upper:]' '[:lower:]')" -f ./swagger.yml --exclude-main -t ./gen

# Change configure file
mv 'gen/restapi/configure_boilerplate.go' "gen/restapi/$(ls gen/restapi/ | grep configure | grep -v -e boilerplate | tr -d '[:space:]')"

# Delete gen gitignore, yourself and git
rm -rf .git
rm gen/.gitignore
rm setup.sh

echo 'You are done! Be sure to initialize a new git repo and commit the project initially.'
