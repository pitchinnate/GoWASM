### build wasm in powershell
cd ./cmd/wasm
$env:GOOS = 'js'; $env:GOARCH = 'wasm'; go build -o  ../../assets/json.wasm

### run web server
cd ./cmd/server
go run .