#! /bin/bash

dep ensure -v
make
make docker-build $1:$2
make docker-push $1:$2
minikube start
make deploy
