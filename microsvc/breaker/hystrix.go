package breaker

import "github.com/afex/hystrix-go/hystrix"

type hystrixBreaker struct {
	name string
}

func newHystrixBreaker(name string) *hystrixBreaker {
	hystrix.ConfigureCommand(name, hystrix.CommandConfig{
		// 执行 command 的超时时间
		Timeout: 10,

		// 最大并发量
		MaxConcurrentRequests: 100,

		// 一个统计窗口 10 秒内请求数量
		// 达到这个请求数量后才去判断是否要开启熔断
		RequestVolumeThreshold: 10,

		// 熔断器被打开后
		// SleepWindow 的时间就是控制过多久后去尝试服务是否可用了
		// 单位为毫秒
		SleepWindow: 500,

		// 错误百分比
		// 请求数量大于等于 RequestVolumeThreshold 并且错误率到达这个百分比后就会启动熔断
		ErrorPercentThreshold: 20,
	})

	return &hystrixBreaker{
		name: name,
	}
}

func (b *hystrixBreaker) Do(fn func() error, onBreakerOpen func(error) error) error {
	return hystrix.Do(b.name, fn, onBreakerOpen)
}

func (b *hystrixBreaker) Go(fn func() error, onBreakerOpen func(error) error) error {
	hystrix.Go(b.name, fn, onBreakerOpen)

	return nil
}
