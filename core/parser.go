package core

func ParseWithJson(request string) {
	//structCreator(request)
	getFields(request)
}

func structCreator(name string) {
	def := "type " + name + " struct {"

	def += " }"
	println(def)
}

func getFields(r string) {
	status := false
	fieldStatus := false
	field := ""
	var fields []string
	var realFields []string
	lastIndex := 0
	for i := 0; i < len(r); i++ {
		if string(r[i]) == "{" {
			status = true
		} else if string(r[i]) == "}" {
			status = false
		}

		if status {
			if string(r[i]) == "\"" {
				fieldStatus = !fieldStatus
				temp := string(r[lastIndex:i])
				println(temp)
				continue
			}
		}

		if string(r[i]) == ":" {
			lastIndex = i + 1
			/*
				i = index
			*/
		}

		if fieldStatus {
			field += string(r[i])
		} else {
			if field != "" {
				//fmt.Println(field)
				fields = append(fields, field)
			}
			field = ""
		}

		if !fieldStatus && string(r[i]) == ":" {
			realFields = append(realFields, fields[len(fields)-1])
		}
	}
	//fmt.Println(fields)
	//fmt.Println(realFields)
}
