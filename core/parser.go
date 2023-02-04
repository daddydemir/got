package core

import (
	"github.com/daddydemir/got/models"
	"strconv"
	"strings"
)

var fields []string
var objects []models.Object
var object string

func ParseWithJson(request string) {
	getFields(request)
	toObject()
	structCreator("YeniObje")
}

func structCreator(name string) {
	def := "type " + name + " struct {\n"
	for _, obj := range objects {
		def += "\t" + obj.Name + "\t" + obj.Type + "\n"
	}
	def += " }"
	object = def
}

const (
	STRING = "string"
	NUMBER = "int"
	BOOL   = "bool"
	DOUBLE = "float32"
)

func getFields(r string) {
	lines := strings.Split(r, ",")
	field := ""
	collection := false
	dendenIndex := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lr := strings.Split(line, ":")
		for _, c := range lr {
			for l := 0; l < len(c); l++ {
				if string(c[l]) == "\"" {
					collection = !collection
					if field != "" && field != " " {
						fields = append(fields, field)
					}
					field = ""
					dendenIndex = l
					continue
				}

				if collection {
					field += string(c[l])
				}

				if l+1 == len(c) && string(c[l]) != "\"" && string(c[l]) != "}" {
					fields = append(fields, c[dendenIndex:l+1])
				}
			}
			dendenIndex = 0
		}
	}
}

func toObject() {

	var o models.Object
	for i := 0; i < len(fields); i += 2 {
		v := strings.Trim(fields[i+1], " ")
		if v == "false" || v == "true" {
			o = models.Object{Name: fields[i], Type: BOOL}
		} else if _, e := strconv.Atoi(v); e == nil {
			o = models.Object{Name: fields[i], Type: NUMBER}
		} else if _, e := strconv.ParseFloat(v, 32); e == nil {
			o = models.Object{Name: fields[i], Type: DOUBLE}
		} else {
			o = models.Object{Name: fields[i], Type: STRING}
		}
		objects = append(objects, o)

	}
}
