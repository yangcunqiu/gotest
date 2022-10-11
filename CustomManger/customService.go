package main

type custom struct {
	id        int
	name      string
	gender    string
	age       int
	tel       string
	email     string
	isDeleted bool
}

var customList []custom

func add(cus custom) {
	customList = append(customList, cus)
}

func del(id int) {
	customList[id].isDeleted = true
}

func update(cus custom) {
	customList[cus.id] = cus
}

func list() []custom {
	var temp []custom
	for _, cus := range customList {
		if !cus.isDeleted {
			temp = append(temp, cus)
		}
	}
	return temp
}

func get(id int) custom {
	return customList[id]
}

//func (cus custom) String() string {
//	return
//}
