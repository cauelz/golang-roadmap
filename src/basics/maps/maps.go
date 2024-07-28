package mapbasics


type Person struct {
	Name string
	Age int
}

type personMappedByName map[string]Person


func CreatePersonMap(n int) personMappedByName {
	return make(personMappedByName, n)
}

func (pm personMappedByName) AddPerson(key string, p Person) {
	pm[key] = p
}

func (pm personMappedByName) HasKey(key string) bool {
	_, ok := pm[key]
	return ok
}

func (pm personMappedByName) GetPerson(key string) (Person, bool) {
	p, ok := pm[key]
	return p, ok
}

func (pm personMappedByName) DeletePerson(key string) {
	delete(pm, key)
}

func (pm personMappedByName) GetKeys() []string {
	keys := make([]string, 0, len(pm))
	for k := range pm {
		keys = append(keys, k)
	}
	return keys
}

func (pm personMappedByName) GetValues() []Person {
	values := make([]Person, 0, len(pm))
	for _, v := range pm {
		values = append(values, v)
	}
	return values
}