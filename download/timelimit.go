package igdl

import (
	"time"
)

type TimeLimiter struct {
	LastTime time.Time
	Interval int64 // unit: second
}

func NewTimeLimiter(interval int64) *TimeLimiter {
	tl := TimeLimiter{
		LastTime: time.Unix(time.Now().Unix()-interval, 0),
		Interval: interval,
	}
	return &tl
}

func (tl *TimeLimiter) WaitAtLeastIntervalAfterLastTime() {
	d := time.Now().Sub(tl.LastTime)
	interval := time.Duration(tl.Interval) * time.Second
	if d < interval {
		time.Sleep(interval - d)
	}
}

func (tl *TimeLimiter) WaitAtLeastNIntervalAfterLastTime(n int64) {
	d := time.Now().Sub(tl.LastTime)
	interval := time.Duration(tl.Interval*n) * time.Second
	if d < interval {
		time.Sleep(interval - d)
	}
}

func (tl *TimeLimiter) IsOverNIntervalAfterLastTime(n int64) bool {
	d := time.Now().Sub(tl.LastTime)
	interval := time.Duration(tl.Interval*n) * time.Second
	if d < interval {
		return false
	}
	return true
}

func (tl *TimeLimiter) SetLastTimeToNow() {
	tl.LastTime = time.Now()
}
