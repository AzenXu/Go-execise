#!/usr/bin/env bash

# 部署脚本
cp ./template ./bin/

mkdir ./bin/videos

cd bin

nohup ./api &
nohup ./dispatcher &
nohup ./stream &
nohup ./web &

echo  "🎉 deploy done~ have fun~~~"