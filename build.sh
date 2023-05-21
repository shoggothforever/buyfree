#!/bin/bash
go mod tidy
swag init
go build main.go