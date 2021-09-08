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
readonly REGISTRATION_STATUS_TABLE_NAME="UTNAFoodRegistrationStatus"

# ==== service ====
readonly SERVICE_STACK_NAME="${SERVICE_NAME}-stack"
readonly FUNCTION_NAME="${SERVICE_NAME}-Function"
readonly API_NAME="${SERVICE_NAME}-API"
