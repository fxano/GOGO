package main

import (
	"fmt"
	"sync"
)
//func main(){
//	for i:=1;i<10 ;i++  {
//		for s:=1;s<=i ;s++  {
//			fmt.Printf("%d*%d=%d\t",i,s,i*s)
//		}
//		fmt.Println()
//	}
//}






var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}






func one(){
		for i:=0;i<100 ;i++  {


			if i==10{
				goto secend
			}
		}

		secend:
			fmt.Print("2")
}