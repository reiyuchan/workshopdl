#!/bin/bash

# Copyright (c) 2024 reiyuchan <iwakura.rei.tk@gmail.com>
# Released under the MIT license
# http://opensource.org/licenses/mit-license.php

OS=("windows" "darwin" "linux")
OSARCH=("amd64" 386)

SRC_DIR="src"
BUILD_DIR="../bin/"
WIN_EXT=".exe"
APP_NAME="workshopdl"

pwd="$(pwd)"
cwd="$(dirname "$0")"

cd $cwd/$SRC_DIR

for os in "${OS[@]}"; do
	for os_arch in "${OSARCH[@]}"; do
		if [ "$os" == "windows" ]; then
			if [ "$os_arch" == "amd64" ]; then
			GOVERSIONINFO_OPTS="-64"
			else
			GOVERSIONINFO_OPTS=""
			fi
		goversioninfo $GOVERSIONINFO_OPTS -icon ../res/icon.ico
		GOOS=$os GOARCH=$os_arch go build -ldflags "-s -w" -o "$BUILD_DIR/$APP_NAME-$os-$os_arch$WIN_EXT"
		rm -rf resource.syso
		else
		GOOS=$os GOARCH=$os_arch go build -ldflags "-s -w" -o "$BUILD_DIR/$APP_NAME-$os-$os_arch"
		fi
	done
done

cd $pwd
