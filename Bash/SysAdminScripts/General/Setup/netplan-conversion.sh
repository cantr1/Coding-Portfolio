#!/bin/bash
# This script grabs the needed info to 
# convert a linux server to netplan from network manager
# Written by: Kelly Cantrell

# VARIABLES
# Get the name of the interface currently in use
ACTIVE_INTERFACE=$(ip addr | grep "mq state UP" | awk -F: '{sub(/^ +/, "", $2); print $2}')

# Get the macaddress of the interface
MAC_ADDRESS=$(ip addr show dev $ACTIVE_INTERFACE | grep "ether" | awk '{print $2}')

# Get the associated IP addresses
IPV4_ADDRESS=$(ip addr show dev $ACTIVE_INTERFACE | grep -w "inet" | awk '{print $2}')
IPV6_ADDRESS=$(ip addr show dev $ACTIVE_INTERFACE | grep -w "inet6" | awk 'NR==1 {print $2; exit}')

# Get the default gateway
IPV4_GATEWAY=$(ip route | grep "default" | awk '{print $3}')


# NETPLAN IMPLEMENTATION
# Echo the config file for netplan into sudo tee to write it to the file
echo "network:
  version: 2
  renderer: networkd
  ethernets:
    ens4f0:
      match:
        macaddress: $MAC_ADDRESS
      set-name: ens4f0
      dhcp4: no
      dhcp6: no
      addresses: [$IPV4_ADDRESS, \"$IPV6_ADDRESS\"]
      routes:
        - to: 0.0.0.0/0
          via: ${IPV4_GATEWAY}
      nameservers:
        addresses: [8.8.8.8, 8.8.4.4]" | sudo tee /etc/netplan/01-network-manager-all.yaml 1>/dev/null
    
sudo systemctl disable NetworkManager
sudo systemctl enable --now systemd-netword
sudo netplan apply
sudo systemctl stop NetworkManager

