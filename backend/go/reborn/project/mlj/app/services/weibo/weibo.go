package weibo

import (
	"fmt"
	"mlj/pkg/database"
)

type Weibo struct {
	Request  WeiboParams
	Response WeiboResponse
}

type WeiboParams struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	Uid       string `form:"uid"`
}

type WeiboResponse struct {
	WeiboCount    int `json:"weibo_count"`
	TbWeiboCount  int `json:"tb_weibo_count"`
	PddWeiboCount int `json:"pdd_weibo_count"`
	JdWeiboCount  int `json:"jd_weibo_count"`
}

func (w *Weibo) Query() (err error) {
	sql := fmt.Sprintf(`select
		count(distinct(item_id)) as weibo_count,
		sum(if(platform in (1,2,3,5), 1, 0)) as 'tb_weibo_count',
		sum(if(platform=6, 1, 0)) as 'pdd_weibo_count',
		sum(if(platform=4, 1, 0)) as 'jd_weibo_count'
	from
		weibo_product_info
	where
	    created_time >= '%s 00:00:00'
		and created_time <= '%s 23:59:59'
		and uid = '%s'`, w.Request.StartTime, w.Request.EndTime, w.Request.Uid)
	fmt.Printf("querysql = %s", sql)
	if err := database.DBWeibo.Raw(sql).Scan(&w.Response).Error; err != nil {
		return err
	}

	fmt.Println(w)

	return nil
}
