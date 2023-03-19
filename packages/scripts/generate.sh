#!/bin/bash

if [ "$#" -ne 4 ]; then
    echo "Uso: $0 tipo nome_do_microservico namespace_do_projeto nome_do_elemento"
    echo "Exemplo: $0 command meu_microservico github.com/usuario CreateUser"
    echo "Tipos: command, query, event, aggregate, entity, valueobject, repository"
    exit 1
fi

TYPE=$1
MICROSERVICE_NAME=$2
NAMESPACE=$3
ELEMENT_NAME=$4

case $TYPE in
    command|query|event|aggregate|entity|valueobject|repository)
        source ./${TYPE}_generator.sh
        generate_${TYPE} "${MICROSERVICE_NAME}" "${NAMESPACE}" "${ELEMENT_NAME}"
        ;;
    *)
        echo "Tipo inválido. Os tipos válidos são: command, query, event, aggregate, entity, valueobject, repository"
        exit 1
        ;;
esac
