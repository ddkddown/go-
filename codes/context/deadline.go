package main

import (
	"context"
	"fmt"
	"time"
)

func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("recv!")
			return
		default:
			fmt.Println("waiting....")
			time.Sleep(1 * time.Second)
		}
	}
}
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))
	go watch(ctx)
	time.Sleep(5 * time.Second)
	fmt.Println("over!")
	defer cancel()
}
