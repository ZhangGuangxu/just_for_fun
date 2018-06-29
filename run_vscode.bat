@echo off

setlocal

if exist run_vscode.bat goto ok
echo run_vscode.bat must be run from its folder
goto end

:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

code just_for_fun.code-workspace

:end
echo finished