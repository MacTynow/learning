# Writing the yaml manifests to deploy an application

In the previous section, we've seen how to start pods via the `kubectl run` command. This doesn't scale and becomes quite confusing when running more complex configurations. We can then write those configurations as yaml manifests that can be commited to version control.

## Deploy the manifest

```bash
kubectl create -f file.yaml
```


Next, let's [easily install popular applications](06-installing-charts.md).
