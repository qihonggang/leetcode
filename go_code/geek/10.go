package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	var a, b, c []int

	// 调用广告
	g.Go(func() error {
		a = []int{0}
		return nil
	})

	// 调用AI
	g.Go(func() error {
		b = []int{1}
		return errors.New("ai")
	})

	// 调用运营平台
	g.Go(func() error {
		c = []int{2}
		return nil
	})

	err := g.Wait()
	fmt.Println(err)

	fmt.Println(ctx.Err())
}
