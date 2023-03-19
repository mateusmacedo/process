#!/bin/bash

function generate_entity() {
    MICROSERVICE_NAME=$1
    NAMESPACE=$2
    ENTITY_NAME=$3

    ENTITY_DIR="${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/entities"
    mkdir -p ${ENTITY_DIR}
    echo "package entities

type ${ENTITY_NAME} struct {
    // Adicione os campos da entidade aqui
}

// Adicione os métodos da entidade aqui
" > "${ENTITY_DIR}/${ENTITY_NAME}.go"
}
