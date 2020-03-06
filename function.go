package function

import (
    "context"
    "fmt"
)

type PubSubMessage struct {
    Data []byte `json:"data"`
}

func Subscription(ctx context.Context, m PubSubMessage) error {
    if err := RemoToBq(ctx); err != nil {
        fmt.Println(err)
        return err
    }
    return nil
}
