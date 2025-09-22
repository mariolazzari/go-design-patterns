package main

import "fmt"

// Component interface
type Member interface {
	printMemberInfo()
}

// Leaf
type TeamMember struct {
	name       string
	teamNumber int
	position   string
}

func (t *TeamMember) printMemberInfo() {
	fmt.Printf("Name: %s Team Number: %d Position: %s\n", t.name, t.teamNumber, t.position)
}

// Composite
type Roster struct {
	members []Member
	name    string
}

// We'll implement the printMemberInfo method
func (r *Roster) printMemberInfo() {
	fmt.Println("Here's the roster for team: " + r.name)
	for _, member := range r.members {
		member.printMemberInfo()
	}
}

func (r *Roster) add(m Member) {
	r.members = append(r.members, m)
}

func main() {

	member1 := &TeamMember{name: "Johnny Rocket", teamNumber: 12, position: "Forward"}
	member2 := &TeamMember{name: "Tim Hoops", teamNumber: 24, position: "Point Guard"}
	member3 := &TeamMember{name: "Billy Banks", teamNumber: 29, position: "Shooting Guard"}

	roster := &Roster{
		name: "Bobcats",
	}

	roster.add(member1)
	roster.add(member2)
	roster.add(member3)

	roster.printMemberInfo()
}
