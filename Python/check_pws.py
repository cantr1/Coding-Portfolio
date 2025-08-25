#!/usr/bin/env python3
import subprocess
import sys

def run(cmd: list[str]) -> str:
    """
    This function takes a command as argument and returns the stdout
    :param cmd: list[str] - takes a command to run as argument and captures stdout
    :returns res.stdout: str - returns the output of the command
    """
    res = subprocess.run(cmd, capture_output=True, text=True, check=True)
    return res.stdout

def login_users() -> set[str]:
    """
    This function finds all active users and returns them in a set
    :returns users: set - all users that have an active shell above UID 1000
    """
    # uid>=1000 and shell not nologin/false
    out = run(["getent", "passwd"])
    users = set()
    for line in out.splitlines():
        name, pw, uid, gid, gecos, home, shell = line.split(":", 6)
        try:
            uid = int(uid)
        except ValueError:
            continue
        if uid >= 1000 and not shell.endswith(("nologin", "false")):
            users.add(name)
    return users

def users_without_hash(users: set[str]) -> list[str]:
    """
    This function checks that all users are in shadow with a hashed password (SHA-512)
    :param users: set - set of all users to check
    :returns nohash: list - all users without a hashed password
    """
    out = run(["getent", "shadow"])
    nohash = []
    for line in out.splitlines():
        name, hash_field, *_ = line.split(":")
        if name in users:
            # Require SHA512 specifically
            if not hash_field.startswith("$6$"):
                nohash.append(name)
    return nohash

def users_without_pw(users: set[str]) -> list[str]:
    """
    This function checks that all users have a password
    :param users: set - all users to check
    :returns nopass: list - users without a password
    """
    out = run(["getent", "passwd"])
    nopass = []
    for line in out.splitlines():
        name, pw_field, *_ = line.split(":")
        if name in users:
            if not pw_field == 'x':
                nopass.append(name)
    return nopass

def main():
    users = login_users()
    bad_hash = users_without_hash(users)
    bad_pass = users_without_pw(users)
    result = 0
    if bad_pass:
        print("❌ Accounts without a usable password:", ", ".join(sorted(bad_pass)))
        result = 1
    if bad_hash:
        print("❌ Accounts without a usable password hash:", ", ".join(sorted(bad_hash)))
        result = 1

    if result == 0:
        print("✅ All login users have a password and appropriate hash.")
    
    sys.exit(result)

if __name__ == "__main__":
    main()
