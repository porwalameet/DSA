package main

type Person struct {
	name     string
	isAlive  bool
	children []*Person
}

type Monarch struct {
	king *Person
}

func NewMonarch(name string) *Monarch {
	return &Monarch{king: &Person{name: name, isAlive: true}}
}

func (m *Monarch) Birth(childName string, parentName string) {
	parent := m.findPerson(m.king, parentName)
	if parent == nil {
		return
	}
	child := &Person{name: childName, isAlive: true}
	parent.children = append(parent.children, child)
}

func (m *Monarch) Death(name string) {
	person := m.findPerson(m.king, name)
	if person == nil {
		return
	}
	person.isAlive = false
}

func (m *Monarch) Succession() []string {
	successors := []string{}
	m.dfs(m.king, &successors)
	return successors
}

func (m *Monarch) findPerson(person *Person, name string) *Person {
	if person.name == name {
		return person
	}
	for _, child := range person.children {
		if p := m.findPerson(child, name); p != nil {
			return p
		}
	}
	return nil
}

func (m *Monarch) dfs(person *Person, successors *[]string) {
	if person.isAlive {
		*successors = append(*successors, person.name)
	}
	for _, child := range person.children {
		m.dfs(child, successors)
	}
}

func main() {
	m := NewMonarch("jake")
	m.Birth("charlie", "jake")
	m.Birth("lucy", "jake")
	m.Birth("mark", "charlie")
	m.Birth("sophie", "charlie")
	m.Birth("james", "lucy")
	m.Birth("lily", "lucy")
	m.Death("charlie")
	successors := m.Succession()
	for _, s := range successors {
		println(s)
	}
}
