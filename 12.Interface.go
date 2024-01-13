package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type staff1 struct {
	empID    int
	basicPay int
}

type staff2 struct {
	empID    int
	basicPay int
	bonus    int
}

func (s staff1) CalculateSalary() int {
	return s.basicPay
}

func (s staff2) CalculateSalary() int {
	return s.basicPay + s.bonus
}

func totalExpense(s []SalaryCalculator) int {

	//类型断言
	for _, value := range s {
		switch value.(type) {
		case staff1:
			println("员工1正在计算中")
		case staff2:
			println("员工2计算中")
		default:
			println("未知错误")
		}
	}

	expense := 0
	for _, value := range s {
		expense = expense + value.CalculateSalary()
	}
	return expense
}

func main() {
	s1 := staff1{empID: 1001, basicPay: 100}
	s2 := staff1{empID: 1002, basicPay: 100}
	s3 := staff2{empID: 1003, basicPay: 100, bonus: 100}

	salCal := []SalaryCalculator{s1, s2, s3}
	expense := totalExpense(salCal)
	fmt.Println("总支出：", expense)

}
