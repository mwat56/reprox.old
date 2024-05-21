#!/usr/bin/env bash
#
# see: https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04
#
# You might need to run
#		sudo apt-get install gcc-multilib g++-multilib upx-ucl
# to avoid compilation errors.
#------------------------------------------------------------------------------
set -u #-x

if [ 1 -gt $# ] ; then
	echo -e "\n\tusage: $0 <package-name> [upx]\n"
	exit 1
fi

package=$1
package_split=(${package//\// })
package_name=${package_split[-1]}

RACE='-race'
UPX=''
if [ 1 -lt $# ] ; then
	# check whether we have an UPX executable
	UPX=$(type -p upx-ucl)
fi
[ -z "${UPX}" ] || RACE=''

# create a bin subdirectory if it doesn't exist
OUTDIR='./bin'
mkdir -pv "${OUTDIR}"

#platforms=("windows/amd64" "windows/386" "darwin/amd64")
platforms=("linux/amd64" "linux/386")

for platform in "${platforms[@]}"; do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	PROGRAM="${package_name%*.go}"
	output_name="${PROGRAM}-${GOOS}-${GOARCH}"
	if [ "windows" = "${GOOS}" ]; then
		output_name+='.exe'
	fi

	env GOOS="${GOOS}" GOARCH="${GOARCH}" \
		go build ${RACE} -ldflags="-s -w" -v -o "${OUTDIR}/${output_name}" "${package}"
	if [ 0 -eq $? ]; then
		if [ ! -z "${UPX}" ] ; then
			time /opt/bin/upx --lzma "${OUTDIR}/${output_name}"
		fi
	else
		echo 'An error has occurred! Aborting script execution ...'
		exit 1
	fi
done

#_EoF_
