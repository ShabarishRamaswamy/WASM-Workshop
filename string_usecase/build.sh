GOOS=js GOARCH=wasm go build -o ../server/main.wasm -ldflags="-s -w"
cp ./index.html ../server/index.html
cd ../server/