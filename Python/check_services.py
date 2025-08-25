#!/usr/bin/env python3
# This script pulls all services enabled to start at boot
# Written by: Kelly Cantrell (8/21/2025)
import subprocess

# Command to collect unit files with their enablement status
CMD = ["systemctl", "list-unit-files", "--type=service", "--state=enabled"]

# Run the command
services = subprocess.run(CMD, capture_output=True, text=True, check=True)

# Split into lines and skip headers/footers
lines = services.stdout.strip().splitlines()

# Enabled services list
enabled_services = []

for line in lines:
    # Output looks like: "sshd.service                enabled"
    parts = line.split()
    if len(parts) >= 2 and parts[-1] == "enabled":
        enabled_services.append(parts[0])

# Print results
print("Services Enabled at Boot")
print("========================")
for svc in enabled_services:
    if 'opentest' in svc:
        pass
    else:
        print(svc)
