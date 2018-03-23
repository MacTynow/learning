# Installing the tools

## Installing docker

Follow the instructions [here](https://store.docker.com/editions/community/docker-ce-desktop-mac)

## Installing kubectl 

Kubectl allows you to interact with the cluster. Type `kubectl` to bring up the reference. `kubectl <command> --help` will give information about each command. 

```
brew install kubectl
```

## Installing helm 

Helm allows you to deploy and manage applications in a cluster, similar to a package manager. It offers templating and simple configuration at runtime. A wide range of popular software is ready to install out of the box.

```
brew install kubernetes-helm
```

## Installing Virtualbox

Virtualbox is a hypervisor. It is used run the virtual machines used by minikube.

```
brew cask install virtualbox
```

## Installing minikube

Minikube is a way to run kubernetes on your machine, in a VM, which is convenient for development and experimentation.

```
brew cask install minikube
```

## Optional : installing stern

Stern is a tool that allows you to tail logs from pods running in a kubernetes cluster. It supports regex.

```
brew install stern
```

Next, [installing the client tools](03-docker-application.md).
