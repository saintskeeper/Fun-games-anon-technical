# Question 1 write up 
I placed the inline policy for aws in the main.tf to show the complete workflow for creating access to the s3 bucket. 

Important block for this requirement; 
```
      {
        Effect = "Allow"
        Action = [
          "s3:GetObject",
          "s3:ListBucket",
        ],
        Resource = [
          "arn:aws:s3:::epic-unreal-infrastructure-bucket",
          "arn:aws:s3:::epic-unreal-infrastructure-bucket/*",
        ],
```

As a note this can be further restricted if need be to specific subpath or resource in the event we need to more tightly control access.

## Bonus 
I've created an `aws_iam_instance_profile` to attach the role  `MyEC2S3AccessRole` to the instance. This is a best practice to ensure that the instance has the correct permissions to access the s3 bucket.