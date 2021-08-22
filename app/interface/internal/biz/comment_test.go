package biz

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTmp1(t *testing.T) {
	do := CommentDO{
		UserId:   1,
		UserName: "xiaom",
		FoodId:   1,
		FoodName: "nihao",
		Comment:  "helel",
	}
	marshal, err := json.Marshal(do)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshal))

}
