# Peace Net SDK for Go
## Installing

Use `go get` to retrieve the SDK to add it to your project's Go module dependencies.
```go
go get github.com/naruse666/peace-net-sdk-go/guardian
```

To update the SDK.
```go
go get -u github.com/naruse666/peace-net-sdk-go/guardian
```

## Quick Examples

```go
package main

import (
	"fmt"
	"github.com/naruse666/peace-net-sdk-go/guardian"
	"os"
)

func main() {
	apiKey, ok := os.LookupEnv("API_KEY")
	if !ok {
		fmt.Println("API_KEY is not set")
		return
	}

	guardianInput := guardian.GuardianInput{
		Text:   "最低な文章！",
		APIKey: apiKey,
	}

	resp, err := guardian.RequestGuardian(guardianInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(resp)
}
```

## Getting Help
Please submit issue to report a Bug or Feature requests.  
[Submit Issue](https://github.com/naruse666/peace-net-sdk-go/issues)
