#!/bin/zsh

if [ $# -lt 2 ]; then
  echo "Error: Missing image tag. Usage: $0 IMAGE_TAG_CLIENT IMAGE_TAG_SERVER"
  exit 1
fi

# Set variables client
IMAGE_NAME_CLIENT=pokedle-client-dev #Changed to '-dev'
DEPLOYMENT_NAME_CLIENT=pokedle-client-dev
CONTAINER_NAME_CLIENT=pokedle-client-sha256-1
IMAGE_PATH_CLIENT=docker/development/vue-client
IMAGE_TAG_CLIENT=$1

# Set variables server
IMAGE_NAME_SERVER=pokedle-server-dev #Changed to '-dev'
DEPLOYMENT_NAME_SERVER=pokedle-server-dev
CONTAINER_NAME_SERVER=pokedle-server-sha256-1
IMAGE_PATH_SERVER=docker/development/go-server
IMAGE_TAG_SERVER=$2

# Build Docker images locally
docker build -t gabr0236/$IMAGE_NAME_CLIENT -f $IMAGE_PATH_CLIENT . &
docker build -t gabr0236/$IMAGE_NAME_SERVER -f $IMAGE_PATH_SERVER .

wait 

# Tag images
docker tag gabr0236/$IMAGE_NAME_CLIENT \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_CLIENT:$IMAGE_TAG_CLIENT &
docker tag gabr0236/$IMAGE_NAME_SERVER \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_SERVER:$IMAGE_TAG_SERVER

wait 

# Push Docker images to GCP Artifact Registry
docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_CLIENT:$IMAGE_TAG_CLIENT &
docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_SERVER:$IMAGE_TAG_SERVER

wait 

# Update Deployments To Use New Image
kubectl set image deployment/$DEPLOYMENT_NAME_CLIENT $CONTAINER_NAME_CLIENT=europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_CLIENT:$IMAGE_TAG_CLIENT &
kubectl set image deployment/$DEPLOYMENT_NAME_SERVER $CONTAINER_NAME_SERVER=europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/$IMAGE_NAME_SERVER:$IMAGE_TAG_SERVER

echo "Deployment Finished"