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

### Start a minikube cluster

This workshop uses minikube. Minikube is a way to run kubernetes on your machine, in a VM, which is convenient for development and experimentation. Install it :

```
$> brew cask install virtualbox
$> brew cask install minikube
```

Then start the local cluster : `minikube start`.

If running into the following error `Error updating cluster:  Error updating localkube from uri: Error creating localkube asset from url:` run :

```
wget https://github.com/kubernetes/minikube/releases/download/v0.25.1/localkube && mv localkube ~/.minikube/cache/localkube/localkube-v1.9.4
```

### Kubectl

Kubectl allows you to interact with the cluster. Type `kubectl` to bring up the reference. `kubectl <command> --help` will give information about each command. 

Commands that are the most used when checking the state of an app or a cluster are :

 + `kubectl get nodes` - displays the list of nodes and their status in the cluster
 + `kubectl get pods` - displays the list of pods running in the current namespace
 + `kubectl describe <resource> <resource-name>` - displays the configuration for a resource and its status as well as some logs
 + `kubectl logs <resource-name>` - displays the logs for a resource (you can also use stern for that purpose)

More advances commands, similar to docker :

 + `kubectl exec -ti <resource-name> bash` - open a shell in a running container
 + `kubectl run -ti --image=alpine alpine sh` - run a new `alpine` container and open a shell session

### Deploying applications

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

These are defined as yaml files and define all the necessary elements for each of the resources to work such as environment variables, the number of CPUs needed, the ports to expose...

