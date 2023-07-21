package settings

import (
	"errors"
	"log"
	"strconv"
)

type DataBaseCredentials struct {
	Host     string
	Password string
	Port     int
	User     string
	Database string
}

// ReadIniFile: fileName: file.ini, section: production, key: password
func readIniFile(fileName string, section string, key string) (string, error) {
	cfg, err := ini.Load(fileName)

	if err != nil {
		log.Panic(err)
		return "", err
	}

	sectionn := cfg.Section(section)

	if sectionn == nil {
		log.Panic("Section not found!")
		return "", errors.New("Section not found!")
	}
	k := sectionn.Key(key)
	if k == nil {
		return "", errors.New("Key not found!")
	}
	return k.String(), nil
}

func GetCredentials() *DataBaseCredentials {
	host, err := readIniFile("C:\\postgres.ini", "desenv", "host")
	password, err := readIniFile("C:\\postgres.ini", "desenv", "password")
	user, err := readIniFile("C:\\postgres.ini", "desenv", "user")
	database, err := readIniFile("C:\\postgres.ini", "desenv", "database")
	port, err := readIniFile("C:\\postgres.ini", "desenv", "port")
	if err != nil {
		log.Println(err)
	}
	p, _ := strconv.Atoi(port)
	return &DataBaseCredentials{
		User:     user,
		Host:     host,
		Password: password,
		Database: database,
		Port:     p,
	}
}
