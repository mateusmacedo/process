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

mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/commands
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/queries
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/events


mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/aggregates
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/entities
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/valueobjects
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/repositories

mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/persistence
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/api/grpc
mkdir -p ${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/infrastructure/api/restful

mkdir -p ${MICROSERVICE_NAME}/pkg/shared
mkdir -p ${MICROSERVICE_NAME}/pkg/shared/errors
mkdir -p ${MICROSERVICE_NAME}/pkg/shared/utils

touch ${MICROSERVICE_NAME}/go.sum
touch ${MICROSERVICE_NAME}/Dockerfile
touch ${MICROSERVICE_NAME}/.gitignore
touch ${MICROSERVICE_NAME}/README.md

# Inicializa o go.mod com o namespace fornecido
cd ${MICROSERVICE_NAME}
go mod init ${NAMESPACE}/${MICROSERVICE_NAME}
