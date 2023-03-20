#!/bin/bash

set -e

function generate_query() {
    MICROSERVICE_NAME=$1
    NAMESPACE=$2
    QUERY_NAME=$3

    HANDLER_DIR="${MICROSERVICE_NAME}/internal/application/queries"
    mkdir -p ${HANDLER_DIR}
    echo "package handlers

import \"${NAMESPACE}/${MICROSERVICE_NAME}/internal/domain\"

type ${QUERY_NAME}Query struct {
    // Adicione os campos da query aqui
}

type ${QUERY_NAME}Handler struct {
    // Adicione as dependências do handler aqui, como repositórios ou serviços
}

func (h *${QUERY_NAME}Handler) Handle(query *${QUERY_NAME}Query) (interface{}, error) {
    // Implemente a lógica do handler aqui
    return nil, nil
}" > "${HANDLER_DIR}/${QUERY_NAME}_query.go"
}
