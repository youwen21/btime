/*
 这是一个把时间按5分钟为一份的，基于二进制的占用时间的包
*/
package btime

import (
	"errors"
	"math"
	"time"
)

type myarr [2]int64

func TryOccupy(startTime string, endTime string, currentOccupy [5]uint64) (init [5]uint64, err error) {
	willUse, err := GetBinary(startTime, endTime)
	if err != nil {
		return init, err
	}
	ret, err := BinaryOccupy(willUse, currentOccupy)
	return ret, err
}

/**
 * [main description]
 *
 * timeByte说明：
 * timeBype[0] = 0-5时，  <60*60*5  每5分钟一份，也就是小于60份
 * timeBype[1] = 5-10时， <60*60*10 每5分钟一份，也就是小于120份
 * timeBype[2] = 10-15时，<60*60*15 每5分钟一份，也就是小于180份
 * timeBype[3] = 15-20时，<60*60*20 每5分钟一份，也就是小于240份
 * timeBype[4] = 20-24时，<60*60*24 每5分钟一份，也就是小于300份
 * @return {[type]} [description]
 *
 * 当前问题 ，以5分钟取整 ，传入的时间必须是5分钟的整数倍
 * 如 9:00到9:16，当前是按9:00到9:15算
 * 如有需要，可以加入math函数，向上取整，可解决此问题
 *
 */
func GetBinary(startTime string, endTime string) (init [5]uint64, err error) {

	if endTime == "24:00:00" {
		endTime = "23:59:00"
	}

	// 当天开始时间
	today := time.Now().Format("2006-01-02")
	todayTimeBegin, err := time.Parse("2006-01-02", today)
	if err != nil {
		return init, errors.New("获取期失败")
	}

	startTimeStamp, errS := getTimeType(startTime, today)
	endTimeStamp, errE := getTimeType(endTime, today)
	if errS != nil || errE != nil {
		return init, errors.New("开始时间或者结束时间不合规，请检查")
	}

	leftStart := startTimeStamp.Unix() - todayTimeBegin.Unix()
	leftEnd := endTimeStamp.Unix() - todayTimeBegin.Unix()

	ret := getTimeGourp(leftStart, leftEnd, 5)

	// 把时间转为5分钟一份
	for key, value := range ret {
		if value[0] > -1 {
			ret[key][0] = int64(math.Floor(float64(value[0]) / float64(60*5)))
		}
		if value[1] > -1 {
			ret[key][1] = int64(math.Ceil(float64(value[1]) / float64(60*5)))
		}
	}

	// 把开始时间到结束时间所有的占用的时间段 设置为1
	// var occupyByte = make([]uint64, 5)
	var occupyByte [5]uint64
	for k, v := range ret {
		if v[0] > -1 {
			occupyByte[k] = getWillUse(uint(v[0]), uint(v[1]))
		} else {
			occupyByte[k] = 0
		}
	}
	return occupyByte, nil
}

/**
 * 尝试占用资源
 * @param  {数组} tryUse        []uint64      尝试新占用的资源
 * @param  {数组} currentOccupy [5]uint64     当前占用的资源
 * 当error为nil时，返回项是成功的资源占用情况，
 * 当error不为nil时， 返回项是冲突明细情况
 * @return {[5]uint64,   error}                [冲突或者占用成功后的结果]
 * @author baiyouwen
 */
func BinaryOccupy(tryUse [5]uint64, currentOccupy [5]uint64) (ret [5]uint64, err error) {
	var occupy = [5]uint64{0, 0, 0, 0, 0}
	var conflict = [5]uint64{0, 0, 0, 0, 0}
	for k, v := range currentOccupy {
		if (v & tryUse[k]) > 0 {
			conflict[k] = v & tryUse[k]
			err = errors.New("有冲突")
		}
		occupy[k] = v ^ tryUse[k]
	}
	if err != nil {
		return conflict, err
	} else {
		return occupy, nil
	}
}

func getTimeGourp(startTime int64, endTime int64, splitTag int) []myarr {

	splitNum := splitNum(splitTag)
	slice1 := make([]myarr, splitNum)

	var tag int = 60 * 60 * splitTag
	var i int
	for i = 1; i < splitNum+1; i++ {
		if (int64(tag*i) - startTime) > int64(tag) {
			slice1[i-1][0] = 0
			if int64(tag*(i-1)) > endTime {
				slice1[i-1][0] = -1
			}
		} else {
			if int64(tag*i)-startTime <= 0 {
				slice1[i-1][0] = -1
			} else {
				slice1[i-1][0] = startTime - int64(tag*(i-1))
			}
		}

		if (int64(tag*i) - endTime) > int64(tag) {
			slice1[i-1][1] = -1
		} else {
			if int64(tag*i)-endTime < 0 {
				if int64(tag*i) > startTime {
					slice1[i-1][1] = int64(tag)
				} else {
					slice1[i-1][1] = -1
				}
			} else {
				slice1[i-1][1] = endTime - int64(tag*(i-1))
			}
		}
	}
	return slice1
}

/**
 * 把24小时分成多少份
 */
func splitNum(delimitNum int) int {
	total := int(math.Ceil(float64(24) / float64(delimitNum)))
	return total
}

/**
 * 生成数量为groupNum份的slice
 */
func bitGroup(groupNum int) []uint64 {
	var bitGroup []uint64 = make([]uint64, groupNum)
	return bitGroup
}

// 把开传入的开始时间和结束时间 加上当天，成为有效时间格式
func getTimeType(timeStr, today string) (time.Time, error) {
	startTimeStamp, err := time.Parse("2006-01-02 15:04:05", today+" "+timeStr)
	return startTimeStamp, err
}

// 占用位
func getWillUse(start, end uint) uint64 {
	var use uint64
	var base uint64 = 1
	var i uint = 0
	for i = start + 1; i <= end; i++ {
		use |= base << (i - 1)
	}
	return use
}

// 获取分份的数量
func getSequence(timeStamp int64) int {
	return int(math.Floor(float64(timeStamp) / float64(60*60*5)))
}
