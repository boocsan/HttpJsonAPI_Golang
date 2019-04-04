#!/bin/bash

git clone https://github.com/iedred7584/HttpJsonAPI_Golang.git
docker build -t jsonapi ./HttpJsonAPI_Golang/
docker run -dp 8888:8888 --name jsonapi jsonapi
