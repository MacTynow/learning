# Dockerizing an application

Make a copy of the `examples/python` directory.

## The Dockerfile

Create a new `Dockerfile` file in that new directory.

### The `FROM` instruction

The `FROM` instruction defines the base image that will be used for a build. Base images can be any image, public or private. When looking for an image to use for an application you can look on the [docker hub](https://hub.docker.com/) or third party public registries such as [quay.io](https://quay.io/).

Images can be tagged :

```Dockerfile
# use the python 3 image as base
FROM python:3
```

Some images can use different (and lighter) OS :

```Dockerfile
# use the python alpine 3 image as base
# alpine linux is a lightweight distribution made for containers
FROM python:3-alpine
```

### The `COPY` instruction

```Dockerfile
COPY app.py . 
```

### The `RUN` instruction

```Dockerfile
RUN apt-get update && apt-get install -y --no-install-recommends build-essentials
```

### The `CMD` instruction

The `CMD` instruction defines the default command that will be used when the image is run. This can be overriden at runtime.

```Dockerfile
CMD python app.py
```

## Building a docker image 

```
docker build -t python-learning .
```

## Running a docker image

```
docker run -p 5000:5000 python-learning
```

## Going further : optimizing your Dockerfiles

Each instruction in the file creates a layer
Each layer can be cached
If a layer changes, it is rebuilt and so are all the following layers
Place instructions that take a long time and that are unlikely to change at the beginning of the file (package installation, dependencies fetching, …)
Ideally, images stay small 
Use smaller base images (slim, alpine, scratch) but those can introduce deps problems
Don’t install recommended packages (--no-install-recommends)
Clear commands that cache things locally (apt again)
