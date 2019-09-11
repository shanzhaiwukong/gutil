package gutil

import (
	"time"
)

// TimerOut 【异步】设置一个超时
//成功执行回调succ
//超时执行回调faild
//退出自动关闭chan
func TimerOut(a chan interface{}, timeout time.Duration, success func(data interface{}), faild func()) {
	ticker := time.NewTimer(timeout)
	go func() {
		defer close(a)
		for {
			select {
			case msg := <-a:
				success(msg)
				return
			case <-ticker.C:
				faild()
				return
			}
		}
	}()
}

// TimerOutSync 【同步】设置一个超时
//成功执行回调succ
//超时执行回调faild
//退出自动关闭chan
func TimerOutSync(a chan interface{}, timeout time.Duration, success func(data interface{}), faild func()) {
	ticker := time.NewTimer(timeout)
	defer close(a)
	for {
		select {
		case msg := <-a:
			success(msg)
			return
		case <-ticker.C:
			faild()
			return
		}
	}
}

// TimesOut 【异步】设置一个执行N次的计时器
//定时执行任务 达到指定数量退出 并回调
//返回一个chan 用于写入值来取消计时
//step 间隔回调 over 执行完毕回调 cancle取消回调
//退出自动关闭通道
func TimesOut(times int32, timed time.Duration, step func(steps int32), over func(), cancle func(steps int32)) (isCancle chan bool) {
	ticker := time.NewTicker(timed)
	isCancle = make(chan bool)
	go func() {
		var currentCount int32 = 0
		defer func() {
			ticker.Stop()
			close(isCancle)
		}()
		for {
			select {
			case <-isCancle:
				cancle(currentCount)
				return
			case <-ticker.C:
				currentCount++
				step(currentCount)
				if currentCount >= times {
					over()
					return
				}
			}
		}
	}()
	return
}

// TimesOutSync 【同步】设置一个执行N次的计时器
//定时执行任务 达到指定数量退出 并回调
//返回一个chan 用于写入值来取消计时
//step 间隔回调 over 执行完毕回调 cancle取消回调
//退出自动关闭通道
func TimesOutSync(times int32, timed time.Duration, step func(steps int32), over func(), cancle func(steps int32)) (isCancle chan bool) {
	ticker := time.NewTicker(timed)
	isCancle = make(chan bool)
	var currentCount int32 = 0
	defer func() {
		ticker.Stop()
		close(isCancle)
	}()
	for {
		select {
		case <-isCancle:
			cancle(currentCount)
			return
		case <-ticker.C:
			currentCount++
			step(currentCount)
			if currentCount >= times {
				over()
				return
			}
		}
	}
}
