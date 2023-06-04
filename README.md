# httplog

[![release](https://img.shields.io/github/v/release/dundunlabs/httplog)](https://github.com/dundunlabs/httplog/releases)
[![license](https://img.shields.io/github/license/dundunlabs/httplog)](https://github.com/dundunlabs/httplog/blob/main/LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://makeapullrequest.com)

HTTP logger for Golang's [net/http](https://pkg.go.dev/net/http)

## Installation

```bash
go get -u github.com/dundunlabs/httplog
```

## Usage

```go
package main

import (
	"net/http"

	"github.com/dundunlabs/httplog"
)

func main() {
	handler := httplog.NewHandler(http.DefaultServeMux)
	http.ListenAndServe(":8080", handler)
}
```

## License

[MIT](https://github.com/dundunlabs/httplog/blob/main/LICENSE)

---

Enjoy your coding 🚀
