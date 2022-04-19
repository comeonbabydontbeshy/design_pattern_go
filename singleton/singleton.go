package singleton

import "sync"

type Singleton struct{}

// 懒汉式单例模式
var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(
			func() {
				lazySingleton = &Singleton{}
			})
	}
	return lazySingleton
}
