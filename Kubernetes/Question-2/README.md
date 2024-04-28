#  Cordon 
Cordoning the node prevents new pods from being scheduled on the node 

# Drain

Draining the node gracefully terminates the pods running on the node and reschedules them on other nodes.

# Get pod Status 
Check pods after draining to make sure you've got enough  compute to handle the load.
If not will need to scale up the cluster to handle the load.


# troubleshoot failign node
hop on the node and check system logs && statuc containers for any issues.

# Return the Node to Service

Once the node is back up and running, uncordon the node to allow new pods to be scheduled on it.



# For a production cluster 
Be sure to scale up the nodegroups prior to draining to ensure there's no outage prior to draining the node.