#!/usr/bin/env bash

APP_NAME=food-delivery

echo "Run stg.sh"

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME} #stop container ${APP_NAME}
# docker rmi $(docker images -qa -f 'dangling=true')

docker run -d --name ${APP_NAME} \
  --network my-net \
  -e VIRTUAL_HOST="159.65.137.18" \
  -e LETSENCRYPT_HOST="159.65.137.18" \
  -e LETSENCRYPT_EMAIL="vanthanh.0610998@gmail.com" \
  -e MYSQL_CONN_STRING="demo:root_password@tcp(mysql:3306)/fd?charset=utf8mb4&parseTime=True&loc=Local" \
  -e S3BucketName=food-deliveri-images \
  -e S3Region=ap-southeast-1 \
  -e S3APIKey=AKIAYDBF24QKNBK6PBME \
  -e S3SecretKey="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" \
  -e S3Domain="https://d1nt4b07cd5syr.cloudfront.net" \
  -e SYSTEM_SECRET=van_thanh \
  -p 8080:8080 \
  ${APP_NAME}
