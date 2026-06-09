# AWS Notes
This file serves as a place to document some useful commands for use with AWS.

## Commands
### General Use
`aws login`
Login to AWS via terminal

### Regions
`aws ec2 describe-regions --output table`
List all available regions

`aws ec2 describe-availability-zones --output table`
List availability zones

`aws configure set region us-east-1`
Set default region

`aws configure get region`
View default region

### Keys
`ssh-keygen -t ed25519 -C "patientping-key" -f ~/.ssh/patientping-key`
Generate an SSH key with a comment and specify file location / name

`aws ec2 import-key-pair --key-name "patientping-key" --public-key-material fileb://$HOME/.ssh/patientping-key.pub`
Upload key to AWS

`aws ec2 describe-key-pairs`
Check Keys

### EC2
`aws ec2 describe-instances --no-cli-pager --output json`
Retrieve data on active instances
