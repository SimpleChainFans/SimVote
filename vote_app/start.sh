#!/bin/bash

docker stop vote_app
docker run -v `pwd`/conf/config.dev.yaml:/app/conf/config.yaml --rm -d -p 7688:7688 --name vote_app vote_app:0.0.1
