resource "aws_iam_role" "ec2_s3_access_role" {
  name = "MyEC2S3AccessRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      },
    ]
  })
}

resource "aws_iam_role_policy" "ec2_s3_access_policy" {
  name   = "MyS3AccessPolicy"
  role   = aws_iam_role.ec2_s3_access_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
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
      },
    ]
  })
}

resource "aws_iam_instance_profile" "ec2_instance_profile" {
  name = "MyInstanceProfile"
  role = aws_iam_role.ec2_s3_access_role.name
}

resource "aws_instance" "myec2instance" {
  # Add your usual EC2 instance configuration here

  iam_instance_profile = aws_iam_instance_profile.ec2_instance_profile.name

  depends_on = [
    aws_iam_role_policy.ec2_s3_access_policy
  ]
}
