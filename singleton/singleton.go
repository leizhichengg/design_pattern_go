package singleton

import "sync"

type Instance struct{}

var instance *Instance
var instanceLock sync.Mutex
var once sync.Once

func GetInstance1() *Instance {
	if instance == nil {
		instanceLock.Lock()
		defer instanceLock.Unlock()
		if instance == nil {
			instance = &Instance{}
		}
	}
	return instance
}

func GetInstance2() *Instance {
	once.Do(func() {
		instance = &Instance{}
	})
	return instance
}
