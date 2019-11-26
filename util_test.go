package gutil_test

import (
	"fmt"
	"gutil"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var key = []byte("1234567800009999")
	var datas = []string{"123456", "abcdefg", "中国", `{"grpc_addr":"localhost:8888"}`}
	var miwen = make([]string, 0, len(datas))
	for _, s := range datas {
		if _s, e := gutil.AesEncrypt([]byte(s), key); e != nil {
			miwen = append(miwen, "")
			t.Error(e)
		} else {
			miwen = append(miwen, _s)
		}
	}

	for i, s := range miwen {
		ors, _ := gutil.AesDecrypt(s, key)
		if string(ors) != datas[i] {
			t.Error("not eq")
		}
	}
}
func TestConvert(t *testing.T) {
	times := []int64{1574734026754, 1574734026, 1000000, 9555566666}
	for _, v := range times {
		var ta time.Time
		tt := gutil.CInt2Time(v, ta)
		fmt.Println(tt, ta, ta == tt)
	}
	objs := []interface{}{1, "a", true, time.Now(), gutil.NewPager(10, 20, 50, 60, nil)}
	for _, v := range objs {
		fmt.Printf("%s\r\n", gutil.CObject2JsonStr(v))
	}
	fmt.Println("over")
}
