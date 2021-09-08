#!/usr/bin/env bash

set -e
set -u

readonly SERVICE_NAME="utna-food"

readonly REGION="ap-northeast-1"

readonly SETUP_STACK_NAME="${SERVICE_NAME}-setup"
readonly ARTIFACT_BUCKET_NAME="${SERVICE_NAME}-artifact-bucket"
readonly ARTIFACT_BUCKET_PREFIX="cloudformation/${SERVICE_NAME}"

readonly SERVICE_STACK_NAME="${SERVICE_NAME}-stack"
readonly FUNCTION_NAME="${SERVICE_NAME}-Function"
readonly TABLE_NAME="UTNAFood"

readonly API_NAME="${SERVICE_NAME}-API"
