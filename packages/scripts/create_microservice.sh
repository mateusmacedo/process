#!/bin/bash

set -e

if [ "$#" -ne 2 ]; then
    echo "Uso: $0 nome_do_microservico namespace_do_projeto"
    echo "Exemplo: $0 meu_microservico github.com/usuario"
    exit 1
fi

MICROSERVICE_NAME=$1
NAMESPACE=$2

mkdir -p ${MICROSERVICE_NAME}/cmd
mkdir -p ${MICROSERVICE_NAME}/internal/application
mkdir -p ${MICROSERVICE_NAME}/internal/application/commands
mkdir -p ${MICROSERVICE_NAME}/internal/application/queries
mkdir -p ${MICROSERVICE_NAME}/internal/application/events


mkdir -p ${MICROSERVICE_NAME}/internal/domain
mkdir -p ${MICROSERVICE_NAME}/internal/domain/aggregates
mkdir -p ${MICROSERVICE_NAME}/internal/domain/entities
mkdir -p ${MICROSERVICE_NAME}/internal/domain/valueobjects
mkdir -p ${MICROSERVICE_NAME}/internal/domain/repositories

mkdir -p ${MICROSERVICE_NAME}/internal/infrastructure
mkdir -p ${MICROSERVICE_NAME}/internal/infrastructure/persistence
mkdir -p ${MICROSERVICE_NAME}/internal/infrastructure/api/grpc
mkdir -p ${MICROSERVICE_NAME}/internal/infrastructure/api/restful

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
