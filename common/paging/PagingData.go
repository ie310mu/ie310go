package paging

//PagingData ..
type PagingData struct {

	//第几页  页数从1开始
	PageIndex int `json:"pageIndex"    `

	//每页行数
	RowsInPage int `json:"rowsInPage"    `

	//总页数
	Pagecount int `json:"pagecount"    `

	//总记录条数
	Total int `json:"total"    `

	//当前页的数据集合
	Rows interface{} `json:"rows"    `

	//当前页的数据集合
	UserData interface{} `json:"userData"    `
}
