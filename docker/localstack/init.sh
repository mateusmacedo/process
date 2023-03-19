#!/bin/bash sh
aws configure set aws_access_key_id key --profile=default
aws configure set aws_secret_access_key secret --profile=default
aws configure set region us-east-1 --profile=default

echo "########### AWS Configure ###########"
aws configure list --profile=default

LOCALSTACK_HOST=localhost
LOCALSTACK_ID=000000000000

create_queue() {
    local QUEUE_NAME_TO_CREATE=$1
    awslocal --endpoint-url=http://${LOCALSTACK_HOST}:4566 sqs create-queue --queue-name ${QUEUE_NAME_TO_CREATE}
}

create_queue_with_dlq() {
    local QUEUE_NAME_TO_CREATE=$1
    local DLQ_QUEUE_NAME_TO_CREATE=$1-dlq
    awslocal --endpoint-url=http://${LOCALSTACK_HOST}:4566 sqs create-queue --queue-name ${DLQ_QUEUE_NAME_TO_CREATE} &&
    local DLQ_QUEUE_ARN=$(guess_queue_arn_from_name $DLQ_QUEUE_NAME_TO_CREATE) &&
    awslocal --endpoint-url=http://${LOCALSTACK_HOST}:4566 sqs create-queue --queue-name ${QUEUE_NAME_TO_CREATE} \
    --attributes '{"RedrivePolicy": "{\"deadLetterTargetArn\":\"'"$DLQ_QUEUE_ARN"'\",\"maxReceiveCount\":\"2\"}","VisibilityTimeout": "10"}'
}

guess_queue_arn_from_name() {
    local QUEUE_NAME=$1
    echo "arn:aws:sqs:${AWS_DEFAULT_REGION}:${LOCALSTACK_ID}:$QUEUE_NAME"
}

# Create a queue to bind messages and a DLQ to check if messages are redelivered
# QUEUE_NAME="users-aws"
# echo "creating queue $QUEUE_NAME"
# QUEUE_URL=$(create_queue_with_dlq ${QUEUE_NAME})
# echo "created queue: $QUEUE_URL"
# QUEUE_ARN=$(guess_queue_arn_from_name $QUEUE_NAME)
