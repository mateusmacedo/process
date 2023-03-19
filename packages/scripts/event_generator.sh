#!/bin/bash

function generate_event() {
    MICROSERVICE_NAME=$1
    NAMESPACE=$2
    EVENT_NAME=$3

    HANDLER_DIR="${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/application/events"
    mkdir -p ${HANDLER_DIR}
    echo "package events

import \"${NAMESPACE}/${MICROSERVICE_NAME}/internal/${MICROSERVICE_NAME}/domain\"

type ${EVENT_NAME}Event struct {
    // Adicione os campos do evento aqui
}

type ${EVENT_NAME}Handler struct {
    // Adicione as dependências do handler aqui, como repositórios ou serviços
}

func (h *${EVENT_NAME}Handler) Handle(event *${EVENT_NAME}Event) error {
    // Implemente a lógica do handler aqui
    return nil
}" > "${HANDLER_DIR}/${EVENT_NAME}_event.go"
}
