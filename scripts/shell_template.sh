#!/bin/bash

LOGPATH="$1"log.log

date +"%Y-%m-%d %H:%M:%S" >> $LOGPATH

echo "There is content." >> $LOGPATH

echo "" >> $LOGPATH