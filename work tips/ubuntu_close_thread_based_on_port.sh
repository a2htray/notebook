#!/bin/bash

port=`lsof -i:$1 | tail -n 1 | awk '{print $2}'`

if [[ ! -z "$port" ]] # 判断字符串不为空
then
  kill -9 ${port}
  echo "close thread on port $port"
else
  echo "port is empty"
fi
