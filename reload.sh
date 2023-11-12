#!/bin/sh
git stash
git pull
git stash pop
make build
fuser -k -n tcp 8000
nohup ./bin/biz-mail -conf ./configs/config.yaml  > biz_mail.log &