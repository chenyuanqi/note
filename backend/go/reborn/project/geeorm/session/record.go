package session

import (
	"geeorm/clause"
	"reflect"
)

func (s *Session) Insert(values ...interface{}) (int64, error) {
	recordValues := make([]interface{}, 0)
	// 多次调用 clause.Set() 构造好每一个子句
	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}

	s.clause.Set(clause.VALUES, recordValues...)
	// 调用一次 clause.Build() 按照传入的顺序构造出最终的 SQL 语句
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// destSlice.Type().Elem() 获取切片的单个元素的类型 destType
	destType := destSlice.Type().Elem()
	// 使用 reflect.New() 方法创建一个 destType 的实例，作为 Model() 的入参，映射出表结构 RefTable()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	// 根据表结构，使用 clause 构造出 SELECT 语句，查询到所有符合条件的记录 rows
	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	// 遍历每一行记录，利用反射创建 destType 的实例 dest，将 dest 的所有字段平铺开，构造切片 values
	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		// 调用 rows.Scan() 将该行记录每一列的值依次赋值给 values 中的每一个字段
		if err := rows.Scan(values...); err != nil {
			return err
		}
		// 将 dest 添加到切片 destSlice 中
		// 循环直到所有的记录都添加到切片 destSlice 中
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()
}
