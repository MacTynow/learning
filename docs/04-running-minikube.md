# Running Minikube

## Start a minikube cluster

https://kubernetes.io/docs/getting-started-guides/minikube/

Start the local cluster : `minikube start`.

If running into the following error `Error updating cluster:  Error updating localkube from uri: Error creating localkube asset from url:` run :

```
wget https://github.com/kubernetes/minikube/releases/download/v0.25.1/localkube && mv localkube ~/.minikube/cache/localkube/localkube-v1.9.4
```

```bash
$ minikube start
Starting local Kubernetes cluster...
Running pre-create checks...
Creating machine...
Starting local Kubernetes cluster...
```

## Kubectl 

```bash
$ kubectl run hello-minikube --image=k8s.gcr.io/echoserver:1.4 --port=8080
deployment "hello-minikube" created

$ kubectl expose deployment hello-minikube --type=NodePort
service "hello-minikube" exposed

# We have now launched an echoserver pod but we have to wait until the pod is up before curling/accessing it
# via the exposed service.
# To check whether the pod is up and running we can use the following:

$ kubectl get pod
NAME                              READY     STATUS              RESTARTS   AGE
hello-minikube-3383150820-vctvh   0/1       ContainerCreating   0          3s

# We can see that the pod is still being created from the ContainerCreating status

$ kubectl get pod
NAME                              READY     STATUS    RESTARTS   AGE
hello-minikube-3383150820-vctvh   1/1       Running   0          13s

# We can see that the pod is now Running and we will now be able to curl it:

$ curl $(minikube service hello-minikube --url)
CLIENT VALUES:
client_address=192.168.99.1
command=GET
real path=/
...

$ kubectl delete deployment hello-minikube
deployment "hello-minikube" deleted

$ minikube stop
Stopping local Kubernetes cluster...
Stopping "minikube"...
```


Next, let's [define our application in yaml files](05-yaml-manifests.md).
