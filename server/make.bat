@echo off

set pwd=%~dp0
SET GO111MODULE=on

SET GOSUMDB=off
SET GOBIN=%pwd%/bin

if [%1] == [] goto:all

if %1==clean (
call:clean
) else if %1==proto (
call:proto %2
) else (
call:build %1
)
goto:exit

:all
call:depend
if %errorlevel%==1 (
	goto:exit
)

call:build server/src/roomserver
pause

exit /b 0

:depend
go mod tidy
if %errorlevel%==0 (
echo go mod success!
) else (
echo go mod error!
exit /b 1
)
exit /b 0

:build
go install %1
if %errorlevel%==0 (
echo build %1 success!
) else (
echo build %1 error!
)
exit /b 0

:proto

exit /b 0

:clean
rm pkg/* -rf
rm bin/*.exe -rf
echo clean ok!
exit /b 0

:exit