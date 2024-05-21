#!/bin/bash
# $Id$
#
#	Replace some text fragment
#
#------------------------------------------------------------------------------
#set -x

OLD='go-ini'
NEW='ini'


function doReplace() {
local f="${1}"
local n="${f}.$$"
#echo $f
	sed -e "s/${OLD}/${NEW}/g" "${f}" > "${n}"
	if [ 0 -eq $? -a -s "${n}" ] ; then
		if cmp -s "${f}" "${n}" ; then
			rm -f "${n}"
		else
			mv -fv "${n}" "${f}"
#exit 0
		fi
	fi
	[ -s "${n}" ] || rm -fv "${n}"
} # doReplace()

function doDir() {
local pwd="${PWD}"
local dir="${1}"
local f
	echo "Directory: ${dir}"
	builtin cd "${dir}" || return
	for f in * ; do
		if [ -d "${f}" ] ; then
			case "${f}" in
				'api-doc'|'CVS'|'misc'|'OLD'|'sessions'|'.git')
					:
					;;
				*)
					doDir "${f}"
					;;
			esac
		elif [ -s "${f}" ] ; then
			if [ "sed.sh" != "${f}" ]; then
#			case "${f##*.}" in
#				'phpX'|'tplX'|'html')
					doReplace "${f}"
#					;;
#				*)
#					:
#					;;
#			esac
			fi
		fi
	done
	builtin cd "${pwd}"
} # doDir()

doDir .

#_EoF_
