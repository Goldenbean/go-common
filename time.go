package common

import (
	"fmt"
	"time"
)

// GetTimeStamp : get timestamp
func GetTimeStamp() {
	now := time.Now()
	//	secs := now.Unix()
	//	nanos := now.UnixNano()
	fmt.Println("now: ", now)
	//fmt.Println("second: ", secs)
	//fmt.Println("nano: ", nanos)
}
