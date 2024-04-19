#!/bin/bash
kubectl cordon <node-name>


kubectl drain <node-name> --ignore-daemonsets --delete-local-data


# --ignore-daemonsets: This is necessary because DaemonSet-managed pods are automatically recreated by Kubernetes, and you don't need to manually manage these during a node drain.
# --delete-local-data: This tells Kubernetes to ignore any concerns about pods with local storage (non-persistent). Be careful with this flag, as it might lead to data loss if the pods use local storage for persistent data.


# return to service
 kubectl uncordon <node-name>