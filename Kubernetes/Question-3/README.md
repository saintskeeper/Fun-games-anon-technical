# High level thoughts 
This setup leverages Vault's ability to securely manage and inject secrets directly into your applications running in Kubernetes, ensuring that secrets are never exposed in logs or plain text configurations.

Polcies and roles are created in vault to ensure that only the correct service accounts have access to the secrets. This prevents unauthorized access to the secrets by other tenents of a k8s cluster.


# 1. Set up vault 
## Step 1: Set Up HashiCorp Vault

```
helm repo add hashicorp https://helm.releases.hashicorp.com
helm install vault hashicorp/vault --set "server.dev.enabled=true"

```


## Step 2: Initialize and Unseal Vault:
kubectl exec -ti vault-0 -- vault operator init
kubectl exec -ti vault-0 -- vault operator unseal


# 2 Configure Kubernetes Authentication

## Step 1: Enable Kubernetes Authentication:
```
kubectl exec -ti vault-0 -- vault auth enable kubernetes

```

## Step 2: Configure Vault to communicate with Kubernetes:
```
VAULT_HELM_SECRET_NAME=$(kubectl get secrets --output=json | jq -r '.items[].metadata | select(.name|startswith("vault-token")).name')
TOKEN_REVIEW_JWT=$(kubectl get secret $VAULT_HELM_SECRET_NAME --output='go-template={{ .data.token }}' | base64 --decode)
KUBE_CA_CERT=$(kubectl config view --raw --minify --flatten --output='jsonpath={.clusters[].cluster.certificate-authority-data}' | base64 --decode)
KUBE_HOST=$(kubectl config view --raw --minify --flatten --output='jsonpath={.clusters[].cluster.server}')

kubectl exec -ti vault-0 -- vault write auth/kubernetes/config \
  token_reviewer_jwt="$TOKEN_REVIEW_JWT" \
  kubernetes_host="$KUBE_HOST" \
  kubernetes_ca_cert="$KUBE_CA_CERT"

```


# Step 3: Create and Configure Policies and Roles in Vault

## Step 1: Write a Policy for Your Application:
```
kubectl exec -ti vault-0 -- vault policy write myapp - <<EOF
path "secret/data/myapp/*" {
  capabilities = ["read"]
}
EOF

```

## Step 2: Create a Role that Maps Kubernetes Service Accounts to Vault Policies:
```
kubectl exec -ti vault-0 -- vault write auth/kubernetes/role/myapp \
  bound_service_account_names=myapp-sa \
  bound_service_account_namespaces=default \
  policies=myapp \
  ttl=1h

```


# Step 4: Use Vault Secrets in Your Kubernetes Application
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  template:
    metadata:
      annotations:
        vault.hashicorp.com/agent-inject: "true"
        vault.hashicorp.com/role: "myapp"
        vault.hashicorp.com/agent-inject-secret-credentials.txt: "secret/data/myapp/config"
        vault.hashicorp.com/agent-inject-template-credentials.txt: |
          {{- with secret "secret/data/myapp/config" -}}
            export API_KEY="{{ .Data.data.api_key }}"
          {{- end }}
    spec:
      serviceAccountName: myapp-sa
      containers:
      - name: myapp
        image: myapp-image
```





