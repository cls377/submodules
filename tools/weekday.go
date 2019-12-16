package tools

import (
	"time"
)

//获取星期 格式：“星期？”
func Weektoday() string {
	var week string
	switch time.Now().Weekday().String() {
	case "Monday":
		week = "星期一"
		break

	case "Tuesday":
		week = "星期二"
		break

	case "Wednesday":
		week = "星期三"
		break

	case "Thursday":
		week = "星期四"
		break

	case "Friday":
		week = "星期五"
		break

	case "Saturday":
		week = "星期六"
		break
	case "Sunday":
		week = "星期日"
		break
	}
	return week
}

//时辰
func Dayhour() string {
	var dayhour string
	daytime := time.Now().Hour()
	switch true {
	case (3 <= daytime && daytime < 6):
		dayhour = "凌晨好"
		break

	case (6 <= daytime && daytime < 8):
		dayhour = "早晨好"
		break
	case (8 <= daytime && daytime < 11):
		dayhour = "上午好"
		break
	case (11 <= daytime && daytime < 13):
		dayhour = "中午好"
		break
	case (13 <= daytime && daytime < 17):
		dayhour = "下午好"
		break
	case (17 <= daytime && daytime < 19):
		dayhour = "傍晚好"
		break
	case (19 <= daytime && daytime < 23):
		dayhour = "晚上好"
		break
	case (23 <= daytime && daytime < 3):
		dayhour = "深夜好"
		break

	}
	return dayhour
}
