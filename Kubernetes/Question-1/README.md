# k8s daemonsets 

## Node Labels 
1. add linux based labels to the linux nodes 
2. add windows based labels to the windows nodes


##  Deployment options for logging aggrigation 
1. DaemonSet: Ensures that a copy of the pod runs on each node in the cluster.
2. Node Affinity: Targets the deployment only to Linux nodes.
3. Tolerations: Allows the pod to be scheduled on nodes with matching taints, if any.

