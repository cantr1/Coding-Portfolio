#!/usr/bin/env python3
"""
This script cleans up an APT cache lock
Author: Kelly Cantrell (8/25/2025)
"""

import subprocess
import psutil  # external package, install with `pip install psutil`

def main():
    # Find running "apt upgrade -y" processes
    for proc in psutil.process_iter(attrs=["pid", "cmdline"]):
        cmdline = " ".join(proc.info["cmdline"])
        if "apt upgrade -y" in cmdline:
            print(f"🪦 Killing process {proc.pid}: {cmdline}")
            subprocess.run(["sudo", "kill", str(proc.pid)], check=False)

    # Fix cache afterwards
    subprocess.run(["sudo", "dpkg", "--configure", "-a"], check=False)

if __name__ == "__main__":
    main()