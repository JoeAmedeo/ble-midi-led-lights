#!/usr/bin/env bash

Version="0.0.1"
Publish=false

for i in "$@"
do
case $i in
    -v=*|--version=*
    Version="${i#*=}"
    shift
    ;;
    -p|--publish
    publish=true
    shift
    ;;
    *)

    ;;
esac
done

docker login -u ${Username} -p ${Password}
docker build -t ${Username}/${Image}:${Version} -t ${Username}/${Image}:latest .
if [ $Publish ]
then
    docker push ${Username}/${Image}:${Version}
    docker push ${Username}/${Image}:latest
fi