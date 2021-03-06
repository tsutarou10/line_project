#!/usr/bin/env bash

set -e
set -u

readonly SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)
source ${SCRIPT_DIR}/../config.sh

readonly TEMPLATE_FILE="${SCRIPT_DIR}/template.yml"
readonly PACKAGED_TEMPLATE_FILE="${SCRIPT_DIR}/packaged_template.yml"

readonly LINE_BOT_CHANNEL_SECRET=$(getSSMValue ${SSM_LINE_BOT_CHANNEL_SECRET})
readonly LINE_BOT_CHANNEL_TOKEN=$(getSSMValue ${SSM_LINE_BOT_CHANNEL_TOKEN})

echo "START BUILD"
make
echo "END BUILD"

aws cloudformation package \
  --region ${REGION} \
  --template-file ${TEMPLATE_FILE} \
  --output-template-file ${PACKAGED_TEMPLATE_FILE} \
  --s3-bucket ${ARTIFACT_BUCKET_NAME} \
  --s3-prefix ${ARTIFACT_BUCKET_PREFIX}

aws cloudformation deploy \
  --no-fail-on-empty-changeset \
  --stack-name ${SERVICE_STACK_NAME} \
  --template-file ${PACKAGED_TEMPLATE_FILE} \
  --capabilities CAPABILITY_NAMED_IAM \
  --parameter-overrides \
    FunctionName=${FUNCTION_NAME} \
    APIName=${API_NAME} \
		DatabaseStackName=${DATABASE_STACK_NAME} \
		LINEBotChannelSecret=${LINE_BOT_CHANNEL_SECRET} \
		LINEBotChannelToken=${LINE_BOT_CHANNEL_TOKEN}

aws cloudformation describe-stacks \
  --region ${REGION} \
  --stack-name ${SERVICE_STACK_NAME} \
  --output table --query Stacks[0].Outputs[0]
