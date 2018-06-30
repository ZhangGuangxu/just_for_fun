@echo off

setlocal

if exist install.bat goto ok
echo install.bat must be run from its folder
goto end

:ok

set OLDGOPATH=%GOPATH%
set GOPATH=%~dp0

gofmt -w "src/q1"
go install "q1"

gofmt -w "src/minstack"
go install "minstack"

gofmt -w "src/q3"
go install "q3"

gofmt -w "src/queue_with_two_stack"
go install "queue_with_two_stack"

gofmt -w "src/q4"
go install "q4"

:end
echo finished