package main

import "fmt"

type Employee interface {
	calculateSalary() int
}

type (
	Fulltimer struct {
		name       string
		dailyWage  int
		daysWorked int
	}
	Contractor struct {
		name       string
		dailyWage  int
		daysWorked int
	}
	Freelancer struct {
		name        string
		hourlyWage  int
		hoursWorked int
	}
)

func (f Fulltimer) calculateSalary() int {
	return f.dailyWage * f.daysWorked
}

func (c Contractor) calculateSalary() int {
	return c.dailyWage * c.daysWorked
}

func (f Freelancer) calculateSalary() int {
	return f.hourlyWage * f.hoursWorked
}

func getSalary(e Employee) int {
	return e.calculateSalary()
}

func main() {
	e1 := Fulltimer{"Yash", 1000, 30}
	e2 := Contractor{"Pavan", 500, 24}
	e3 := Freelancer{"Naman", 100, 150}
	fmt.Println("Salary of", e1.name, "is", getSalary(e1))
	fmt.Println("Salary of", e2.name, "is", getSalary(e2))
	fmt.Println("Salary of", e3.name, "is", getSalary(e3))
}
