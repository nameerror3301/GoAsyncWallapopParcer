package main

import (
	"GoAsyncWallapopParcer/internal/app"
	"runtime"

	"github.com/sirupsen/logrus"
)

// type G struct {
// 	mx  sync.Mutex
// 	Arr []int
// }

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: false,
	})
}

// func (g *G) AddNum(num int) {
// 	g.mx.Lock()
// 	defer g.mx.Unlock()

// 	g.Arr = append(g.Arr, num)
// }

// func (g *G) GetNum() {
// 	for _, val := range g.Arr {
// 		fmt.Println(val)
// 	}
// }

// var wg sync.WaitGroup
// var g G

//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go func(i int) {
//			g.AddNum(i)
//			defer wg.Done()
//		}(i)
//	}
//
// wg.Wait()
// 2.136161837s
func main() {
	app.Run()
}
