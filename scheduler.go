package main

import "time"

/*
Scheduler ...
*/
type Scheduler struct {
	timerMap map[string]uint64
	Timer    time.Timer
}

/*
Register ...
*/
func (s *Scheduler) Register(key string, time uint64) {
	timerMap[key] = time
}
