package web

import "time"

type LogRequest struct {
	StartTime 	time.Time
	EndTime 	time.Duration
	StatusCode 	int
	ClientIp 	string
	Method 		string
	Path 		string
	UserAgent 	string
}