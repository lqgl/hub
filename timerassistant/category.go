package timerassistant

import "time"

type Category interface {
	ShouldCall() bool
	SetLastCallTime(time int64)
}

// Once call
type Once struct {
	Hour         int
	Min          int
	Sec          int
	lastCallTime int64
}

func (o *Once) ShouldCall() bool {
	if o.lastCallTime != 0 {
		return false
	}
	now := time.Now()
	hour, min, sec := now.Local().Clock()
	if (hour*3600 + min*60 + sec) >= (o.Hour*3600 + o.Min*60 + o.Sec) {
		return true
	}
	return false
}

// SetLastCallTime set last call time
func (o *Once) SetLastCallTime(lastCallTime int64) {
	o.lastCallTime = lastCallTime
}

// Daily call
type Daily struct {
	Hour         int64
	Min          int64
	Sec          int64
	lastCallTime int64
}

func (d *Daily) ShouldCall() bool {
	now := time.Now()
	hour, min, sec := now.Local().Clock()
	nowSec := int64(hour*3600) + int64(min*60) + int64(sec)

	if d.lastCallTime == 0 {
		if nowSec >= (d.Hour*3600 + d.Min*60 + d.Sec) {
			return true
		}
	} else {
		// elapsed one day
		if nowSec >= (d.Hour*3600+d.Min*60+d.Sec) && (nowSec-d.lastCallTime >= int64(86400)) {
			return true
		}
	}
	return false
}

func (d *Daily) SetLastCallTime(lastCallTime int64) {
	d.lastCallTime = lastCallTime
}

// Weekly call
type Weekly struct {
	Hour         int64
	Min          int64
	Sec          int64
	lastCallTime int64
}

func (w *Weekly) ShouldCall() bool {
	now := time.Now()
	hour, min, sec := now.Local().Clock()
	nowSec := int64(hour*3600) + int64(min*60) + int64(sec)

	if w.lastCallTime == 0 {
		if nowSec >= (w.Hour*3600 + w.Min*60 + w.Sec) {
			return true
		}
	} else {
		// elapsed one week
		if nowSec >= (w.Hour*3600+w.Min*60+w.Sec) && (nowSec-w.lastCallTime >= int64(7*86400)) {
			return true
		}
	}
	return false
}

func (w *Weekly) SetLastCallTime(lastCallTime int64) {
	w.lastCallTime = lastCallTime
}
