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
    --publish
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
    -t=*|--target=*
    target="${i#*=}"
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

buildVersion=$(! [ -v $version ] && echo "-t $username/$image:$version" || echo "")

buildVerionLatest="-t $username/$image:latest"

buildTarget=$(! [ -v $target ] && echo "--target $target" || echo "")

docker login -u $user -p $password
docker build $buildTarget $buildVersion $buildVersionLatest .
if [ $publish = true ]
then
    docker push $username/$image:$version
    docker push $username/$image:latest
fi