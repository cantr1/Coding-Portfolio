#!/usr/bin/python3

import os
import sys
import re

PAM_MODULE_FILE = '/etc/pam.d/common-auth'

# The machine is not using GDM as a display manager. Skipping the validation.
if not os.path.isfile(PAM_MODULE_FILE):
    print("✅ GDM not a display manager: PASS")
    sys.exit(0)

with open(PAM_MODULE_FILE, 'r', encoding='utf-8') as f:
    pam_file_content = f.read()

# Check if gdm-password is correctly configured.
# There should be three lines in the gdm-password file:
# - auth requisite pam_faillock.so preauth silent deny=4 unlock_time=1200
# - auth requisite pam_faillock.so authfail deny=4 unlock_time=1200
# - auth requisite pam_faillock.so authsucc deny=4 unlock_time=1200

if not re.search(r'pam_faillock\.so preauth silent deny=(\d+) unlock_time=(\d+)',
                 pam_file_content):
    print(f'FAIL: preauth for pam_filelock.so not found in {PAM_MODULE_FILE}\n')
    sys.exit(1)
elif not re.search(r'pam_faillock\.so authfail deny=(\d+) unlock_time=(\d+)',
                 pam_file_content):
    print(f'FAIL: authfail for pam_filelock.so not found in {PAM_MODULE_FILE}\n')
    sys.exit(1)
elif not re.search(r'pam_faillock\.so authsucc deny=(\d+) unlock_time=(\d+)',
                 pam_file_content):
    print(f'FAIL: authsucc for pam_filelock.so not found in {PAM_MODULE_FILE}\n')
    sys.exit(1)
else:
    print("✅ PAM Configuration Confirmed: PASS")
    sys.exit(0)