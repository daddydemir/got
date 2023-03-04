package core

import (
	"fmt"

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

// problem is array reference,
func getterCreator(modelName string) {
	fmt.Println("getterCreator is start --")
	def := ""
	tmp := ""
	for i, obj := range objects {
		def += "func (" + string(obj.Name[0:1]) + " " + modelName + ") " + Objects[i].Name + "() " + obj.Type + " { \n"
		def += "\treturn " + string(obj.Name[0:1]) + "." + Objects[i].Name + "\n"
		def += "}\n\n"
		getters = append(getters, def)
		tmp += def
		def = ""
	}
	get = tmp
	fmt.Println(get)
}

func setterCreator(modelName string) {
	fmt.Println("setterCreator is start --")
	def := ""
	tmp := ""
	for i, obj := range objects {
		def += "func (" + string(obj.Name[0:1]) + " *" + modelName + ") " + "Set" + Objects[i].Name + "(" + obj.Name + " " + obj.Type + ") {\n"
		def += "\t " + string(obj.Name[0:1]) + "." + obj.Name + " = " + obj.Name + "\n"
		def += "}\n\n"
		setters = append(setters, def)
		tmp += def
		def = ""
	}
	set = tmp
	fmt.Println(set)
}
func ParseWithJson(request string) {
	Creator(request)
	objects = Objects
	structCreator("Model")
	getterCreator("Model")
	setterCreator("Model")
}

func structCreator(name string) {
	fmt.Println("structCreator is start --")
	def := "type " + name + " struct {\n"
	for _, obj := range Objects {
		def += "\t" + obj.Name + "\t" + obj.Type + "\n"
	}
	def += " }\n"
	object = def
	mod = def
	fmt.Println(def)
}

const (
	STRING = "string"
	NUMBER = "int"
	BOOL   = "bool"
	DOUBLE = "float32"
)
