package ard

func Copy(value interface{}) interface{} {
	switch value_ := value.(type) {
	case Map:
		map_ := make(Map)
		for key, value_ := range value_ {
			map_[key] = Copy(value_)
		}
		return map_

	case List:
		list := make(List, len(value_))
		for index, entry := range value_ {
			list[index] = Copy(entry)
		}
		return list
	}

	return value
}
