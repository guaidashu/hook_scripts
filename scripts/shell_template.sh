#!/bin/bash

LOGPATH="$1"logs/project_log.log

date +"%Y-%m-%d %H:%M:%S" >> $LOGPATH

cd project_root_dir
git pull >> $LOGPATH 2>&1
echo "" >> $LOGPATH
sudo supervisorctl stop project_name >> $LOGPATH 2>&1

echo "" >> $LOGPATH

export GOPATH=/home/ubuntu/
export GOCACHE="/home/ubuntu/.cache/go-build"
go build >> $LOGPATH 2>&1

sudo supervisorctl start project_name >> $LOGPATH 2>&1

echo "" >> $LOGPATH

echo "" >> $LOGPATH