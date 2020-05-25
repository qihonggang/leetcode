package design_pattern

import "sync"

type Singleton struct{}

var singleton *Singleton
var once sync.Once

func GetInstence() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}