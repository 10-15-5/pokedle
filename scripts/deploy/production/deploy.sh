#!/bin/zsh

if [ $# -lt 2 ]; then
  echo "Error: Missing image tag. Usage: $0 IMAGE_TAG_CLIENT IMAGE_TAG_SERVER"
  exit 1
fi

# Set variables client
IMAGE_NAME_CLIENT=pokedle-client
DEPLOYMENT_NAME_CLIENT=pokedle-client-temp
CONTAINER_NAME_CLIENT=pokedle-client-sha256-1
IMAGE_PATH_CLIENT=docker/production/vue-client
IMAGE_TAG_CLIENT=$1

# Set variables server
IMAGE_NAME_SERVER=pokedle-server
DEPLOYMENT_NAME_SERVER=pokedle-server
CONTAINER_NAME_SERVER=pokedle-server-sha256-1
IMAGE_PATH_SERVER=docker/production/go-server
IMAGE_TAG_SERVER=$2

# Build Docker images locally
docker build -t gabr0236/$IMAGE_NAME_CLIENT -f $IMAGE_PATH_CLIENT . &
docker build -t gabr0236/$IMAGE_NAME_SERVER -f $IMAGE_PATH_SERVER .

# Tag images
docker tag gabr0236/$IMAGE_NAME_CLIENT \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_CLIENT:$IMAGE_TAG_CLIENT &
docker tag gabr0236/$IMAGE_NAME_SERVER \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_SERVER:$IMAGE_TAG_SERVER

# Push Docker images to GCP Artifact Registry
docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_CLIENT:$IMAGE_TAG_CLIENT &
docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_SERVER:$IMAGE_TAG_SERVER

# Update Deployments To Use New Image
kubectl set image deployment/$DEPLOYMENT_NAME_CLIENT $CONTAINER_NAME_CLIENT=europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_CLIENT:$IMAGE_TAG_CLIENT &
kubectl set image deployment/$DEPLOYMENT_NAME_SERVER $CONTAINER_NAME_SERVER=europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_SERVER:$IMAGE_TAG_SERVER

echo "Deployment Finished"