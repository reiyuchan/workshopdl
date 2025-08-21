:: Copyright (c) 2024 reiyuchan <iwakura.rei.tk@gmail.com>
:: Released under the MIT license
:: http://opensource.org/licenses/mit-license.php

@echo off

set pwd=%cd%
cd /d %~dp0

set OS=windows linux darwin
set OSARCH=amd64 386

set SRC_DIR=src
set BUILD_DIR=..\bin
set WIN_EXT=.exe
set APP_NAME=app

cd "%SRC_DIR%"
for %%o in (%OS%) do (
	for %%a in (%OSARCH%) do (
		set GOOS=%%o
		set GOARCH=%%a
		if "%%o"=="windows" (
			if "%%a" == "amd64" (
			set GOVERSIONINFO_OPTS=-64
			) else (
			set GOVERSIONINFO_OPTS=
			)
		goversioninfo.exe %GOVERSIONINFO_OPTS% -icon ..\res\icon.ico
		go build -ldflags "-s -w" -o "%BUILD_DIR%\%APP_NAME%-%%o-%%a%WIN_EXT%"
		del resource.syso
		)
		else (
		go build -ldflags "-s -w" -o "%BUILD_DIR%\%APP_NAME%-%%o-%%a"
		)
	)
)

cd /d %pwd%
