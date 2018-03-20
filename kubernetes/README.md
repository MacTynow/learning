# Python examples

## 1. Dockerfile 

Build the container

```
docker build -t python-learning .
```

Run the container, exposing it on local port 5000

```
docker run -p 5000:5000 python-learning
```

## 2. Running in kubernetes

This service is an API which will be accessed through an endpoint which means that there are at least 3 mandatory resources that need to be provisioned in kubernetes : 

 + Deployment - provisions a pod template, allowing for easy horizontal scaling 
 + Service - exposes the pods in the cluster, takes care of the networking part
 + Ingress - exposes the service through a reverse proxy, making it accessible through an endpoint to external services

Some optional resources an also be added such as :

 + Configmap - can hold configuration files or variables
 + Horizontal Autoscaler - scales a deployment automatically based on CPU and memory usage

If role-based access control (RBAC) is enabled on a cluster you also need to make sure that the following exist :
 
 + Cluster role
 + Cluster role binding
