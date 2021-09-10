#!/usr/bin/env bash

set -e
set -u

# ==== common ====
readonly SERVICE_NAME="utna-food"
readonly REGION="ap-northeast-1"

# ==== setup ====
readonly SETUP_STACK_NAME="${SERVICE_NAME}-setup"
readonly ARTIFACT_BUCKET_NAME="${SERVICE_NAME}-artifact-bucket"
readonly ARTIFACT_BUCKET_PREFIX="cloudformation/${SERVICE_NAME}"

# ==== database ====
readonly DATABASE_STACK_NAME="${SERVICE_NAME}-database"
readonly UTNA_FOOD_TABLE_NAME="UTNAFood"
readonly VISITED_RESTAURANT_TABLE_NAME="VisitedRestaurant"

# ==== service ====
readonly SERVICE_STACK_NAME="${SERVICE_NAME}-stack"
readonly FUNCTION_NAME="${SERVICE_NAME}-Function"
readonly API_NAME="${SERVICE_NAME}-API"

# ==== ssm ====
readonly SSM_LINE_BOT_CHANNEL_SECRET="/${SERVICE_NAME}/LINE_BOT_CHANNEL_SECRET"
readonly SSM_LINE_BOT_CHANNEL_TOKEN="/${SERVICE_NAME}/LINE_BOT_CHANNEL_TOKEN"

function getSSMValue() {
	local value=$(aws --region ${REGION} ssm get-parameter --name ${1} --with-decryption --output text --query "Parameter.Value" 2>&1)
	echo ${value}
}
