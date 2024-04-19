// profile to assume to ensure this code is deployed in us-east-1
provider "aws" {
  region = "us-east-1"
}
//  base vpc 
resource "aws_vpc" "my_vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "MyVPC"
  }
}

// internet gateway to allow public access 
resource "aws_internet_gateway" "my_igw" {
  vpc_id = aws_vpc.my_vpc.id

  tags = {
    Name = "MyInternetGateway"
  }
}

// eip nat gateway for private subnet access
resource "aws_eip" "nat_eip" {
  vpc = true
}

resource "aws_nat_gateway" "my_nat_gateway" {
  allocation_id = aws_eip.nat_eip.id
  subnet_id     = aws_subnet.public[0].id

  tags = {
    Name = "MyNATGateway"
  }
}

// aws public subnets 
// set count to 5 to iterate over each sublic subnet 
resource "aws_subnet" "public" {
  count                   = 5
  vpc_id                  = aws_vpc.my_vpc.id
  cidr_block              = cidrsubnet(aws_vpc.my_vpc.cidr_block, 4, count.index)
  map_public_ip_on_launch = true

  tags = {
    Name = "PublicSubnet-${count.index}"
  }
}

// aws private subnets
// set count to 5 to iterate over each private subnet 
resource "aws_subnet" "private" {
  count      = 5
  vpc_id     = aws_vpc.my_vpc.id
  cidr_block = cidrsubnet(aws_vpc.my_vpc.cidr_block, 4, count.index + 5)

  tags = {
    Name = "PrivateSubnet-${count.index}"
  }
}

// public route tables for public subnets 
resource "aws_route_table" "public_route_table" {
  vpc_id = aws_vpc.my_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.my_igw.id
  }

  tags = {
    Name = "PublicRouteTable"
  }
}

// assosication of the above route tables 
resource "aws_route_table_association" "public_route_table_assoc" {
  count          = 5
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public_route_table.id
}

// creation of private route tables
resource "aws_route_table" "private_route_table" {
  vpc_id = aws_vpc.my_vpc.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.my_nat_gateway.id
  }

  tags = {
    Name = "PrivateRouteTable"
  }
}

// association of private route tables
resource "aws_route_table_association" "private_route_table_assoc" {
  count          = 5
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private_route_table.id
}
