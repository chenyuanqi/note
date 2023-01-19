package Business

import (
	"fmt"
	"mlj/pkg/database"
)

type Business struct {
	Request  BusinessParams
	Response BusinessResponse
}

type BusinessParams struct {
	StartTime string `form:"start_time"`
	EndTime   string `form:"end_time"`
	Uid       string `form:"uid"`
}

type BusinessResponse struct {
	Computed    float64 `json:"computed"`
	TbComputed  float64 `json:"tb_computed"`
	PddComputed float64 `json:"pdd_computed"`
	JdComputed  float64 `json:"jd_computed"`
	Count       int     `json:"count"`
	TbCount     int     `json:"tb_count"`
	PddCount    int     `json:"pdd_count"`
	JdCount     int     `json:"jd_count"`
}

func (b *Business) Query() (err error) {
	sql := fmt.Sprintf(`select
		sum(computed) as 'computed',
		sum(if(platform='ali', computed, 0)) as 'tb_computed',
		sum(if(platform='pdd', computed, 0)) as 'pdd_computed',
		sum(if(platform='jd', computed, 0)) as 'jd_computed',
		sum(order_success_count) as 'count',
		sum(if(platform='ali', order_success_count, 0)) as 'tb_count',
		sum(if(platform='pdd', order_success_count, 0)) as 'pdd_count',
		sum(if(platform='jd', order_success_count, 0)) as 'jd_count'
	from
		ho_order_statistics
	where
		date >= '%s'
		and date <= '%s'
		and uid = '%s'`, b.Request.StartTime, b.Request.EndTime, b.Request.Uid)
	fmt.Printf("querysql = %s", sql)
	if err := database.DBItMain.Raw(sql).Scan(&b.Response).Error; err != nil {
		return err
	}

	return nil
}
