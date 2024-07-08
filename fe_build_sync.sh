#!/bin/bash 
set -e

if [[ $# -eq 0 ]] ; then
    echo 'required params: apigw_domain s3_bucket'
    exit 1
fi

rm -rf frontend/react/build/
cd frontend/react
echo "REACT_APP_USER_URL='$1'" > .env
export NODE_OPTIONS=--openssl-legacy-provider
npm run build:dev
cd build
aws s3 sync . s3://$2
