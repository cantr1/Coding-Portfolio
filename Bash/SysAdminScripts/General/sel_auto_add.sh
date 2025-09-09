#!/bin/bash

#This script is meant to sit on the Rundeck server and to be triggered by the managed server script. 
#The managed server updates the file locations below, then startes this script based on the SEL selection.

# Written By: Braxton Acheson

#Set script static variables:
#Hostfile Location
hostfile_loc=/etc/ansible/sel-hosts

#Variables set to the temp file IPs/Hostnames dumped by the Rundeck-Client_AUTO.sh script
client_ip=$(head -1 /etc/ansible/tmp/IPADDRESS.txt)
client_hostname=$(head -1 /etc/ansible/tmp/HOSTNAME.txt)

#Static User Password
userpasswd=

#Rundeck/Ansible Key Locations
qmnasnible_key_loc="/home/qmnansible/.ssh/id_rsa.pub"
rundeck_key_loc="/home/qmnansible/.ssh/rundeckid_rsa.pub"

#Start of Script
echo "This script adds the IP and Hostname to the SEL Hostfile in the Default Group"

#Copy Keys to the Client system
echo "Copying the Keys for Rundeck to the Client Server"
sshpass -p "$userpasswd" -S ssh-copy-id -o StrictHostKeyChecking=accept-new -i $qmnasnible_key_loc qmnansible@$client_ip
sshpass -p "$userpasswd" -S ssh-copy-id -o StrictHostKeyChecking=accept-new -i $rundeck_key_loc rundeck@$client_ip

#Adds the Hostname and IP into the proper Host File from the temporary files.
echo "Updating the SEL Hostfile now: "
echo "$client_hostname ansible_host=$client_ip" | tee -a $hostfile_loc

#Restart the Rundeck Client to re-populate the list of nodes:
sudo systemctl restart rundeckd
sleep 5

#Ansible Ping the server for validation
ansible -i /etc/ansible/sel-hosts -m ping "$client_hostname"