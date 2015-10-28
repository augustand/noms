# Splore

This is a generic noms data explorer.

## Requirements

* Node.js: https://nodejs.org/download/
* You probably want to configure npm to [use a global module path that your user owns](https://docs.npmjs.com/getting-started/fixing-npm-permissions)

## Build

* `./link.sh`
* `npm install`
* `npm run build`

## Run

* `python -m SimpleHTTPServer 8082` (expects ../server to run on same host, port 8000)
* navigate to http://localhost:8082/

## Develop

* `npm run start`

This will start watchify which is continually building a non-minified (and thus debuggable) build.