#!/bin/sh
#
# Change git `master` branch to `main`.
#
# 2020-06-12  Matthias Watermann  <support@mwat.de>
#-------------------------------------------------------------------------------
#set -x

git branch -m master main || exit
git push -u origin main || exit
git symbolic-ref refs/remotes/origin/HEAD refs/remotes/origin/main || exit
#
echo "
See the web-page shown above and change the default branch on github to 'main'"
read -p "press <RETURN> when done." REPLY
#
git branch -d master 2>/dev/null
git push origin --delete master
git push -u origin main --tags
echo

# _EoF_
