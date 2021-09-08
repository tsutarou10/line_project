#!/usr/bin/env bash

set -e
set -u

readonly SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)
source ${SCRIPT_DIR}/../config.sh

readonly TEMPLATE_FILE="${SCRIPT_DIR}/template.yml"

aws cloudformation deploy \
  --region ${REGION} \
  --stack-name ${SETUP_STACK_NAME} \
  --template-file ${TEMPLATE_FILE} \
  --capabilities CAPABILITY_IAM \
  --parameter-overrides \
    ArtifactBucketName=${ARTIFACT_BUCKET_NAME}

aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name ${SETUP_STACK_NAME} \
  --output table --query Stacks[0].Outputs[0]
