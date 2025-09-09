#!/usr/bin/env bash

set -e  # Exit immediately if a command exits with a non-zero status

SUDOERS_FILE="/etc/sudoers"
BACKUP_FILE="/etc/sudoers.bak_$(date +%Y%m%d)"
TEMP_FILE="/tmp/sudoers.new"

# 1) Make a backup
cp "${SUDOERS_FILE}" "${BACKUP_FILE}"

# 2) Use awk to process the lines
awk '
  # We will track lines that match the it_tech_mod pattern to remove duplicates
  /it_tech_mod/ {
    if (!seen_it_tech_mod[$0]++) {
      print
    }
    next
  }

  # If line contains sel-te_mod (old group) but not qmn_sel_te_mod, skip it
  /sel-te_mod/ && !/qmn_sel_te_mod/ {
    next
  }

  # Otherwise, print the line
  {
    print
  }
' "${SUDOERS_FILE}" > "${TEMP_FILE}"

# 3) Validate the new file syntax using visudo
if visudo -cf "${TEMP_FILE}" ; then
  # If syntax is OK, move it into place
  cp "${TEMP_FILE}" "${SUDOERS_FILE}"
  echo "Successfully updated ${SUDOERS_FILE}."
else
  echo "visudo check failed; restoring backup."
  cp "${BACKUP_FILE}" "${SUDOERS_FILE}"
  exit 1
fi

rm -f "${TEMP_FILE}"
