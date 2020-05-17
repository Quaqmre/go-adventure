package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	ctx := context.Background()
	behavior(ctx, 1)
	behavior(ctx, 1)
	behavior(ctx, 0)
	// den()()
	// den()()
	// den()()
}

func behavior(ctx context.Context, i int) {
	defer func() {
		val := ctx.Value("err")
		if val != nil {
			fmt.Printf("bir error yakalandı:%v", val)
		}
	}()

	if i == 0 {
		err := errors.New("bu bir errordür")
		ctx = context.WithValue(ctx, "err", err)
	}
}

func den() func() {
	i := 0
	return func() {
		i = i + 1
		fmt.Println(i)
	}
}
