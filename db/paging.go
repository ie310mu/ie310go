package db

import (
	"strconv"

	"github.com/ie310mu/ie310go/forks/github.com/ilibs/gosql"
	"github.com/ie310mu/ie310go/forks/github.com/jmoiron/sqlx"
)

//PageIndexToLimit 将页码转换为limit的起始序号
func PageIndexToLimit(pageIndex int, rowsInPage int) (offset int) {
	if pageIndex < 1 {
		pageIndex = 1
	}
	offset = (pageIndex - 1) * rowsInPage
	return
}

//Select 加入分页信息后进行查询
//rowsInPage等于0时不附加limit，可以自行在sql中附加
func Select(DB *gosql.Wrapper, tx *sqlx.Tx, pageIndex int, rowsInPage int, dest interface{}, sqlStr string, args ...interface{}) error {
	if rowsInPage != 0 {
		offset := PageIndexToLimit(pageIndex, rowsInPage)
		sqlStr = sqlStr + " limit " + strconv.Itoa(offset) + "," + strconv.Itoa(rowsInPage)
	}

	if tx == nil {
		if DB == nil {
			return gosql.Select(dest, sqlStr, args...)
		}
		return DB.Select(dest, sqlStr, args...)
	}
	return gosql.WithTx(tx).Select(dest, sqlStr, args...)
}
