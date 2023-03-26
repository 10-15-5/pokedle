#!/bin/zsh

if [ $# -lt 1 ]; then
  echo "Error: Missing image tag. Usage: $1 IMAGE_TAG"
  exit 1
fi

echo "param is: $1"

# Set variables
IMAGE_NAME=pokedle-server-dev #Changed to '-dev'
IMAGE_TAG=$1

docker build -t gabr0236/$IMAGE_NAME -f docker/development/go-server .

docker tag gabr0236/$IMAGE_NAME \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME:$IMAGE_TAG

docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME:$IMAGE_TAG
