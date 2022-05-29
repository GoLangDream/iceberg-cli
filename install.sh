#!/bin/bash

cd /tmp

git clone https://github.com/GoLangDream/iceberg-cli
cd iceberg-cli

go mod tidy
go install


