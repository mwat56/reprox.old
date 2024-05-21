#!/bin/bash
set -u -x

declare -r TARGETDIR="${PWD}.git"

git cvsimport -v -C "${TARGETDIR}" -a -A ~/.cvsnames -R -o main
[ $? -eq 0 ] && cd "${TARGETDIR}" && git symbolic-ref HEAD refs/heads/main
#_EoF_
