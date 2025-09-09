#!/bin/bash
# This script must be run as root

# Step 1: Create the login banner file in /etc/ssh
BANNER_FILE="/etc/ssh/login-banner.txt"

echo "Creating login banner at $BANNER_FILE"
cat << 'EOF' > "$BANNER_FILE"
           __________                                 
         .'----------`.                              
         | .--------. |                             
         | |########| |       __________              
         | |########| |      /__________\             
.--------| `--------' |------|    --=-- |-------------.
|        `----,-.-----'      |o ======  |             | 
|       ______|_|_______     |__________|             | 
|      /  %%%%%%%%%%%%  \                             | 
|     /  %%%%%%%%%%%%%%  \                            | 
|     ^^^^^^^^^^^^^^^^^^^^                            | 
+-----------------------------------------------------+
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^ 
EOF


# Step 2: Edit /etc/ssh/sshd_config
SSHD_CONFIG="/etc/ssh/sshd_config"
echo "Modifying $SSHD_CONFIG"

# Replace "#Banner none" with our banner file path
sed -i 's/^#Banner none/Banner \/etc\/ssh\/login-banner.txt/' "$SSHD_CONFIG"

# Change "#PrintLastLog yes" to "PrintLastLog no"
sed -i 's/^#PrintLastLog yes/PrintLastLog no/' "$SSHD_CONFIG"


# Step 3: Edit /etc/pam.d/sshd to disable default MOTD
PAM_SSHD="/etc/pam.d/sshd"
echo "Modifying $PAM_SSHD to disable default MOTD sessions"

# Comment out the line: session  optional  pam_motd.so motd=/run/motd.dynamic noupdate
sed -i '/^session\s\+optional\s\+pam_motd\.so\s\+motd=\/run\/motd\.dynamic\s\+noupdate/s/^/#/' "$PAM_SSHD"

# Comment out the line: session  optional  pam_motd.so # [1]
sed -i '/^session\s\+optional\s\+pam_motd\.so\s\+# \[1\]/s/^/#/' "$PAM_SSHD"

# Comment out the line: session  optional  pam_motd.so standard noenv # [1]
sed -i '/^session\s\+optional\s\+pam_motd\.so\s\+standard\s\+noenv\s\+# \[1\]/s/^/#/' "$PAM_SSHD"


# Step 4: Create the post-login banner in /etc/motd
MOTD_FILE="/etc/motd"
echo "Creating post-login banner at $MOTD_FILE"
cat << 'EOF' > "$MOTD_FILE"
 _   _ _                 _           _           _     
| | | | |               | |         | |         | |    
| | | | |__  _   _ _ __ | |_ _   _  | |     __ _| |__  
| | | | '_ \| | | | '_ \| __| | | | | |    / _` | '_ \ 
| |_| | |_) | |_| | | | | |_| |_| | | |___| (_| | |_) |
 \___/|_.__/ \__,_|_| |_|\__|\__,_| \_____/\__,_|_.__/ 
EOF

# Step 5: Disable other login messages
sudo chmod -x /etc/update-motd.d/*

# Step 6: Restart ssh services
sudo systemctl restart ssh
