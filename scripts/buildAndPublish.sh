#!/usr/bin/env bash

variable_exists () {
    if ! [ -v $0 ]
    then
        echo "Missing variable: " + $1
        exit 1
    fi
}

for i in "$@"
do
case $i in
    -v=*|--version=*
    version="${i#*=}"
    shift
    ;;
    -p|--publish
    publish=true
    shift
    ;;
    -u=*|--user=*
    user="${i#*=}"
    shift
    ;;
    -p=*|--password=*
    password="${i#*=}"
    shift
    ;;
    -i=*|--image=*
    image="${i#*=}"
    shift
    ;;
    *)

    ;;
esac
done

variable_exists user "user"
variable_exists password "password"
variable_exists image "image"

if ! [ -v version] && [ -v publish ]
then
    echo "A version number must be specified in order to publish a package"
    exit 1
fi

docker login -u ${Username} -p ${Password}
docker build -t ${Username}/${Image}:${version} -t ${Username}/${Image}:latest .
if [ $publish = true ]
then
    docker push ${Username}/${Image}:${version}
    docker push ${Username}/${Image}:latest
fi