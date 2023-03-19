#!/bin/bash

function generate_aggregate() {
    MICROSERVICE_NAME=$1
    NAMESPACE=$2
    AGGREGATE_NAME=$3

    AGGREGATE_DIR="${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/aggregates"
    mkdir -p ${AGGREGATE_DIR}
    echo "package aggregates

type ${AGGREGATE_NAME} struct {
    // Adicione os campos do agregado aqui
}

// Adicione os métodos do agregado aqui
" > "${AGGREGATE_DIR}/${AGGREGATE_NAME}.go"
}
