### build wasm
```
cd ./cmd/wasm
$env:GOOS = 'js'; $env:GOARCH = 'wasm'; go build -o  ../../assets/json.wasm // powershell
GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm  // bash
```

### run web server
```
cd ./cmd/server
go run .
```