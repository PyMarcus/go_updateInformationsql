package main 

import (
	"log"
	"time"
	"auto_sql/dao"
)

func main(){
	start := time.Now()
	dao.Insert()
	log.Println("Time spended: ", time.Since(start))
}
