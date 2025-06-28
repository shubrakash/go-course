package basics

import (
	"fmt"
	"maps"
)

func mapsMain() {
	// var variableName map[keyType]valueType

	// var variableName = make(map[keyType]valueType)

	// variableName := make(map[keyType]valueType)

	// map literals
	map1 := make(map[int]string)
	fmt.Println(map1)

	map1[1] = "one"
	fmt.Println(map1)
	map1[2] = "two"
	fmt.Println(map1[1])
	fmt.Println(map1[3]) // non-existent key : returns ""(empty string); if value is <int>, 0 is returned

	map1[2] = "2"
	fmt.Println(map1)

	map1[3] = "three"
	fmt.Println(map1)

	delete(map1, 3)
	fmt.Println(map1)

	clear(map1)
	fmt.Println(map1)

	map1[1] = "one"
	map1[2] = "two"

	value, unknownValue := map1[1]
	fmt.Printf("value: %v, unknownValue: %v\n", value, unknownValue)

	_, doesEntryExist := map1[3]
	fmt.Println("doesEntryExist for key '3'? ::", doesEntryExist)

	///-----

	map2 := map[string]int{"key1": 1, "key2": 10, "key3": 30}
	fmt.Println(map2)

	map3 := map[string]int{"key1": 1, "key2": 10, "key3": 30}
	fmt.Println(map2)

	if maps.Equal(map3, map2) {
		fmt.Println("map3 and map2 are equal")
	}

	map3["key3"] = 20
	if !maps.Equal(map3, map2) {
		fmt.Println("map3 and map2 are not equal")
	}

	for _, v := range map3 {
		fmt.Println("value:", v)
	}

	_, ok := map3["key2"] // ok is a convention name, to check whether value against <key> exists or not

	if ok {
		fmt.Println("entry for key: 'key2' exists in map3")
	}

	var map4 map[string]string /// gets initialized to nil

	if map4 == nil {
		fmt.Println("map4 is intialized to nil value")
	} else {
		fmt.Println("map4 is not intialized to nil value")
	}

	val := map4["key"]
	fmt.Println("val:", val)

	// map4["key"] = "ten" // fails with a panic; can't assign to nil value
	// fmt.Println(map4)

	var map5 = make(map[string]string)
	map5["key"] = "value"
	fmt.Println("map5:", map5)
	fmt.Println("length of map5 is", len(map5)) // counts no. of keys

	var map6 = make(map[string]map[int]string)

	map6["map1"] = map1
	fmt.Println("map6:", map6)
}
