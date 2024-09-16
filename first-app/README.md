# Build golang app with Docker

### Docker build/update command

```
docker build . -t go-containerized:latest
```

Explain:

- docker build: command to build docker image
- .(dot): Represents the build context, which is the current directory containing the Dockerfile and other necessary files for building the image.
- -t (or --tag): the tag of the image for docker
- (go-containerized:latest): the name of the image(go-containerized) and tag(latest) for the docker image

### Docker run app

```
docker run -d --name go-app -p 8080:8080 go-containerized:latest
```

explain:

- docker run: command to run as a container
- (-d): running docker in detached mode, meaning docker run on background, the terminal is not blocked
- (--name go-app): Assigns a custom name (go-app) to the running container. Without this, Docker will assign a random name.
- (-p 8080:8080): Maps port 8080 on the local machine (left) to port 8080 in the Docker container (right). This allows traffic on your local machine at localhost:8080 to reach the container's service running on port 8080.
- (go-containerized:latest): Specifies the Docker image (go-containerized) and the tag (latest) to run the container from.

### Docker stop app

```
docker stop go-app
```

### Docker remove app

```
docker rm go-app
```

## Push docker image into docker hub

### Create repository in docker hub

### Tag the docker image

format:

```
docker tag image-name:tag username/repository:tag
```

command:

```
docker tag go-containerized:latest egrizq/first-app:latest
```

### Push the image to docker hub

format:

```
docker push username/repository:tag
```

command:

```
docker push egrizq/first-app:latest
```
