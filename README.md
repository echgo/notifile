# notifile

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/echgo/notifile.svg)](https://golang.org/) [![Go](https://github.com/echgo/notifile/actions/workflows/go.yml/badge.svg)](https://github.com/echgo/notifile/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/echgo/notifile)](https://goreportcard.com/report/github.com/echgo/notifile) [![Go Doc](https://godoc.org/github.com/echgo/notifile?status.svg)](https://pkg.go.dev/github.com/echgo/notifile) [![GitHub issues](https://img.shields.io/github/issues/echgo/notifile)](https://github.com/echgo/notifile/issues) [![GitHub forks](https://img.shields.io/github/forks/echgo/notifile)](https://github.com/echgo/notifile/network) [![GitHub stars](https://img.shields.io/github/stars/echgo/notifile)](https://github.com/echgo/notifile/stargazers) [![GitHub license](https://img.shields.io/github/license/echgo/notifile)](https://github.com/echgo/notifile/blob/master/LICENSE)

With this library you can easily create notification files for your Go projects. These must then be stored in the directory you specify so that echGo can read them.

## Install

```console
go get github.com/echgo/notifile
```

## How to use?

You can use the following example to create a file. Please note that there are currently only the following notification channels: **gotify, matrix, telegram, trello, smtp  & webhook**.

```go
data := notifile.Data{
    Channels:  []string{"gotify", "telegram"},
    Headline: "New notification!",
    Message:  "Here you will find your message.",
}

err := notifile.Create(data, "/var/lib/echgo/notification")
if err != nil {
    log.Fatalln(err)
}
```