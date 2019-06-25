
# for what

- this project used to gin middware
- [ ] replace pkg name `github.com/bar-counter/template` to target

## dependInfo

| lib | url | version |
|:-----|:-----|:-----|
| gin | https://github.com/gin-gonic/gin | 1.4.0 |

# demo

```bash
make init
make dep
# ensure right then
make dev
# and open url
# health http://127.0.0.1:38000/status/health
# pprof http://127.0.0.1:38000/debug/pprof/
```

# use middleware

```bash
# go get
go get -v github.com/bar-counter/template

# dep go 1.7 -> 1.11
dep ensure --add github.com/bar-counter/template@1.0.0
dep ensure -v
```

