# 2017-kubecon-eu

This repository contains code for KubeCon EU 2017 that implements a simple controller that
watches for a [Third Party Resource](https://kubernetes.io/docs/user-guide/thirdpartyresources/)
called `Backup`. When the controller receives any event on `Backup`, it fetches the state of 
some part of the cluster in which it's running and logs it to the console.

See below for usage instructions.

## Install the Watcher

Install the watcher to your cluster with the following command:

```console
helm install --name=kubecon-watcher --namespace=default --set docker.tag=c0ffcab ./charts/watcher
```

Note that the watcher may restart a few times and even go into `CrashLoopBackOff` until the 
`Backup` third party resource definition it relies upon is created. This behavior is expected, and
good practice.

## Install a `Backup`

To tell the watcher to back up all pods, create a `Backup` in the cluster:

```console
helm install --name=kubecon-backup --namespace=default ./charts/backup
```

## Clean Up

Simply run the following command to clean up:

```console
helm delete --purge kubecon-watcher kubecon-backup
```
