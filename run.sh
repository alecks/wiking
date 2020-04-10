command -v go >/dev/null 2>&1 || { echo >&2 "'go' is required but it either isn't installed or isn't in PATH. Abort."; exit 1; }
command -v ng >/dev/null 2>&1 || { echo >&2 "'ng' is required but it either isn't installed or isn't in PATH. Abort."; exit 1; }

cd ./angular/
ng build --prod=true

cd ../api/
go get -v ./...
go run .