package main

import (
	"context"
	"fmt"
)

func main() {
	ctx:=context.WithValue(context.Background(),"hello","world")
	ctx=context.WithValue(ctx,123,456)
	process(ctx)
}

func process(ctx context.Context) {
	result:=ctx.Value("hello").(string)
	fmt.Println(result)
	result1,ok:= ctx.Value(123).(int)
	if !ok{
		fmt.Println("youwu")
		//return
	}
	fmt.Println(result1)
}
