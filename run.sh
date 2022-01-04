#!/bin/bash
go mod tidy
go mod download
go run src/app.go
