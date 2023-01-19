package bi

import (
	"mlj/app/services/business"
	"mlj/app/services/weibo"
	"sync"
)

type Bi struct {
	Request  BiParams
	Response BiResponse

	bs business.Business
	w  weibo.Weibo
}

type BiParams struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	Uid       string `form:"uid"`
}

type BiResponse struct {
	business.BusinessResponse
	weibo.WeiboResponse
}

func (b *Bi) Query() (err error) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		b.bs.Request.Uid = b.Request.Uid
		b.bs.Request.StartTime = b.Request.StartTime
		b.bs.Request.EndTime = b.Request.EndTime
		b.bs.Query()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		b.w.Request.Uid = b.Request.Uid
		b.w.Request.StartTime = b.Request.StartTime
		b.w.Request.EndTime = b.Request.EndTime
		b.w.Query()
	}()

	wg.Wait()

	// assembly
	b.Response.BusinessResponse = b.bs.Response
	b.Response.WeiboResponse = b.w.Response

	return nil
}
