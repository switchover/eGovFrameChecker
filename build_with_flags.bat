@echo off

set TARGET=github.com/switchover/eGovFrameChecker/cmd/ver

set CUR_HH=%time:~0,2%
if %CUR_HH% lss 10 (set CUR_HH=0%time:~1,1%)
set CUR_NN=%time:~3,2%
set CUR_SS=%time:~6,2%
set CUR_MS=%time:~9,2%
set BUILD_TIME=%date% %CUR_HH%:%CUR_NN%:%CUR_SS%.%CUR_MS%

FOR /F "tokens=*" %%g IN ('go version') do (SET GO_VERSION=%%g)
call set GO_VERSION=%%GO_VERSION:go version =%%

set "GIT_CMD=git log -1 --pretty^=format:%%h"
FOR /F "tokens=*" %%g IN ('%GIT_CMD%') do (SET COMMIT_HASH=%%g)

echo Build time : %BUILD_TIME%
echo Go version : %GO_VERSION%
echo Commit hash : %COMMIT_HASH%

set FLAG=-X '%TARGET%.BuildTime=%BUILD_TIME%'
set FLAG=%FLAG% -X '%TARGET%.GoVersion=%GO_VERSION%'
set FLAG=%FLAG% -X '%TARGET%.CommitHash=%COMMIT_HASH%'

go build -v -ldflags "%FLAG%" %1 %2 %3 %4 %5 %6 %7 %8 %9
