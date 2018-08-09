package btime

import (
	"btime"
	"fmt"
	// "github.com/imroc/biu"
	"reflect"
	"testing"
)

func Test_GetBinary(t *testing.T) {
	startTime := "08:00:00"
	endTime := "09:00:00"
	ret, err := btime.GetBinary(startTime, endTime)
	if nil != err {
		t.Error("测试没通过")
	} else {
		t.Log("测试通过")
		fmt.Println("type:", reflect.TypeOf(ret))
		// for _, v := range ret {
		// fmt.Println(biu.ToBinaryString(v))
		// fmt.Println(v)
		// }
	}
}

func Test_BinaryOccupy(t *testing.T) {
	startTime := "09:00:00"
	endTime := "10:00:00"
	current := [5]uint64{0, 281406257233920, 0, 0, 0}
	ret, err := btime.GetBinary(startTime, endTime)
	if nil != err {
		t.Error("测试没通过")
	}
	occupyRet, err := btime.BinaryOccupy(ret, current)
	if nil != err {
		t.Error("测试没通过")
	} else {
		t.Log("测试通过")
		fmt.Println("type:", reflect.TypeOf(occupyRet))
	}
}

func Test_BinaryOccupy_conflict(t *testing.T) {
	startTime := "08:55:00"
	endTime := "10:00:00"
	current := [5]uint64{0, 281406257233920, 0, 0, 0}
	ret, err := btime.GetBinary(startTime, endTime)
	if nil != err {
		t.Error("测试没通过")
	}
	occupyRet, err := btime.BinaryOccupy(ret, current)
	if nil != err {
		t.Log("测试通过,冲突正常")
	} else {
		t.Error("测试没通过,冲突占用不正常")
		fmt.Println("type:", reflect.TypeOf(occupyRet))
	}
}
