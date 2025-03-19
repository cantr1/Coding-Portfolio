#!/bin/bash
# This script grabs the needed info to 
# establish the USB/RJ45 adapter to DHCP setup and appends to the netplan config
# Written by: Kelly Cantrell

# VARIABLES
# Get the name of the interface
usb_interface=$(ip addr | grep "enx00" | awk -F: '{sub(/^ +/, "", $2); print $2}')

# Get the macaddress of the interfaces
usb_mac=$(ip addr show dev $usb_interface | grep "ether" | awk '{print $2}')


# NETPLAN IMPLEMENTATION
# Echo the config file for netplan into sudo tee to write it to the file
echo "
    ens3f0:
      match:
        macaddress: $usb_mac
      set-name: ens3f0
      dhcp4: no
      dhcp6: no
      addresses:
      routes:
        - to: 0.0.0.0/0
          via: 192.168.1.1
      nameservers:
        addresses: [8.8.8.8, 8.8.4.4]" | sudo tee -a /etc/netplan/01-network-manager-all.yaml 1>/dev/null	