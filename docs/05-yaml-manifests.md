# Writing the yaml manifests to deploy an application

In the previous section, we've seen how to start pods via the `kubectl run` command. This doesn't scale and becomes quite confusing when running more complex configurations. We can then write those configurations as yaml manifests that can be commited to version control.

## Deploy the manifest

```bash
kubectl create -f file.yaml
```


THIS SHOULD GO IN THE POWERPOINT
## Pod definition

## The service

The service resource is used to expose pods on the network.

### ClusterIP

This service type exposes the pods on an IP in the cluster range, using the port defined in the pod definition.

### NodePort

This service type exposes the application on the same port on every node accross the cluster (the default port range is 30000 to 32?33000).

### LoadBalancer

This service type interacts with the cloud provider to provision a load balancer in your account and attaches the nodes to it, exposing the pods on a NodePort.

Next, let's [easily install popular applications](06-installing-charts.md).
