# Dockerizing an application

Make a copy of the `examples/python` directory.

## The Dockerfile

Create a new `Dockerfile` file in that new directory. A dockerfile is a list of instructions that will be run one after the other. Each instruction run is called a layer, and each layer is cached. The cached layers will be reused unless they are changes at their level, in which case they will be rebuilt from scratch as well as all the following layers. Here is a short description of the most commonly used instructions. The full list and documentation is available at https://docs.docker.com/engine/reference/builder/.

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

This instruction is used to copy files from the working directory into the container filesystem.

```Dockerfile
COPY app.py . 
```

### The `RUN` instruction

The `RUN` instruction is used to execute commands in the container. Usually mostly used to install dependencies, libraries or systems packages.

```Dockerfile
RUN apt-get update && apt-get install -y --no-install-recommends build-essentials
```

### The `CMD` instruction

The `CMD` instruction defines the default command that will be used when the image is run. This can be overriden at runtime.

```Dockerfile
CMD python app.py
```

## Building a docker image 

```bash
$ docker build -t python-learning .
Sending build context to Docker daemon   5.12kB
Step 1/6 : FROM python:3-alpine
 ---> 29b5ce58cfbc
Step 2/6 : WORKDIR /app
 ---> Using cache
 ---> fa43cb2c2a8e
Step 3/6 : COPY requirements.txt .
 ---> Using cache
 ---> fc5afb2ea7a3
Step 4/6 : RUN pip install -r requirements.txt
 ---> Using cache
 ---> bed10fc0d87e
Step 5/6 : COPY app.py .
 ---> a63b038f6364
Step 6/6 : CMD gunicorn -b 0.0.0.0:5000 app:app
 ---> Running in cb680b9c04a0
Removing intermediate container cb680b9c04a0
 ---> cdd952787e82
Successfully built cdd952787e82
Successfully tagged python-learning:latest
```

## Running a docker image

```bash
$ docker run --rm -d -p 5000:5000 python-learning
1e9e3c5ecbd7648baa858823c294524be57a6718dcdc46253c469e377a01f35a

# We launched our container in detached mode, exposed locally on port 5000

$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED                  STATUS              PORTS                    NAMES
1e9e3c5ecbd7        python-learning     "/bin/sh -c 'gunicor…"   Less than a second ago   Up 1 second         0.0.0.0:5000->5000/tcp   vibrant_fermi

$ http localhost:5000/
HTTP/1.1 200 OK
Connection: close
Content-Length: 5
Content-Type: text/html; charset=utf-8
Date: Mon, 26 Mar 2018 06:39:24 GMT
Server: gunicorn/19.7.1

Hello

$ docker stop 1e9e3c5ecbd7
```

## Going further : optimizing your Dockerfiles

Each instruction in the file creates a layer, which is then cached. If a layer changes, it is rebuilt and so are all the following layers : place instructions that take a long time and that are unlikely to change at the beginning of the file (package installation, dependencies fetching, …).
Ideally, images stay small in order to ship them faster. Use smaller base images (slim, alpine, scratch) but those can introduce dependency problems.
If using Debian base images and apt, don’t install recommended packages (--no-install-recommends). Clear commands that cache things locally (apt again).


Next, let's [start a local kubernetes cluster](04-running-minikube.md).
