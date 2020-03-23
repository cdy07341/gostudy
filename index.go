package main

import (
	"fmt"
	"sync"
	"time"

	"icode.baidu.com/baidu/gdp/gdp"
	"icode.baidu.com/baidu/gdp/gdp/conf"
)

var config *gdp.AppConfig

func main() {

	var wg sync.WaitGroup
	start := time.Now()
	var rwl sync.RWMutex
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {

			if config == nil {
				rwl.Lock()
				if config == nil {
					fmt.Println("config is not set")
					config = &gdp.AppConfig{}
					if err := conf.ReadFile("app.toml", config); err != nil {
						panic(fmt.Sprintf("Can't load config file %s: %s", "app.toml", err.Error()))
					}
				}
				rwl.Unlock()
			}
			rwl.RLock()
			fmt.Println(config)
			rwl.RUnlock()
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(config)
}
