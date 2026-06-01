# Simple Golang Structure

This is the blueprint of a simple golang project. The aim is to quickly prototype
code without much care about maintainability and testability.

As the name suggests, a simple structure will only work on simple projects, of a
few files and with a very well-defined context.

This is usually a small CLI or a very simple API server.

**Focus:**

- Simplicity

- Minimal time to prod

- Low implementation cost

- No abstractions

There are two main usages here:

- **Simple Service:** For long-running services of a few use-cases.
- **Simple CLI:** CLIs require a bit more code to deal with each command, but apart from that, follow the same structure.


## Components Description

The core components of this architecture are the `main.go` file, which should deal with initialization code, the `app` package where all business logic lives, and auxiliary packages for external calls.

Within the `app` package, we'll have configuration, and the app structure/functions that are to be called from within the `main.go` file to start the service loop (for simple-service) or to perform the desired command.


### CMD

Following best practices for Golang projects, we'll use the `cmd` folder to hold all the binaries `main.go`. This is usually a single one to bootstrap the application. Independently of the nature of the app (server or CLI), It should initialize whatever dependencies is has and start the main loop/command.

```
root/
  cmd/
    simple-service/
      main.go
    simple-cli/
      main.go
  internal/
    app/
      config.go
      app.go
      models.go
      handle_x.go
      command_x.go
```

Configuration should live in `config.go` file, and the business logic within `app.go` and any other auxiliary file within the same package.

### Configuration

Although the goal is for simplicity and speed, a bit of structure is needed, and I believe structured configuration is a minimum requirement for any good software.

The dependencies to achieve that are `godotenv` for helping loading environment files, and `env`, a small library that simply parses env variables to a struct with defaults. 

### Simple API Server

If the project requires a simple API server, it should declare it in `app.go`, and use `http.ServeMux` for the handlers. Route definition lives inside `app.go`, but the handlers can be declared in dedicated files in the package, to avoid it getting too big, and organize code by functionality/context. It's preferable to use method handlers, so we can store configuration within the server struct instead of using globals.

To avoid overengineering, the handlers will contain all business logic, common code should still be isolated to keep it DRY, and single responsibility principles should also be applied.

Handler files should declare their input/output structs in the same file, and right above the handler method, for simplicity and ease to trace its definition.

### Simple CLI

For CLIs, we'll need another dependency, [`github.com/alecthomas/kong`](https://github.com/alecthomas/kong). This is an alternative to cobra, and serves the purpose of dealing with multiple-commands with a better lighter interface.

Commands should be declared in their own file inside `app` package, and `main.go` will be responsible for initialization only. This pattern will allow business rules to live inside the `app` package, and follow the same pattern as the API server, where each handler is a dedicated file, and owns business logic.

### Models

The models file should include anything related to the business rules data. It's not the API input/output, but rather internal data representations that are core to the app.

### External Calls

Usually an application has to call external services through their APIs, to collect data. For those, we'll keep it simple here, and just have another package if that involves **more than two functions**. Otherwise, it's safe to keep it in a file within the `app` package or in the handler if that is only called by it.
