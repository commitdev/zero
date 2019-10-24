#
# VPC Resources
#  * VPC
#  * Subnets
#  * Internet Gateway
#  * Route Table
#

resource "aws_vpc" "demo" {
  cidr_block = "10.0.0.0/16"

  tags = "${
    map(
      "Name", "terraform-eks-demo-node",
      "kubernetes.io/cluster/${var.cluster-name}", "shared",
    )
  }"
}

resource "aws_subnet" "demo-0" {
  # count = 2

  availability_zone = "${data.aws_availability_zones.available.names[0]}"
  # cidr_block        = "10.0.${count.index}.0/24"
  cidr_block        = "10.0.0.0/24"
  vpc_id            = "${aws_vpc.demo.id}"

  tags = "${
    map(
      "Name", "terraform-eks-demo-node",
      "kubernetes.io/cluster/${var.cluster-name}", "shared",
    )
  }"
}

resource "aws_subnet" "demo-1" {
  # count = 2

  availability_zone = "${data.aws_availability_zones.available.names[1]}"
  # cidr_block        = "10.0.${count.index}.0/24"
  cidr_block        = "10.0.1.0/24"
  vpc_id            = "${aws_vpc.demo.id}"

  tags = "${
    map(
      "Name", "terraform-eks-demo-node",
      "kubernetes.io/cluster/${var.cluster-name}", "shared",
    )
  }"
}

resource "aws_internet_gateway" "demo" {
  vpc_id = "${aws_vpc.demo.id}"

  tags = {
    Name = "terraform-eks-demo"
  }
}

resource "aws_route_table" "demo" {
  vpc_id = "${aws_vpc.demo.id}"

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.demo.id}"
  }
}

resource "aws_route_table_association" "demo-0" {
  # count = 2

  subnet_id      = "${aws_subnet.demo-0.id}"
  route_table_id = "${aws_route_table.demo.id}"
}

resource "aws_route_table_association" "demo-1" {
  # count = 2

  subnet_id      = "${aws_subnet.demo-1.id}"
  route_table_id = "${aws_route_table.demo.id}"
}
