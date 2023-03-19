#!/bin/bash

function generate_command() {
    MICROSERVICE_NAME=$1
    NAMESPACE=$2
    COMMAND_NAME=$3

    HANDLER_DIR="${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/cqrs/command/handlers"
    mkdir -p ${HANDLER_DIR}
    echo "package handlers

import \"${NAMESPACE}/${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain\"

type ${COMMAND_NAME}Command struct {
    // Adicione os campos do comando aqui
}

type ${COMMAND_NAME}Handler struct {
    // Adicione as dependências do handler aqui, como repositórios ou serviços
}

func (h *${COMMAND_NAME}Handler) Handle(command *${COMMAND_NAME}Command) error {
    // Implemente a lógica do handler aqui
    return nil
}" > "${HANDLER_DIR}/${COMMAND_NAME}_command.go"
}
