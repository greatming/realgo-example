package middleware

import (
	"realgo"
	"fmt"
)

func Test(ctx *realgo.WebContext)  {
	fmt.Println("this is a mid")
	ctx.Set("name", "haoming")
}
