package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "foo", "bar")
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		fmt.Printf("fetch user data error: %v\n", err)
	}
	fmt.Printf("fetch user data: %v\n", val)
	fmt.Printf("time elapsed: %v\n", time.Since(start))
}

type Response struct {
	val int
	err error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	value := ctx.Value("foo")
	fmt.Println(value)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	reschan := make(chan Response)
	go func() {
		value, err := fetchThirdParty()
		reschan <- Response{val: value, err: err}
	}()
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took so long")
		case r := <-reschan:
			return r.val, r.err
		}

	}
}

func fetchThirdParty() (int, error) {
	time.Sleep(time.Millisecond * 150)
	return 66, nil
}
