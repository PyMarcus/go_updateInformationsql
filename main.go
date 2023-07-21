package main 

import (
	"log"
	"time"
	"work_ss/Processos/include_foreign_keys_into_passenger/dao"
)

func main(){
	start := time.Now()
	dao.Insert()
	log.Println("Time spended: ", time.Since(start))
}
