/*
created by ie310 at 2019.02.01

mapper的基础类，基于gosql封装而成，gosql github：https://github.com/ilibs/gosql
而gosql是对sqlx的封装：https://blog.csdn.net/wdy_yx/article/details/78262528

mapper的设计类似于java mybatis中的mapper
不同的是，mybatis中sql语句有两种地方可写：注解和xml，这里的go mapper则不一样，sql语句更多的用参数直接传入
配合代码自动生成工具，各个数据表的XXXMapper可自动生成所有的增、删、改、查操作，且可自定义其他操作

*/

package db

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/ie310mu/ie310go/forks/database/sql"

	"github.com/ie310mu/ie310go/common/logsagent"
	"github.com/ie310mu/ie310go/common/obj"
	"github.com/ie310mu/ie310go/common/throw"
	_ "github.com/ie310mu/ie310go/forks/github.com/go-sql-driver/mysql" //导入mysql驱动
	"github.com/ie310mu/ie310go/forks/github.com/ilibs/gosql"
	"github.com/ie310mu/ie310go/forks/github.com/jmoiron/sqlx"
)

//BaseMapper ..
type BaseMapper struct {
	TableName string
	DB        *gosql.Wrapper
}

//Tx 开启一个事务，执行失败后回滚并抛出异常
/*
mp.Tx(func(tx *sqlx.Tx) (r error) {
		defer func() {
			if err := recover(); err != nil {
				msg := obj.InterfaceToStr(err)
				r = errors.New(msg)
			}
		}()

		//do.....................

		return nil
	})
*/
func (m BaseMapper) Tx(fn func(tx *sqlx.Tx) error) {
	logsagent.Info("begin tx")
	err := m.DB.Tx(fn)
	if err == nil {
		logsagent.Info("end tx without err")
	} else {
		logsagent.Info("end tx with err")
		logsagent.Error("%s", err)
	}
	throw.CheckErr(err)
}

type fnDeferError struct {
	fnDonotDeferError func(tx *sqlx.Tx)
}

func (fd fnDeferError) fn(tx *sqlx.Tx) (r error) {
	defer func() {
		if err := recover(); err != nil {
			msg := obj.InterfaceToStr(err)
			r = errors.New(msg)
		}
	}()

	fd.fnDonotDeferError(tx)

	return nil
}

//Tx2 开启一个事务，执行失败后回滚并抛出异常
//2019.06.28简化Tx的封装，之前的实现函数还需要defer err，现在不需要了
/*
mp.Tx2(func(tx *sqlx.Tx)  {
		//do.....................
	})
*/
func (m BaseMapper) Tx2(fnDonotDeferError func(tx *sqlx.Tx)) {
	logsagent.Info("begin tx")

	fd := &fnDeferError{fnDonotDeferError: fnDonotDeferError}
	err := m.DB.Tx(fd.fn)
	if err == nil {
		logsagent.Info("end tx without err")
	} else {
		logsagent.Info("end tx with err")
		logsagent.Error("%s", err)
	}

	throw.CheckErr(err)
}

//ScanItem 获取一行数据，没有数据返回nil，出错抛出异常
func (m BaseMapper) ScanItem(tx *sqlx.Tx, dest interface{}, sqlStr string, args ...interface{}) (r interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if isNoRowsError(err) {
				r = nil
			} else {
				panic(err)
			}
		}
	}()

	var err error
	if tx == nil {
		err = m.DB.Get(dest, sqlStr, args...)
	} else {
		err = gosql.WithTx(tx).Get(dest, sqlStr, args...)
	}
	throw.CheckErr(err)

	logsagent.Info("return " + reflect.ValueOf(dest).String())
	return dest
}

//ScanItems 获取多行数据，没有数据不会在数组中增加元素，出错抛出异常
//rowsInPage都等于0时代表不处理分页，可自行在sql中进行处理
func (m BaseMapper) ScanItems(tx *sqlx.Tx, pageIndex int, rowsInPage int, dest interface{}, sqlStr string, args ...interface{}) {
	defer func() {
		if err := recover(); err != nil {
			if isNoRowsError(err) {
				logsagent.Info("return 0 rows")
			} else {
				panic(err)
			}
		}
	}()

	err := Select(m.DB, tx, pageIndex, rowsInPage, dest, sqlStr, args...) //Select方法在没有数据时不会有err?
	throw.CheckErr(err)
	logsagent.Info("return n rows") //不知道怎样获取dest的实际长度?
	return
}

