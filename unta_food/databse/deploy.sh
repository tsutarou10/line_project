#!/usr/bin/env bash

set -e
set -u

readonly SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)
source ${SCRIPT_DIR}/../config.sh

readonly TEMPLATE_FILE="${SCRIPT_DIR}/template.yml"

aws cloudformation deploy \
  --no-fail-on-empty-changeset \
  --region ${REGION} \
  --stack-name ${DATABASE_STACK_NAME} \
  --template-file ${TEMPLATE_FILE} \
  --capabilities CAPABILITY_NAMED_IAM \
  --parameter-overrides \
    UTNAFoodTableName=${UTNA_FOOD_TABLE_NAME} \
		VisitedRestaurantTableName=${VISITED_RESTAURANT_TABLE_NAME} \

aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name ${DATABASE_STACK_NAME} \
  --output table
