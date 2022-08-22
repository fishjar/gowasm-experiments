```sh
wget https://raw.githubusercontent.com/tinygo-org/tinygo/release/targets/wasm_exec.js
tinygo build -o ./main.wasm -target wasm -no-debug ./main.go
npx serve .
```
