package main

import "fmt"

type Employee struct {
	EmployeeName string
	EmployeeId   int
	Salary       int
}

type Company struct {
	CompanyName string
	Emp         Employee
}

func (e *Employee) EmployeeDetail() {
	fmt.Println("Printing details:", e)
}

func NewEmployee(name string, Id, salary int) *Employee {
	return &Employee{
		EmployeeName: name,
		EmployeeId:   Id,
		Salary:       salary,
	}
}

func main() {
	//Anonymous struct
	// emp := struct {
	// 	EmployeeName string
	// 	EmployeeId   int
	// }{
	// 	EmployeeName: "Raj",
	// 	EmployeeId:   1234,
	// }
	// fmt.Println(emp)

	// struct Embedding

	// c := Company{
	// 	CompanyName: "ABC",
	// 	Emp: Employee{
	// 		EmployeeName: "XYZ",
	// 		EmployeeId:   1,
	// 		Salary:       3000,
	// 	},
	// }

	// fmt.Printf("Company details \n Companyname:%s Employee\n EmployeeName", c)
	emp1 := NewEmployee("Raj", 1, 49029394)
	emp1.EmployeeDetail()

}
