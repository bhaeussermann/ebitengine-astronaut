rm -rf bin/web
mkdir bin/web
env GOOS=js GOARCH=wasm go build -o bin/web/astronaut.wasm github.com/bhaeussermann/ebitengine-astronaut/src
cp web/*.* bin/web
