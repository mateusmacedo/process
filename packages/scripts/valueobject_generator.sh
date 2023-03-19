#!/bin/bash

function generate_valueobject() {
    MICROSERVICE_NAME=$1
    NAMESPACE=$2
    VALUEOBJECT_NAME=$3

    VALUEOBJECT_DIR="${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain/valueobjects"
    mkdir -p ${VALUEOBJECT_DIR}
    echo "package valueobjects

type ${VALUEOBJECT_NAME} struct {
    // Adicione os campos do objeto de valor aqui
}

// Adicione os métodos do objeto de valor aqui
" > "${VALUEOBJECT_DIR}/${VALUEOBJECT_NAME}.go"
}
