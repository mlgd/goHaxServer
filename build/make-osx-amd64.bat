@ECHO OFF
SETLOCAL
SET GOOS=darwin
SET GOARCH=amd64

cd ..
go build -o goHaxServer.app main.go

ENDLOCAL
