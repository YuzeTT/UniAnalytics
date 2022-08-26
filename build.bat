SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o UniAnalytics_windows_amd64.exe main.go

SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o UniAnalytics_linux_amd64 main.go