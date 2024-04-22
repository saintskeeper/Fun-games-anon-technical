# lifecycle policies 
You can use the lifecycle meta-argument in Terraform to prevent the deletion of the existing security group before the new security group is created. Specifically, you can use the prevent_destroy attribute within the lifecycle block to prevent Terraform from attempting to delete the existing security group during the apply process.

## example 
``` 
resource "aws_security_group" "example" {
  # Define your security group configuration here

  lifecycle {
    prevent_destroy = true
  }
}
```

By setting prevent_destroy to true, Terraform will prevent the deletion of the existing security group during the apply process. This allows Terraform to first create the new security group with the desired name, attach it to the network adapter, and then optionally delete the old security group afterwards.