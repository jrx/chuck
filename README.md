# chuck

## compile
```bash
export GOPATH=/Users/jan/go
go install github.com/jrx/chuck/
```

## run
```bash
$GOPATH/bin/chuck
```

## build
```bash
env GOOS=linux GOARCH=amd64 go build -v github.com/jrx/chuck/
```

## deploy to DC/OS

```bash
dcos marathon app add chuck.json
```
