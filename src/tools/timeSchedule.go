package tools

import "time"

// scheduleWeekly 启动一个周期性任务，每周特定时间执行
func ScheduleWeekly(targetWeekday time.Weekday, targetHour, targetMinute int, task func()) {
	go func() {
		for {
			next := nextTriggerTime(targetWeekday, targetHour, targetMinute)
			timer := time.NewTimer(time.Until(next))
			<-timer.C
			task()
		}
	}()
}

// nextTriggerTime 计算下一个触发时间点
func nextTriggerTime(targetWeekday time.Weekday, targetHour, targetMinute int) time.Time {
	now := time.Now()

	// 构造当前日期的目标时间（时分秒归零）
	t := time.Date(now.Year(), now.Month(), now.Day(), targetHour, targetMinute, 0, 0, now.Location())

	// 计算到目标星期所需的天数
	daysToAdd := (int(targetWeekday) - int(t.Weekday()) + 7) % 7
	if daysToAdd == 0 {
		// 今天已是目标星期，检查是否已过目标时间
		if now.After(t) {
			daysToAdd = 7 // 已过时间，延至下周
		}
	}
	t = t.AddDate(0, 0, daysToAdd)

	// 处理跨月/年后时间仍早于当前时间的情况（罕见情况）
	if t.Before(now) {
		t = t.AddDate(0, 0, 7)
	}

	return t
}

// ScheduleDaily 启动每日定时任务
func ScheduleDaily(targetHour, targetMinute, targetSecond int, task func()) {
	go func() {
		for {
			next := nextDailyTriggerTime(targetHour, targetMinute, targetSecond)
			timer := time.NewTimer(time.Until(next))
			<-timer.C
			task()
		}
	}()
}

// nextDailyTriggerTime 计算下个触发时间点（精确到秒）
func nextDailyTriggerTime(h, m, s int) time.Time {
	now := time.Now()

	// 构造今日目标时间
	today := time.Date(now.Year(), now.Month(), now.Day(), h, m, s, 0, now.Location())

	// 若当前时间已过今日目标时间，则设置为明日
	if now.After(today) {
		return today.Add(24 * time.Hour)
	}
	return today
}
