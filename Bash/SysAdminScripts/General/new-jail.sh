#!/bin/bash
# This is less a script than it is a process flow.
# These commands work to set up a chroot jail on a Ubuntu server
# Written By: Kelly Cantrell

# 1. Create Base Directory and Structure
sudo mkdir -p /home/Jails
sudo mkdir -p /home/Jails/{bin,dev,etc,home,lib,usr,usr/bin,usr/lib,lib64}
sudo mkdir -p /home/Jails/{proc,sys,run,tmp}

# Make sure tmp is writable
sudo chmod 1777 /home/Jails/tmp


# 2. Set Up Device Nodes
sudo mknod -m 666 /home/Jails/dev/null c 1 3
sudo mknod -m 666 /home/Jails/dev/tty c 5 0
sudo mknod -m 666 /home/Jails/dev/zero c 1 5
sudo mknod -m 666 /home/Jails/dev/random c 1 8
sudo mknod -m 666 /home/Jails/dev/ptmx c 5 2


# Ensure the jail root is owned by root (chroot directories must not be writable by non-root users)
sudo chown -R root:root /home/Jails
sudo chmod 0755 /home/Jails

# Psuedo Terminals
sudo mkdir -p /home/Jails/dev/pts
sudo mount --bind /dev/pts /home/Jails/dev/pts
echo "/dev/pts /home/Jails/dev/pts none bind 0 0" | sudo tee -a /etc/fstab


# 3. Copy Essential Binaries
sudo cp /usr/bin/ssh /home/Jails/usr/bin/
sudo cp /usr/bin/ping /home/Jails/usr/bin/
sudo cp /usr/bin/bash /home/Jails/bin/bash

# Set setuid on ping for proper privileges
sudo chmod u+s /home/Jails/usr/bin/ping

# Bind mount /lib and /usr/lib for required libraries
sudo mount --bind /lib /home/Jails/lib
sudo mount --bind /usr/lib /home/Jails/usr/lib

#Binding the /lib64 dir
sudo mount --bind /lib64 /home/Jails/lib64

# Add persistent bind mounts to /etc/fstab
echo "/lib /home/Jails/lib none bind 0 0" | sudo tee -a /etc/fstab
echo "/lib64 /home/Jails/lib64 none bind 0 0" | sudo tee -a /etc/fstab
echo "/usr/lib /home/Jails/usr/lib none bind 0 0" | sudo tee -a /etc/fstab

# 5. Bind Mount Configuration Files for Dynamic LDAP
sudo cp /etc/nsswitch.conf /home/Jails/etc/

# Copy user and group files for local users
sudo cp /etc/passwd /home/Jails/etc/passwd
sudo cp /etc/group /home/Jails/etc/group

# 6. Set Up SSH Client Configuration inside the Jail
sudo mkdir -p /home/Jails/etc/ssh
sudo cp /etc/ssh/ssh_config /home/Jails/etc/ssh/

# 7. Configure SSHD for the Chroot Jail
# Append the chroot configuration for the designated user (replace "ssh-ituser" with your user)
sudo tee -a /etc/ssh/sshd_config <<EOF

Match Group qmn_sel_te_mod,it_tech_mod
ChrootDirectory /home/Jails
EOF

# Restart the SSH service to apply changes
sudo systemctl restart ssh


# 8. Copy Authentication Libraries (NSS and PAM)
for lib in libnss_files.so libnss_compat.so libnss_dns.so libnss_ldap.so libpam.so; do
    find /lib /usr/lib -name "$lib*" -exec sudo cp -v '{}' /home/Jails/lib/ \;
done

# Make sssd directory
sudo mkdir -p /home/Jails/etc/sssd

# Copy files
sudo cp /etc/sssd/sssd.conf /home/Jails/etc/sssd
# Permissions
sudo chmod 600 /home/Jails/etc/sssd/sssd.conf

# Create pam directory
sudo mkdir -p /home/Jails/etc/pam.d

# Copy files
sudo cp /etc/pam.d/sshd /home/Jails/etc/pam.d/

# Create mc and pipes directories
#sudo mkdir -p /home/Jails/var/lib/sss/{mc,pipes}

# Edit /etc/sssd/sssd.conf
sudo tee -a /etc/sssd/sssd.conf <<EOF

[nss]
allow_all_chrooted = true
EOF

# Bind the sss dir
sudo mkdir -p /home/Jails/var/lib/sss
sudo mount --bind /var/lib/sss /home/Jails/var/lib/sss
echo "/var/lib/sss /home/Jails/var/lib/sss none bind 0 0" | sudo tee -a /etc/fstab

# Create fstab entry to ensure persistent mount
#echo "/var/lib/sss/pipes/ /home/Jails/var/lib/sss/pipes/ none bind 0 0" | sudo tee -a /etc/fstab

# Restart sssd
sudo systemctl restart sssd

# Edit ssh known hosts for a shared file
sudo mkdir -p /home/Jails/var/shared_ssh

#Perms
sudo chmod 777 /home/Jails/var/shared_ssh/

#Edit /home/Jails/etc/ssh/ssh_config
sudo tee -a /home/Jails/etc/ssh/ssh_config <<EOF
    UserKnownHostsFile /var/shared_ssh/known_hosts
EOF


sudo systemctl restart sssd
sudo systemctl restart sssd

echo "Chroot jail setup complete."
