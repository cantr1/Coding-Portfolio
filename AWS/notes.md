# AWS Notes
This file serves as a place to document some useful commands for use with AWS.

## Commands
### Regions
`aws ec2 describe-regions --output table`
List all available regions

`aws ec2 describe-availability-zones --output table`
List availability zones

`aws configure set region us-east-1`
Set default region

`aws configure get region`
View default region