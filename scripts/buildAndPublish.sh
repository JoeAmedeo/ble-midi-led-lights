#!/usr/bin/env bash

for i in "$@"
do
case $i in
    -v=*|--version=*)
    version="${i#*=}"
    shift
    ;;
    --publish)
    publish=true
    shift
    ;;
    -u=*|--user=*)
    user="${i#*=}"
    shift
    ;;
    -p=*|--password=*)
    password="${i#*=}"
    shift
    ;;
    -i=*|--image=*)
    image="${i#*=}"
    shift
    ;;
    -t=*|--target=*)
    target="${i#*=}"
    shift
    ;;
    *)

    ;;
esac
done

# TODO: very repetative, fix later
if ! [ -v user ]
then
    echo "Missing variable: " + $1
    exit 1
fi

if ! [ -v password ]
then
    echo "Missing variable: " + $1
    exit 1
fi

if ! [ -v image ]
then
    echo "Missing variable: " + $1
    exit 1
fi
# TODO: end

if ! [ -v version ] && [ -v publish ]
then
    echo "A version number must be specified in order to publish a package"
    exit 1
fi

buildVersion=$(! [ -v $version ] && echo "-t $user/$image:$version" || echo "")

buildVersionLatest=$( echo "-t $user/$image:latest" )

buildTarget=$(! [ -v $target ] && echo "--target $target" || echo "")

docker login -u $user -p $PASSWORD
docker buildx build --platform linux/arm64 $buildTarget $buildVersion $buildVersionLatest .
if [ "$publish" = true ]
then
    docker push $user/$image:$version
    docker push $user/$image:latest
fi