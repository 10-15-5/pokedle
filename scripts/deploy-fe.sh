#!/bin/zsh
echo "$PWD"
docker build -t gabr0236/pokedle-client -f docker/vue-client .

docker tag gabr0236/pokedle-client \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/pokedle-client

docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/pokedle-client

echo "Building is complete"