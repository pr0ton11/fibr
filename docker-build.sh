#!/bin/sh

TAG="${DOCKER_TAG:=ghcr.io/pr0ton11/fibr}"
CMD="${DOCKER_CMD:=docker}"

echo "Building base docker container..."
${CMD} build -t ${TAG}:base .

echo "Building the image tags for different entrypoints..."
for dockerfile in Dockerfile.*; do
    PART=$(echo $dockerfile | cut -d '.' -f 2)
    echo "Building TAG: ${TAG}:${PART}..."
    sed -i "s/FROM tag/FROM ${TAG}:base" $dockerfile
    echo "FROM ${TAG}:base" > $dockerfile
    cat $dockerfile >> $dockerfile
    ${CMD} build -t ${TAG}:${PART} -f $dockerfile .
done

echo "Cleanup of temporary files..."
rm *.tmp
