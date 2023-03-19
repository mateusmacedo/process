#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Uso: $0 nome_do_microservico namespace_do_projeto"
    echo "Exemplo: $0 meu_microservico github.com/usuario"
    exit 1
fi

MICROSERVICE_NAME=$1
NAMESPACE=$2

mkdir -p ${MICROSERVICE_NAME}/cmd/${MICROSERVICE_NAME}
touch ${MICROSERVICE_NAME}/cmd/${MICROSERVICE_NAME}/main.go

mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/cqrs/command/handlers
touch ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/cqrs/command/commands.go
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/cqrs/query/handlers
touch ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/cqrs/query/queries.go

mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/eventsourcing/eventstore
touch ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/eventsourcing/events.go

mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/aggregates
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/entities
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/valueobjects
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/repositories

mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/persistence/repository_impl
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/api/grpc
touch ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/api/grpc/server.go
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/api/restful
touch ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/api/restful/server.go

mkdir -p ${MICROSERVICE_NAME}/pkg/shared/errors
mkdir -p ${MICROSERVICE_NAME}/pkg/shared/utils

touch ${MICROSERVICE_NAME}/go.sum
touch ${MICROSERVICE_NAME}/Dockerfile
touch ${MICROSERVICE_NAME}/.gitignore
touch ${MICROSERVICE_NAME}/README.md

# Inicializa o go.mod com o namespace fornecido
cd ${MICROSERVICE_NAME}
go mod init ${NAMESPACE}/${MICROSERVICE_NAME}
