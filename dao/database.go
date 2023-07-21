package dao

import (
	"auto_sql/connection"
	"auto_sql/settings"
	"fmt"
	"log"
	"time"
)

type SpecificTable struct {
	id         int
	rfid       string
	local_time time.Time
	unit       int
	positionId int
}

// select id from dev_status_30
func selectData() []*SpecificTable {

	data := []*SpecificTable{}

	conn, e := connection.ConnectWithDatabase()

	defer conn.Close()

	if e != nil {
		conn.Close()
	}

	for _, date := range settings.GenerateDateRange("2023-01-01", settings.DateNow()) {
		log.Println("Selecting data from ", date)
		rows, err := conn.Query(fmt.Sprintf(`
									SELECT d.id, d.prfid, d.local_time, d.unit, d.position_id
										FROM specificTable d
									JOIN another_table pb ON
									d.td = pb.td
									AND d.local_time = pb.local_time
									AND d.unit = pb.unit
									AND pb.rfid = d.prfid 
									AND pb.position_id = d.position_id
									AND to_char(pb.local_time, 'YYYY-MM-DD') = '%s' LIMIT 10;
								`, date))
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			ds30 := &SpecificTable{}

			rows.Scan(&ds30.id, &ds30.rfid, &ds30.local_time, &ds30.unit, &ds30.positionId)

			data = append(data, ds30)
		}
	}
	return data
}

func Insert() {
	inserts := []string{}

	data := selectData()

	for _, d := range data {
		inserts = append(inserts, fmt.Sprintf(`
		UPDATE anotherTable set id = %d
		WHERE rfid = '%s' and local_time = '%v' and unit = %d and poi_id = %d;`, d.id, d.rfid,
			d.local_time.Format("2006-01-02 15:04:05"), d.unit, d.positionId))
	}

	for _, v := range inserts {
		fmt.Println(v)
	}

	conn, e := connection.ConnectWithDatabase()

	defer conn.Close()

	if e != nil {
		log.Fatalln("Fail to get connection to insert data")
	}

	for _, insert := range inserts {
		conn.Exec(insert)
		log.Println(insert)
	}

	log.Println("Successfully")

}
