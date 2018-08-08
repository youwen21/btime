# btime包是一个GO语言开发的二进制时间包

## btime说明 

btime包把24小时分成5份， 分别为：1－5，5－10，10－15，55－20，20－24  
其中每一小份按5分钟为一个单位，每小时12个单位时间  
每份为 12 * 5 ＝60 ，需要60个位就能表示 5个小时的时间


### DEMO1
```go

func main() {
    // 获取开始时间，结束时间
    startTime := "07:00:00"
    endTime := "08:00:00"

    var current = [5]uint64{0, 4503530907893760, 0, 0, 0}

    ret, err := btime.GetBinary(startTime, endTime)
    if err != nil {
        fmt.Println(err)
    }
    for _, v := range ret {
        fmt.Println(biu.ToBinaryString(v))
        fmt.Println(v)
    }

    tryuse, err := btime.BinaryOccupy(ret, current)
    if err != nil {
        fmt.Println(err)
        fmt.Println(tryuse)
    } else {
        fmt.Println(tryuse)
        for _, v := range tryuse {
            fmt.Println(biu.ToBinaryString(v))
        }
    }
}

```

### DEMO2
```go
func test2() {
    // 获取开始时间，结束时间
    startTime := "07:00:00"
    endTime := "10:00:00"

    var current = [5]uint64{0, 4503530907893760, 0, 0, 0}
    ret, err := btime.TryOccupy(startTime, endTime, current)
    if err != nil {
        fmt.Println(err)
    }
    for _, v := range ret {
        fmt.Println(biu.ToBinaryString(v))
        fmt.Println(v)
    }
}
```