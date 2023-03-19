#!/bin/bash

function generate_repository() {
    MICROSERVICE_NAME=$1
    NAMESPACE=$2
    REPOSITORY_NAME=$3

    REPOSITORY_DIR="${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/repositories"
    mkdir -p ${REPOSITORY_DIR}
    echo "package repositories

type ${REPOSITORY_NAME} interface {
    // Adicione os métodos do repositório aqui
}
" > "${REPOSITORY_DIR}/${REPOSITORY_NAME}.go"
}
