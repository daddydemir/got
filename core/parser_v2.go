package core

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/daddydemir/got/models"
)

var modelList []models.Object

func Creator(data string) {
	c := make(map[string]json.RawMessage)
	log.Println("gelen :", data)
	err := json.Unmarshal([]byte(data), &c)
	if err != nil {
		log.Println("json'a cevirirken hata: ", err)
	}

	keys := make([]string, len(c))
	list := make([]models.Object, len(c))

	i := 0
	temp := ""
	for s, _ := range c {
		keys[i] = s
		temp = string(c[s])
		if string(temp[0]) == "\"" && string(temp[len(temp)-1]) == "\"" {
			list[i] = models.Object{Name: s, Type: STRING}
		} else if string(temp) == "true" || "false" == string(temp) {
			list[i] = models.Object{Name: s, Type: BOOL}
		} else if _, e := strconv.Atoi(string(temp)); e == nil {
			list[i] = models.Object{Name: s, Type: NUMBER}
		} else if _, e := strconv.ParseFloat(string(temp), 32); e == nil {
			list[i] = models.Object{Name: s, Type: DOUBLE}
		} else {
			log.Println("hata")
		}
		i++
	}
	Objects = list
}
