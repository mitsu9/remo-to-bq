package main

import (
    "context"
	"fmt"

    "github.com/mitsu9/remo-to-bq"
)

func main() {
    ctx := context.Background()
    if err := function.RemoToBq(ctx); err != nil {
        fmt.Println(err)
        return
    }
    return
}
