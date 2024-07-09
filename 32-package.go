package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context, id string) (string, error) {

	//deriving context with a parent context
	ctx, cancelFun := context.WithTimeout(ctx, 3*time.Second)
	defer cancelFun()

	resCh := make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		resCh <- "akhlaq"
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resCh:
		return res, nil
	}
}
func datKu(ctx context.Context) {

	ctx = context.WithValue(ctx, "key", "val")
	//goroutine is prone to data race condition
	//so passing values with context is safe for goroutine

	data := ctx.Value("key")
	fmt.Println(data)
}

func main() {
	ctx := context.Background() //creating parent context
	id := "111"

	res, err := fetchData(ctx, id)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}
