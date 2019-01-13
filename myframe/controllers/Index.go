package controllers

import (
	"realgo"
	"myframe/model"
	"fmt"
	"myframe/utils/errors"
)

func Index(ctx *realgo.WebContext)  {
	res := NewResponse(ctx)
	defer res.RequestDone()

	db, err := model.GetBackUpDBHandler().GetInstance()

	if err != nil{
		fmt.Println(err)
		res.Error = errors.DB_ERROR
		return
	}

	sql := "select name from IK_Config limit 1;"
	rows,_ := db.Query(sql)
	for rows.Next(){
		var name string
		rows.Scan(&name)
		fmt.Println(name)
	}



	ret := make(map[string]string)

	ret["name"] = "haoming"
	res.Result = ret


}
