# 1. Choose a Remote Backend:
 lets use S3 for this example and a bucket named `terraform-remote-state-fun-games`.

# 2. Configure Backend in Terraform Configuration  & Bonus: 

create a dynamodb table to lock the state file to prevent concurrent writes called `your-dynamodb-lock-table`

```
terraform {
  backend "s3" {
    bucket         = "terraform-remote-state-fun-game"
    key            = "terraform.tfstate"
    region         = "your-region"
    dynamodb_table = "your-dynamodb-lock-table"
  }
}

```

## NOTES 
This effectivly locks the "workspace" when a `terraform apply` is triggered
