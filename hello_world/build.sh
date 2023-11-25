GOOS=js GOARCH=wasm go build -o ../server/main.wasm
cp ./index.html ../server/index.html
cd ../server/
go run ../server/main.go