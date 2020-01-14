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

func TestWriteFile(t *testing.T) {
	ts := time.Now().Format("06-01-02 15:04:05.000000")
	if err := gutil.WriteFileByOverwrite("./abc/123.txt", []byte(ts+"abc123")); err != nil {
		t.Errorf("WriteFileByOverwrite [./abc/123.txt] error %v \r\n", err)
	}
	if err := gutil.WriteFileByOverwrite("./456.txt", []byte(ts+"456456")); err != nil {
		t.Errorf("WriteFileByOverwrite [./456.txt] error %v \r\n", err)
	}
	if err := gutil.WriteFileByAppend("./abc/123.txt", "Hello world "+ts); err != nil {
		t.Errorf("WriteFileByAppend [./abc/123.txt] error %v \r\n", err)
	}
}
