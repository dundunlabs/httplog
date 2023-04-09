# httplog
Golang net/http logging

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