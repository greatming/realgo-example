package controllers

import (
	"realgo"
	"realgo/lib/logger"
	"myframe/utils/errors"
	"myframe/utils"
	"fmt"
	"time"
	"strconv"
	"github.com/bitly/go-simplejson"
)

type Response struct {
	Ctx *realgo.WebContext
	Error   errors.AppError
	Result  interface{}
	Header  map[string]string
	StartTime time.Time
}

func NewResponse(ctx *realgo.WebContext) *Response {
	r := &Response{
		Ctx:ctx,
	}
	r.Init()
	return r
}

func (r *Response)Init()  {
	r.Ctx.Logger = logger.New()
	r.Ctx.Logger.SetBaseInfo("ip", r.Ctx.Request.RemoteAddr)
	r.Ctx.Logger.SetBaseInfo("url", r.Ctx.Request.RequestURI)
	r.Error = errors.SUCCESS
	r.StartTime = time.Now()
}

func (r *Response)RequestDone()  {
	r.Ctx.Response.WriteHeader(r.GetStatus())
	for key, val := range r.GetHeader() {
		r.Ctx.Response.Header().Set(key, val)
	}
	r.Ctx.Response.Write(r.GetBody())
	r.Ctx.Logger.PushNotice("exec_time", fmt.Sprintf("%v", time.Since(r.StartTime)))
	r.Ctx.Logger.PushNotice("errno", strconv.Itoa(r.GetErrno()))
}


func (t *Response) GetStatus() int {
	return t.Error.Status
}
func (t *Response) GetBody() []byte {
	res := simplejson.New()
	result, _ := utils.ToJSONObject(t.Error)
	resultMap, _ := result.Map()
	for k, v := range resultMap {
		res.Set(k, v)
	}
	if t.Result != nil {
		record, _ := utils.ToJSONObject(t.Result)
		recordMap, _ := record.Map()
		for k, v := range recordMap {
			res.Set(k, v)
		}
	}
	body, _ := res.MarshalJSON()
	return body
}
func (t *Response) GetHeader() map[string]string {
	if t.Header == nil {
		t.Header = make(map[string]string, 1)
	}
	t.Header["Content-Type"] = "application/json; charset=UTF-8"
	return t.Header
}
func (t *Response) GetErrno() int {
	return t.Error.Errno
}
