#!/bin/bash

cd /tmp

git clone --depth=1 https://github.com/GoLangDream/iceberg-cli
cd iceberg-cli

go mod tidy
go install


