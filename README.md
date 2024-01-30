# WASM-Workshop
The code for the WASM Tech-Tok conducted for ConnectWise on 17th January 2024 and much more.

### NOTE: This is a WIP repository.

## Run the individual examples
1. cd into any example directory.
2. cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
3. GOOS=js GOARCH=wasm go build -o ./main.wasm
4. python3 -m http.server [This is to start the http server at the current directory]


# Modules included
## Hello World
Initial example to demonstrate how a WASM based website works.

## Bcrypt Hashing and SHA-256 hashing
An example of how to reuse the crypto package from golang in the browser environment.

## String Usecase
An example to show how a GOLANG FUNCTION can be used in JS and vice-versa.
