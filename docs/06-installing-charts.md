# Installing a stable helm chart

Plain yaml manifests are convenient, but when we need to redeploy an application often, or if we want to manage our application versions, they are quite lacking. This where Helm comes in : it allows to write templated yaml manifests, also known as charts.

## Browsing the curated charts

A lot of popular applications already have charts written that can simply be installed using `helm install stable/nginx` for example. Browse the entire list at https://github.com/kubernetes/charts.

## Installing the ingress controller

Let's install an nginx ingress controller. This is basically a reverse proxy which interacts with ingress resources to allow external access to our applications. The definition can be find at https://github.com/kubernetes/charts/tree/master/stable/nginx-ingress.

```bash
$ helm install stable/nginx-ingress
```

This installs nginx with the default values https://github.com/kubernetes/charts/blob/master/stable/nginx-ingress/values.yaml. All these can be overridden on the command line or on using a file.


Next, let's [write a chart for our application](07-writing-chart.md).
