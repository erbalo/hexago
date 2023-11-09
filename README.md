# hexago

This project has the purpose of showing the hexagonal architecture implementation using golang.

## How to run?

### With Makefile

Option a

```shell
$ make all
```

Option b

```shell
$ make dependencies
$ make run
```

Clean the project

```shell
$ make clean
```

> **Note:** You can specify the CLI including the parameter CLI_PARAM

| parameter | expectations | default | example |
| :---: | :--- | :---: | :---: |
| http  |   The project sets a server in localhost:8080 | yes | `CLI_PARAM=http` |
| cli | The project sets a terminal CLI for running the application | no | `CLI_PARAM=cli` |

#### Example

Option a

```shell
$ make all CLI_PARAM=cli
```

Option b

```shell
$ make dependencies
$ make run CLI_PARAM=cli
```

### With Go

**Option a.** The HTTP server

```shell
$ go run cmd/http/main.go
```

**Option b.** The CLI application

```shell
$ go run cmd/cli/main.go
```
