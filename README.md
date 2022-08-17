# go webassembly experiments

Requires go1.13

- [bouncy](https://stdiopt.github.io/gowasm-experiments/bouncy)
- [rainbow-mouse](https://stdiopt.github.io/gowasm-experiments/rainbow-mouse)
- [repulsion](https://stdiopt.github.io/gowasm-experiments/repulsion)
- [bumpy](https://stdiopt.github.io/gowasm-experiments/bumpy)
- [splashy](https://stdiopt.github.io/gowasm-experiments/splashy)
- [arty](https://stdiopt.github.io/gowasm-experiments/arty/client)
- [hexy](https://stdiopt.github.io/gowasm-experiments/hexy)
- [bittune](https://stdiopt.github.io/gowasm-experiments/bittune)

## Building and running

```sh
$ cd {proj} # sub folder (i.e. bouncy, rainbow-mouse)
$ go get -v # ignore the js warning
$ ./build.sh
$ go run ../serve.go
```

Serve with caddy or anything else that is able to set the mimetype
'application/wasm' for .wasm files

## geek

```sh
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ./static

GOOS=js GOARCH=wasm go build -o geek/main.wasm ./geek
GOOS=js GOARCH=wasm go build -o geek2/main.wasm ./geek2
GOOS=js GOARCH=wasm go build -o geek3/main.wasm ./geek3
GOOS=js GOARCH=wasm go build -o geek4/main.wasm ./geek4

go run serve.go
```

## hello

```sh
export PATH="$PATH:$(go env GOROOT)/misc/wasm"
GOOS=js GOARCH=wasm go run ./hello

# or
GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" ./hello

# or
GOOS=js GOARCH=wasm go build -o ./hello/main.wasm ./hello
$(go env GOROOT)/misc/wasm/go_js_wasm_exec ./hello/main.wasm
```

## emcc

```sh
git clone https://github.com/juj/emsdk.git
cd emsdk
./emsdk install latest
./emsdk activate latest
source "./emsdk_env.sh"
echo './emsdk_env.sh"' >> $HOME/.bash_profile

emcc hello.c -s WASM=1 -o hello.html
```
