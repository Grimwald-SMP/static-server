#!/bin/bash
set -e

echo " - Pulling new server version from git"

git pull origin master

go build -o server main.go

# systemctl restart static-server

echo " - Server restarted with new version"