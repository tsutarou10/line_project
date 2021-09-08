#!/usr/bin/env bash

set -e
set -u

readonly SCRIPT_DIR=$(cd $(dirname ${BASH_SOURCE:-$0}); pwd)
source ${SCRIPT_DIR}/../../config.sh

read -sp "LINE_BOT_CHANNEL_SECRET: " secret
echo ""
aws ssm put-parameter \
	--region ${REGION} \
	--overwrite \
	--name ${SSM_LINE_BOT_CHANNEL_SECRET} \
	--type "SecureString" \
	--value ${secret}

read -sp "LINE_BOT_CHANNEL_TOKEN: " token
echo ""
aws ssm put-parameter \
	--region ${REGION} \
	--overwrite \
	--name ${SSM_LINE_BOT_CHANNEL_TOKEN} \
	--type "SecureString" \
	--value ${token}
