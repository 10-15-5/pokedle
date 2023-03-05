# Pokedle

Test your traine knowledge!

## Recommended IDE Setup

[VSCode](https://code.visualstudio.com/) 
[Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) 

## Customize configuration

See [Vite Configuration Reference](https://vitejs.dev/config/).

## Run

### With download/install of node packages

```sh
npm run dev
```

### Without download/install of node packages

```sh
npm run dev-fast
```


## google cloud

Connect to cluster
gcloud container clusters get-credentials pokedle-cluster --region europe-west1 --project encoded-might-376114

### Add image to gcloud using local cli
https://cloud.google.com/artifact-registry/docs/docker/store-docker-container-images?cloudshell=true#gcloud 


#### Tag image
CLIENT: docker tag gabr0236/pokedle-client \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/pokedle-client

SERVER: docker tag gabr0236/pokedle-server \europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/pokedle-server

#### Push to gcloud repo
CLIENT: docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/pokedle-client
SERVER: docker push europe-west1-docker.pkg.dev/encoded-might-376114/pokedle-repo/pokedle-server

### DEPLOY
https://cloud.google.com/kubernetes-engine/docs/tutorials/hello-app#console_1


