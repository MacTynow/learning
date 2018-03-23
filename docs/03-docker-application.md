# Dockerizing an application

Make a copy of the `examples/python` directory.

## The Dockerfile

Create a new `Dockerfile` file in that new directory.

### The `FROM` statement

### The `COPY` statement

### The `RUN` statement

### The `CMD` statement

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
