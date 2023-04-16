package core

import (
	"log"
	"os"
	"strings"

	"github.com/daddydemir/got/models"
)

var fields []string
var objects []models.Object
var Objects []models.Object
var object string
var getters []string
var setters []string
var get string
var set string
var mod string
var nameOfModel string

// problem is array reference,
func getterCreator(modelName string) {
	def := ""
	tmp := ""
	for i, obj := range objects {
		def += "func (" + string(nameOfModel[0:1]) + " " + modelName + ") " + strings.Title(strings.ToLower(Objects[i].Name)) + "() " + obj.Type + " { \n"
		def += "\treturn " + string(nameOfModel[0:1]) + "." + Objects[i].Name + "\n"
		def += "}\n\n"
		getters = append(getters, def)
		tmp += def
		def = ""
	}
	get = tmp
}

func setterCreator(modelName string) {
	def := ""
	tmp := ""
	for i, obj := range objects {
		def += "func (" + string(nameOfModel[0:1]) + " *" + modelName + ") " + "Set" + strings.Title(strings.ToLower(Objects[i].Name)) + "(" + obj.Name + " " + obj.Type + ") {\n"
		def += "\t " + string(nameOfModel[0:1]) + "." + obj.Name + " = " + obj.Name + "\n"
		def += "}\n\n"
		setters = append(setters, def)
		tmp += def
		def = ""
	}
	set = tmp
}

func ParseWithJson(request string, name string) string {
	nameOfModel = name
	Creator(request)
	objects = Objects
	structCreator(name)
	getterCreator(name)
	setterCreator(name)
	createAll(name)
	return mod + get + set
}

func structCreator(name string) {
	def := "package models\n\ntype " + name + " struct {\n"
	for _, obj := range Objects {
		def += "\t" + obj.Name + "\t" + obj.Type + "\n"
	}
	def += " }\n"
	object = def
	mod = def
}

const (
	STRING = "string"
	NUMBER = "int"
	BOOL   = "bool"
	DOUBLE = "float32"
)

func createAll(name string) {
	err := os.WriteFile("./output/"+name+".go", []byte(mod+get+set), 0644)
	if err != nil {
		log.Println(err)
	}
	log.Println(name, "is created at location:", "./output/"+name+".go")
}
