package core

import (
	"fmt"
	"github.com/daddydemir/got/models"
	"strconv"
	"strings"
)

var fields []string
var objects []models.Object
var Objects []models.Object
var object string
var getters []string
var setters []string

// problem is array reference,
func getterCreator(modelName string) {
	def := ""
	for i, obj := range objects {
		def += "func (" + string(obj.Name[0:1]) + " " + modelName + ") " + Objects[i].Name + "() " + obj.Type + " { \n"
		def += "\treturn " + string(obj.Name[0:1]) + "." + Objects[i].Name + "\n"
		def += "}\n"
		getters = append(getters, def)
		println(def)
		def = ""
	}
}

func setterCreator(modelName string) {
	def := ""
	for i, obj := range objects {
		def += "func (" + string(obj.Name[0:1]) + " " + modelName + ") " + "Set" + Objects[i].Name + "(" + obj.Name + " " + obj.Type + ") {\n"
		def += "\t " + string(obj.Name[0:1]) + "." + obj.Name + " = " + obj.Name + "\n"
		def += "}\n"
		setters = append(setters, def)
		fmt.Println(def)
		def = ""
	}
}
func ParseWithJson(request string) {
	getFields(request)
	toObject()
	serialize()
	structCreator("Model")
	getterCreator("Model")
	setterCreator("Model")
	fmt.Println(object)
}

func structCreator(name string) {
	def := "type " + name + " struct {\n"
	for _, obj := range Objects {
		def += "\t" + obj.Name + "\t" + obj.Type + "\n"
	}
	def += " }"
	object = def
}

func initialLetterLarger(v string) string {
	temp := ""
	for i := 0; i < len(v); i++ {
		if i == 0 {
			temp += strings.ToLower(string(v[i]))
		} else {
			temp += string(v[i])
		}
	}
	return temp
}

func serialize() {

	for i, v := range Objects {
		v.Name = initialLetterLarger(v.Name)
		if v.Type == BOOL {
			if v.Name[0:2] == "is" {
				Objects[i].Name = initialLetterLarger(v.Name)
			}
		} else {
			Objects[i].Name = initialLetterLarger(v.Name)
		}
	}
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
	Objects = objects
}
