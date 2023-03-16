package main

import (
	"fmt"
	"time"
)

type User struct {
	userName string
	workTime string
	lateTime string
	pens     bool
}

type Worker interface {
	name() string
	time() string
	penns() bool
	lates() string

	addName(string)
	addWorkTime(string)
	addLateTime(string)
	addPens(bool)
	wTime() time.Time
}

//	var curUser = User{
//		userName: "",
//		workTime: "12h15min",
//		lateTime: "0min",
//		pens:     false,
//	}
func (u *User) name() string {
	return u.userName
}
func (u *User) time() string {
	return u.workTime
}
func (u *User) penns() bool {
	return u.pens
}
func (u *User) lates() string {
	return u.lateTime
}
func (u *User) addName(name string) {
	u.userName = name
}
func (u *User) addWorkTime(time string) {
	u.workTime = time
}
func (u *User) addLateTime(lates string) {
	u.lateTime = lates
}
func (u *User) addPens(b bool) {
	u.pens = true
}

func (u *User) wTime() time.Time {
	timeNow := time.Now()
	endTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 21, 15, 00, 0, time.Local)
	diff := endTime.Sub(timeNow)
	out := time.Time{}.Add(diff)
	workTime := fmt.Sprintf(out.Format("15:04"))
	fmt.Println("workTime", workTime)

	startTime := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 9, 15, 00, 0, time.Local)
	diff = timeNow.Sub(startTime)
	out = time.Time{}.Add(diff)
	lateTime := fmt.Sprintf(out.Format("15:04"))
	if diff.Hours() < 0 {
		lateTime = "-"
	}
	fmt.Println("lateTime", lateTime)

	u.addWorkTime(workTime)
	u.addLateTime(lateTime)

	if diff.Minutes() > 15 {
		u.addPens(true)
	}

	return timeNow
}
