# Docker Configuration: 
Configure Docker Engine to use process isolation for running Windows containers. This configuration setting can typically be specified in the Docker configuration file (daemon.json). Set the default-runtime to process to enable process isolation.Example daemon.json:


# Container Networking: 
Configure networking settings for your Windows containers, including IP addresses, DNS settings, and network adapters. Ensure that containers can communicate with each other and with external resources as needed.


# Resource Allocation:
Manage resource allocation for containers, including CPU, memory, and storage resources. Configure resource limits and constraints based on your application requirements and server capabilities.


## Pros:
- Lightweight: Process isolation is generally lighter in terms of resource overhead compared to full virtualization methods like Hyper-V isolation.

- Better Performance: Since containers share the host operating system kernel, there is less overhead compared to running each container with its own kernel.

- Easier Management: Process isolation provides a simpler management experience compared to full virtualization solutions, with faster startup times and less administrative overhead.

## Cons 

- Shared Kernel: Containers share the host operating system kernel, which may introduce security risks if vulnerabilities are exploited.

- Limited Isolation: Process isolation provides less isolation compared to full virtualization methods
like Hyper-V isolation. Malicious code within a container may potentially affect other containers on the same host.

- Compatibility Issues: Some applications may not be fully compatible with process isolation due to dependencies or system requirements. Compatibility testing is essential before migrating applications to containers.



# Comparison with Hyper-V Isolation:
- Isolation Level: Hyper-V isolation provides stronger isolation by running containers in separate lightweight VMs, offering better security and isolation compared to process isolation.

- Resource Overhead: Hyper-V isolation typically incurs higher resource overhead compared to process isolation due to the additional layer of virtualization.

- Compatibility: Hyper-V isolation may be required for certain applications with specific compatibility or security requirements. However, it may introduce additional complexity and management overhead compared to process isolation.


# what i would do instead.

Instead of dealing with windows isolation, I'd  use EKS windows AMI  nodes with a windows workload to handle resource isolation, it's cheaper and more effective. 


# docs on this 
https://learn.microsoft.com/en-us/virtualization/windowscontainers/manage-containers/hyperv-container
