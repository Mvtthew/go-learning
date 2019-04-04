package myPackage

type person struct {
	name string
	age  int
}

type world struct {
	persons []person
	name    string
}

func NewPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func NewWorld(name string) world {
	return world{
		persons: nil,
		name:    name,
	}
}

func (w *world) AddToWorld(person person)  {
	w.persons = append(w.persons, person)
}