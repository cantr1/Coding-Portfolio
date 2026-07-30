# This file contains some common git alias commands I like to use

# Returns the last commit info
git config --global alias.last 'log -1 HEAD --pretty=format:"%C(magenta)%h%Creset -%C(red)%d%Creset %s %C(dim green)(%cr) [%an]"' # => git last

# Returns the last few commits in one line style
git config --global alias.recents 'log --oneline -n 5 --pretty=format:"%C(magenta)%h%Creset -%C(red)%d%Creset %s %C(dim green)(%cr) [%an]"' # => git recents

# Lazy status
git config --global alias.st 'status' # => git st

# Even lazier status
git config --global alias.s 'status -s' # => git s

# Branch work
git config --global alias.br 'branch' # => git br -c new_branch

# Checkout a branch
git config --global alias.co 'checkout' # => git co new_branch

# To remove an alias
# git config --global --unset alias.cmd