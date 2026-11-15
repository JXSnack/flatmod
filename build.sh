CGO_ENABLED=0

# Windows
GOOS=windows GOARCH=amd64 go build -o "dist/fmod-windows-amd64.exe" -buildvcs=false
GOOS=windows GOARCH=arm64 go build -o "dist/fmod-windows-arm64.exe" -buildvcs=false

# Linux
GOOS=linux GOARCH=amd64 go build -o "dist/fmod-linux-amd64" -buildvcs=false
GOOS=linux GOARCH=arm64 go build -o "dist/fmod-linux-arm64" -buildvcs=false

# macOS
GOOS=darwin GOARCH=amd64 go build -o "dist/fmod-macos-amd64" -buildvcs=false
GOOS=darwin GOARCH=arm64 go build -o "dist/fmod-macos-arm64" -buildvcs=false