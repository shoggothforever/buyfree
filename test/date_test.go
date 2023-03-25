package test

import (
	"buyfree/utils"
	"testing"
)

func TestDatefunc(t *testing.T) {
	//y, m, d := utils.GetFirstDayOfWeek()
	//t.Error(utils.GetTimeKeyByMode("dsm", 0))
	//t.Log(utils.GetTimeKeyByMode("dsm", 1))
	//t.Log(utils.GetTimeKeyByMode("dsm", 2))
	//t.Log(utils.GetTimeKeyByMode("dsm", 3))
	//t.Log(utils.GetTimeKeyByMode("dsm", 4))
	//t.Log(utils.GetTimeKeyByMode("dsm", 5))
	t.Error(utils.GetAllKeys("dsm"))
	//if y != 2023 || m != int(time.Now().Month()) || d != 20 {
	//	t.Error("函数实现出错1", y, m, d)
	//}
	//y, m, d = utils.GetFirstDayOfMonth()
	//if y != 2023 || m != int(time.Now().Month()) || d != 1 {
	//	t.Error("函数实现出错2", y, m, d)
	//}
	//y, m, d = utils.GetFirstDayOfYear()
	//if y != 2023 || m != 1 || d != 1 {
	//	t.Error("函数实现出错3", y, m, d)
	//}
	//s := utils.GetBeginningOfTheDay()
	//if s != "2023-3-22:00:00:00:" {
	//	t.Error("函数实现出错4", s)
	//}
}
