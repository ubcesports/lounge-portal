# UBCEA Lounge Portal

The UBCEA Lounge Portal is the one stop for all things UBCEA.

## Setup

1. Download and install version 1.23 of [Go](https://go.dev/)
2. Install [Templ](https://templ.guide/)

```
go install github.com/a-h/templ/cmd/templ@latest
```

3. Go into the root directory and run the command to install Templ in the repository.

```
go get -tool github.com/a-h/templ/cmd/templ@latest
```

### Extensions

The extensions below are highly recommended for ease of development.

- [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go)
- [Templ extension](https://marketplace.visualstudio.com/items?itemName=a-h.templ)

## Starting the App
To run the app, run

```
go run .
```

To see the app, go to `localhost:3000`.
