#!/bin/zsh

# Require tag param
if [ $# -lt 1 ]; then
  echo "Error: Missing image tag. Usage: $1 IMAGE_TAG"
  exit 1
fi

# Set variables
IMAGE_NAME=pokedle-client
DEPLOYMENT_NAME=pokedle-client-temp
CONTAINER_NAME=pokedle-client-sha256-1
IMAGE_PATH=docker/production/vue-client
IMAGE_TAG=$1

# Build Docker image locally
docker build -t gabr0236/$IMAGE_NAME -f $IMAGE_PATH .

# Tag image
docker tag gabr0236/$IMAGE_NAME \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME:$IMAGE_TAG

# Push Docker image to GCP Artifact Registry
docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME:$IMAGE_TAG

# Update Deployment To Use New Image
kubectl set image deployment/$DEPLOYMENT_NAME $CONTAINER_NAME=europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME:$IMAGE_TAG
