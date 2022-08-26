SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=amd64
go build -o UniAnalytics_windows_amd64.exe main.go

SET CGO_ENABLED=1
SET GOOS=darwin
SET GOARCH=amd64
go build -o UniAnalytics_macos_amd64 main.go

SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=amd64
go build -o UniAnalytics_linux_amd64 main.go