package btime

import (
	"btime"
	"fmt"
	// "github.com/imroc/biu"
	// "reflect"
	"testing"
)

func BenchmarkGetBinary(b *testing.B) {
	startTime := "08:00:00"
	endTime := "09:00:00"
	for i := 0; i < b.N; i++ {
		btime.GetBinary(startTime, endTime)
	}
}

func BenchmarkBinaryOccupy(b *testing.B) {
	startTime := "09:00:00"
	endTime := "10:00:00"
	for i := 0; i < b.N; i++ {
		current := [5]uint64{0, 281406257233920, 0, 0, 0}
		ret, err := btime.GetBinary(startTime, endTime)
		if nil != err {
			fmt.Println("有错了")
		}
		btime.BinaryOccupy(ret, current)
	}
}
