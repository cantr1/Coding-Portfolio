#!/bin/bash
# This script is used to help with using SSH to targeted locations.
# Works best with keys in place to smoothly automate SSH connections.

MAP_FILE="/etc/target-hosts"

if [[ -z "$1" ]]; then
    echo "Usage: SSH <location>"
    exit 1
fi

LOCATION="$1"
IP=$(awk -v loc="$LOCATION" '$1 == loc {print $2}' "$MAP_FILE")

if [[ -z "$IP" ]]; then
    echo "Error: Location not found"
    exit 1
fi

if [[ "$1" == *"X"* ]]; then
    exec "ssh user@$IP"
elif [[ "$1" == *"Y"* ]]; then
    exec "ssh user@$IP"
else
    echo "Host ID not compatable."
fi
