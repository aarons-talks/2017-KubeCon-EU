# 2017-kubecon-eu

Code for KubeCon EU 2017

Install the watcher to your cluster with the following command:

```console
helm install --name=kubecon-watcher --namespace=default ./charts/watcher
```

Then, tell the watcher to back up all pods with the following command:

```console
helm install --name=kubecon-backup --namespace=default --set docker.tag=c0ffcab ./charts/backup
```

To clean up:

```console
helm delete --purge kubecon-watcher kubecon-backup
```