//ScanInt 获取整数，没有返回0，出错抛出异常
func (m BaseMapper) ScanInt(tx *sqlx.Tx, sqlStr string, args ...interface{}) (r int) {
	defer func() {
		if err := recover(); err != nil {
			if isNoRowsError(err) {
				logsagent.Info("return 0")
				r = 0
			} else {
				panic(err)
			}
		}
	}()

	var row *sqlx.Row
	if tx == nil {
		row = m.DB.QueryRowx(sqlStr, args...)
	} else {
		row = gosql.WithTx(tx).QueryRowx(sqlStr, args...)
	}
	var count int
	err := row.Scan(&count)
	throw.CheckErr(err)
	logsagent.Info("return " + strconv.Itoa(count))
	return count
}

//ScanFloat 获取数值，没有返回0，出错抛出异常
func (m BaseMapper) ScanFloat(tx *sqlx.Tx, sqlStr string, args ...interface{}) (r float64) {
	defer func() {
		if err := recover(); err != nil {
			if isNoRowsError(err) {
				logsagent.Info("return 0")
				r = 0
			} else {
				panic(err)
			}
		}
	}()

	var row *sqlx.Row
	if tx == nil {
		row = m.DB.QueryRowx(sqlStr, args...)
	} else {
		row = gosql.WithTx(tx).QueryRowx(sqlStr, args...)
	}
	var v float64
	err := row.Scan(&v)
	throw.CheckErr(err)
	logsagent.Info("return " + strconv.FormatFloat(v, 'f', 2, 64))
	return v
}

//ScanString 获取字符串值，没有返回""，出错抛出异常
func (m BaseMapper) ScanString(tx *sqlx.Tx, sqlStr string, args ...interface{}) (r string) {
	defer func() {
		if err := recover(); err != nil {
			if isNoRowsError(err) {
				r = ""
			} else {
				panic(err)
			}
		}
	}()

	str := ""
	var row *sqlx.Row
	if tx == nil {
		row = m.DB.QueryRowx(sqlStr, args...)
	} else {
		row = gosql.WithTx(tx).QueryRowx(sqlStr, args...)
	}

	err := row.Scan(&str)
	throw.CheckErr(err)

	return str
}

//DeleteOrUpdateItems 删除或更新数据，出错抛出异常
func (m BaseMapper) DeleteOrUpdateItems(tx *sqlx.Tx, sqlStr string, args ...interface{}) int {
	var sqlResult sql.Result
	var err error

	if tx == nil {
		sqlResult, err = m.DB.Exec(sqlStr, args...)
	} else {
		sqlResult, err = gosql.WithTx(tx).Exec(sqlStr, args...)
	}
	throw.CheckErr(err)
	rows, err := sqlResult.RowsAffected()
	throw.CheckErr(err)

	logsagent.Info("update " + strconv.Itoa(int(rows)) + " rows")
	return int(rows)
}

//InsertItem 插入一行数据，出错抛出异常  如果id是整形，可以返回id(第1个参数，第2个参数是插入的行数)
func (m BaseMapper) InsertItem(tx *sqlx.Tx, sqlStr string, args ...interface{}) (int, int) {
	var sqlResult sql.Result
	var err error

	if tx == nil {
		sqlResult, err = m.DB.Exec(sqlStr, args...)
	} else {
		sqlResult, err = gosql.WithTx(tx).Exec(sqlStr, args...)
	}
	throw.CheckErr(err)

	rows, err := sqlResult.RowsAffected()
	throw.CheckErr(err)
	logsagent.Info("update " + strconv.Itoa(int(rows)) + " rows")

	lastID, err := sqlResult.LastInsertId()
	if err == nil {
		return int(lastID), int(rows)
	}
	return 0, int(rows)
}

//判断错误是否是"没有结果"
func isNoRowsError(err interface{}) bool {
	if strings.Index(fmt.Sprintf("%s", err), "no rows in result set") < 0 {
		return false
	}
	return true
}

//Test 测试数据库是否连接（测试方式，指定sql语句返回一个字符串，如果测试失败，函数会返回失败信息，否则返回""）
func (m BaseMapper) Test(selectStr string, args ...interface{}) (r error) {
	defer func() {
		if err := recover(); err != nil {
			if !isNoRowsError(err) {
				r = fmt.Errorf("%s", err)
			}
		}
	}()

	row := m.DB.QueryRowx(selectStr, args...)
	var result string
	err := row.Scan(&result)
	throw.CheckErr(err)

	return nil
}

//CheckOrderStr ..
func CheckOrderStr(fieldsMap map[string]int, orderStr string) {
	if orderStr == "" {
		return
	}

	ks := strings.FieldsFunc(orderStr, func(r rune) bool { return r == ' ' })
	for _, k := range ks {
		k = strings.Trim(k, " ")
		k = strings.ToLower(k)
		if k == "asc" || k == "desc" {
			continue
		}
		if _, ok := fieldsMap[k]; !ok {
			panic("error part in orderStr : " + k)
		}
	}
}
