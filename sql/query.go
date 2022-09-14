package sql

import (
	"github.com/TurboHsu/Posta/util"
)

func QueryArticleByID(id int) string {
	return ""
}

func RunQuery(q string) (ret []TestQuery) {
	rows, err := DB.Query(q)
	util.HandleErr(err)
	defer rows.Close()

	for rows.Next() {
		var id int
		var data string
		err = rows.Scan(&id, &data)
		util.HandleErr(err)
		ret = append(ret, TestQuery{ID: id, Data: data})
	}

	return
}

type TestQuery struct {
	ID   int    `json:"id"`
	Data string `json:"data"`
}
