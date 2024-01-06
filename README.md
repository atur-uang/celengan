# Celengan

Simple project to save your money in Go.

## GO installation

I suggest you install GO using `goenv`, not using standard Go installation.
This have benefit to install multiple GO version in your computer.
Take a look [this](https://github.com/go-nv/goenv/blob/master/INSTALL.md) to install `goenv`.

## Project installation

```
git clone git@github.com:atur-uang/celengan.git
go mod tidy
```

## Live reload

Run with live reload for development

```
# Go inside project directory
cd celengan

# Download the gin binary
go install github.com/codegangsta/gin@latest

# Run local server
gin -a 8080 run main.go

```

Then open browser http://localhost:3000.