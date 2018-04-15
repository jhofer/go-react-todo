# go-react-todo

    - Todo REST API in Go
    - React frontend client .
    - Setup with docker & docker-compose.

## Requirements

* nodejs / yarn
* docker
* docker-compose
* golang
* glide: https://github.com/Masterminds/glide

## Development setup:

### Configure golang

this optional and just for tools like the golang addon for vscode working correctly.

add path of the server to the gopath

```
export GOPATH=SOME/OTHER/PATH:$HOME/workspace/go-react-todo/server
```

### Start development

run the following code in the project root

```
./dev.sh
```

access webapp via: localhost:4444

## Production setup:

In production livereloading of Go code and a frontend container are not needed. There are now only three containers: postgres database, go backend and a nginx container. A separate docker-compose file is used, as seen below. Run the following scripts in your cloned git directory:

    # will create build/ directory in ../nginx/, where it is then
    # mounted into nginx container to be served as static files
    $ cd client/ && npm run-script build
    $ cd ../ && docker-compose -f docker-compose.prod.yml up -d

You server container will be listening at port 8002 as exposed in docker-compose.prod.yml.

## API endpoints

* **GET** api/v1/todos/
* **GET** api/v1/todos/:id
* **POST** api/v1/todos
* **PUT** api/v1/todos/:id
* **DELETE** api/v1/todos/:id
