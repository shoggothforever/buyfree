package utils

import (
	"fmt"
	"time"
)

//获得一天开头的确切时间
func GetBeginningOfTheDay(offset int) string {
	y, m, d := time.Now().In(time.Local).AddDate(0, 0, offset).Date()
	return fmt.Sprintf("%d-%d-%d 00:00:00", y, m, d)
}

func GetDailySalesKey(uname string, offset int) string {
	y, timem, d := time.Now().In(time.Local).AddDate(0, 0, offset).Date()
	m := int(timem)
	return fmt.Sprintf("%s:%s:%d-%d-%d", uname, DailySalesKey, y, m, d)
}

func GetAllTimeKeys(uname string) []string {
	keys := []string{
		GetBeginningOfTheDay(0),
		GetBeginningOfTheDay(-1),
		GetTimeKeyByMode(uname, 0),
		GetTimeKeyByMode(uname, 1),
		GetTimeKeyByMode(uname, 2),
		GetTimeKeyByMode(uname, 3),
		GetTimeKeyByMode(uname, 4),
		GetTimeKeyByMode(uname, 5),
		GetDailySalesKey(uname, 0),
		GetDailySalesKey(uname, -1),
		GetDailySalesKey(uname, -2),
		GetDailySalesKey(uname, -3),
		GetDailySalesKey(uname, -4),
		GetDailySalesKey(uname, -5),
		GetDailySalesKey(uname, -6),
	}
	return keys
}

//根据模式获取相应的时间 0一天的开始，1：一周的开始，2：当月第一天，3：当年第一天.4:连续七天,5:总榜
func GetTimeKeyByMode(uname string, mode int) string {
	now := time.Now().In(time.Local)
	y, timem, d := now.Date()
	m := int(timem)
	if mode == 0 {
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, DailySalesKey, y, m, d)
	} else if mode == 1 {
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6
		}
		y, m, d := time.Now().In(time.Local).AddDate(0, 0, offset).Date()
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, WeeklySalesKey, y, m, d)
	} else if mode == 2 {
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, MonthSalesKey, y, m, 1)
	} else if mode == 3 {
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, AnnuallySalesKey, y, 1, 1)
	} else if mode == 4 {
		return fmt.Sprintf("%s:%s", uname, Constantly7aysSalesKey)
	} else if mode == 5 {
		return fmt.Sprintf("%s:%s", uname, TOTALSalesKey)
	}
	return "root:root:2006-1-2 15:-4:-5"
}

//根据模式获取相应的时间 0一天的开始，1：周排行榜键名，2：月排行榜键名，3：年排行榜键名.4:总榜
func GetRankKeyByMode(adp, uname string, mode int) string {
	now := time.Now().In(time.Local)
	y, timem, d := now.Date()
	m := int(timem)
	if mode == 0 {
		//uname:adp..DailyRanksKey:y-m-d
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, adp+DailyRanksKey, y, m, d)
	} else if mode == 1 {
		offset := int(time.Monday - now.Weekday())
		if offset > 0 {
			offset = -6
		}
		y, m, d := time.Now().In(time.Local).AddDate(0, 0, offset).Date()
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, adp+WeeklyRanksKey, y, m, d)
	} else if mode == 2 {
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, adp+MonthRanksKey, y, m, 1)
	} else if mode == 3 {
		return fmt.Sprintf("%s:%s:%d-%d-%d", uname, adp+AnnuallyRanksKey, y, 1, 1)
	} else if mode == 4 {
		return fmt.Sprintf("%s:%s", uname, adp+TOTALRanksKey)
	}
	return "root:root:2006-1-2 15:-4:-5"
}

func GetAllTypeRankKeys(adp, uname string) []string {
	s := []string{}
	for i := 0; i <= 4; i++ {
		s = append(s, GetRankKeyByMode(adp, uname, i))
		s = append(s, GetRankKeyByMode("All"+adp, uname, i))
	}
	fmt.Println(s)
	return s
}

//func GetAllADRankKeys(adid string) []string {
//	s := []string{
//		GetRankKeyByMode(adid, 0),
//		GetRankKeyByMode("AD", 0),
//		GetRankKeyByMode(adid, 1),
//		GetRankKeyByMode("AD", 1),
//		GetRankKeyByMode(adid, 2),
//		GetRankKeyByMode("AD", 2),
//		GetRankKeyByMode(adid, 3),
//		GetRankKeyByMode("AD", 3),
//		GetRankKeyByMode(adid, 4),
//		GetRankKeyByMode("AD", 4),
//	}
//	fmt.Println(s)
//	return s
//}
