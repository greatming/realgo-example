package middleware

import (
	"realgo"
)

func PrintLog(ctx *realgo.WebContext)  {
	if ctx.Logger != nil{
		ctx.Logger.Notice("request done")
	}
}

