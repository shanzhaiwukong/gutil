package gutil

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Pager 分页
type Pager struct {
	// Condition 分页条件
	Condition *PagerCondition `json:"condition"`
	// Size 每页大小
	Size int64 `json:"size"`
	// Index 当前页索引
	Index int64 `json:"index"`
	// Total 总页数
	Total int64 `json:"total"`
	// Items 总条数
	Items int64 `json:"items"`
	// List 记录列表
	List interface{} `json:"list"`
	// Extra 扩展字典
	Extra map[string]interface{} `json:"extra"`
}

// PagerCondition 分页条件
type PagerCondition struct {
	// Eqs 等于
	Eqs map[string]interface{} `json:"eqs"`
	// Nes 不等于
	Nes map[string]interface{} `json:"nes"`
	// Gts 大于
	Gts map[string]interface{} `json:"gts"`
	// Ges 大等于
	Ges map[string]interface{} `json:"ges"`
	// Lts 小于
	Lts map[string]interface{} `json:"lts"`
	// Les 小等于
	Les map[string]interface{} `json:"les"`
	// Lks 模糊查询
	Lks map[string]interface{} `json:"lks"`
	// Sorts 查询条件
	Sorts []string `json:"sorts"`
}

// DeserializePager 反序列化pager对象
func DeserializePager(str string) (*Pager, error) {
	p := &Pager{}
	err := json.Unmarshal([]byte(str), p)
	return p, err
}

// NewPager 初始化对象
func NewPager(size, index, total, items int64, list interface{}) *Pager {
	return &Pager{
		Condition: &PagerCondition{},
		Size:      size,
		Index:     index,
		Total:     total,
		Items:     items,
		List:      list,
		Extra:     make(map[string]interface{}),
	}
}

// ComputePage 计算页数
func (that *Pager) ComputePage(items int64) {
	if that.Size == 0 {
		that.Size = 20
	}
	that.Items = items
	that.Total = items / that.Size
	if items%that.Size > 0 {
		that.Total++
	}
}

// ToSQL 转换为查询语句
// return (where 1=1 and xx=xx and yy=yy),sort (order by xx,yy),limit (limit 0,10),params
func (that *Pager) ToSQL() (string, string, string, []interface{}) {
	var where = make([]string, 0)
	where = append(where, " where 1=1")
	var sort string
	var limit string
	var param = make([]interface{}, 0)
	if that == nil || that.Condition == nil {
		return "", "", "", nil
	}
	for k, v := range that.Condition.Eqs {
		where = append(where, k+"=?")
		param = append(param, v)
	}
	for k, v := range that.Condition.Nes {
		where = append(where, k+"!=?")
		param = append(param, v)
	}
	for k, v := range that.Condition.Gts {
		where = append(where, k+">?")
		param = append(param, v)
	}
	for k, v := range that.Condition.Lts {
		where = append(where, k+"<?")
		param = append(param, v)
	}
	for k, v := range that.Condition.Ges {
		where = append(where, k+">=?")
		param = append(param, v)
	}
	for k, v := range that.Condition.Les {
		where = append(where, k+"<=?")
		param = append(param, v)
	}
	for k, v := range that.Condition.Lks {
		where = append(where, k+" like ?")
		param = append(param, "%"+v.(string)+"%")
	}
	sort = strings.Join(that.Condition.Sorts, ",")
	if sort != "" {
		sort = "order by " + sort
	}
	limit = fmt.Sprintf("limit %d,%d", (that.Index-1)*that.Size, that.Size)
	return strings.Join(where, " and "), sort, limit, param
}
