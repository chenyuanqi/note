package bussiness

import (
	"fmt"
	"mlj/pkg/database"
)

type Bussiness struct {
	Request  BussinessParams
	Response interface{}
}

type BussinessParams struct {
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

func (b Bussiness) Query() (err error) {
	res := make([]*BusinessResponse, 1)
	sql := fmt.Sprintf(`select
	sum(computed) as 'Computed',
	sum(if(platform='ali', computed, 0)) as 'TbComputed',
	sum(if(platform='pdd', computed, 0)) as 'PddComputed',
	sum(if(platform='jd', computed, 0)) as 'JdComputed',
	sum(order_success_count) as 'Count',
	sum(if(platform='ali', order_success_count, 0)) as 'TbCount',
	sum(if(platform='pdd', order_success_count, 0)) as 'PddCount',
	sum(if(platform='jd', order_success_count, 0)) as 'JdCount'
  from
	ho_order_statistics
  where
	date >= '%s'
	and date <= '%s'
	and uid = '%s'`, b.Request.StartTime, b.Request.EndTime, b.Request.Uid)
	fmt.Printf("querysql = %s", sql)
	if err := database.DBItMain.Raw(sql).Scan(&res).Error; err != nil {
		return err
	}

	b.Response = res[0]
	// fmt.Println(res[0])

	return nil
}
