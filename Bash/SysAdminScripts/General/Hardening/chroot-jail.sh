#!/bin/bash
# Script to make chroot jail

# Create the ssh-user
sudo useradd -m -s /bin/bash -G Jumper ssh-user && echo "ssh-user:password" | sudo chpasswd

# Make jails Dir
sudo mkdir -p /home/Jails
	
# Make file system
sudo mkdir -p /home/Jails/dev
sudo mkdir -p /home/Jails/home
sudo mkdir -p /home/Jails/etc
sudo mkdir -p /home/Jails/bin
sudo mkdir -p /home/Jails/lib64
sudo mkdir -p /home/Jails/lib
sudo mkdir -p /home/Jails/usr
sudo mkdir -p /home/Jails/usr/bin

# Create the /dev file system
sudo mknod -m 666 /home/Jails/dev/null c 1 3
sudo mknod -m 666 /home/Jails/dev/tty c 5 0
sudo mknod -m 666 /home/Jails/dev/zero c 1 5
sudo mknod -m 666 /home/Jails/dev/random c 1 8

#Change ownership and permissions
sudo chown root:root /home/Jails
sudo chmod 0755 /home/Jails

# Copy binaries
sudo cp /bin/bash /home/Jails/bin/
sudo cp /usr/bin/bash /home/Jails/usr/bin/

# Binding the /lib directory
sudo mount --bind /lib /home/Jails/lib
#sudo mount -o remount,ro /home/Jails/lib
#The above did not work

#Binding the /lib64 dir
sudo mount --bind /lib64 /home/Jails/lib64
#sudo mount -o remount,ro /home/Jails/lib64
#The above did not work

# Ensure persistent mounts
echo "
/lib /home/Jails/lib none bind 0 0
/lib64 /home/Jails/lib64 none bind 0 0" | sudo tee -a /etc/fstab

# Symlink the /etc/passwd and /etc/group to the Jail
sudo ln -s /etc/passwd /home/Jails/etc/passwd
sudo ln -s /etc/group /home/Jails/etc/group

# Edit the ssh config
echo "
# Define username to apply chroot jail to
Match User ssh-user
# Specify chroot jail
ChrootDirectory /home/Jails" | sudo tee -a /etc/ssh/sshd_config

# Restart the ssh service
sudo systemctl restart ssh

# Copy the necessary binaries
sudo cp -v /usr/bin/ssh /home/Jails/usr/bin/
sudo cp -v /usr/bin/telnet /home/Jails/usr/bin/
sudo cp -v /usr/bin/cat /home/Jails/usr/bin/
sudo cp -v /usr/bin/less /home/Jails/usr/bin/
sudo cp -v /usr/bin/ls /home/Jails/usr/bin/
# sudo cp -v /usr/bin/systemctl /home/Jails/usr/bin/
# systemd will not work inside the Jail

# Change ownership of ssh-user files
sudo chown ssh-user:ssh-user /home/ssh-user
sudo chown ssh-user:ssh-user /home/ssh-user/.bash_logout
sudo chown ssh-user:ssh-user /home/ssh-user/.bashrc
sudo chown ssh-user:ssh-user /home/ssh-user/.cache
sudo chown ssh-user:ssh-user /home/ssh-user/.config
sudo chown ssh-user:ssh-user /home/ssh-user/.profile
sudo chown ssh-user:ssh-user /home/ssh-user/snap
# No .cache, .config, or snap
