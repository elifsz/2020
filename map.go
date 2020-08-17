package main

import "fmt"

func main() {
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "foo"
	myMap[12] = "bar"
	fmt.Println(myMap)

	states := make(map[string]string)
	states["IST"] = "istanbul"
	states["ANK"] = "ankara"
	states["ANT"] = "antalya "
	fmt.Println(states)

	antalya := states["ANT"]
	fmt.Println("seçili şehir: ", antalya)

	//silme işlemi
	delete(states, "ANK")
	fmt.Println(states)

	states["ERZ"] = "Erzurum"

	//döngüde yazdırma
	for k, v := range states {
		fmt.Printf("%v %v\n", k, v)
	}

	//states map nesnesinin elemanı kadar key listesi oluşturma
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)

	//key listesinde indexs değerlerine göre yazdırma
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
