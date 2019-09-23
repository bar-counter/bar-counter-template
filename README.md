
# for what

- this project used to gin middware
- [ ] replace pkg name `github.com/bar-counter/bar-counter-template` to target

## dependInfo

| lib | url | version |
|:-----|:-----|:-----|
| go sdk | https://golang.org/ | 1.12+ |
| gin | https://github.com/gin-gonic/gin | 1.4.0 |

# demo

```bash
make init
# ensure right then
make dev
# run as docker contant
$ make dockerRun
# stop or remove docker
$ make dockerStop
$ make dockerRemove
```

# import plugin

```bash
# above go 1.12
GOPROXY=https://mirrors.aliyun.com/goproxy/ GO111MODULE=on go mod edit -require=github.com/json-iterator/go@v1.1.7
make dep
```

```go
import "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary
```
