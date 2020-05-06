package main

import (
	"fmt"
	"log"
	"organization/org"
	"organization/tree"
)

func main() {
	var choice int
	var err error
	var organization *tree.Tree

	for {
		printMenu()
		fmt.Scanf("%v", &choice)
		switch choice {
		case 0:
			organization, err = org.CreateOrganization()
			if err != nil {
				log.Fatal("error occurred: ", err)
			}
			fmt.Println("--- organization is created successfully, CEO employee id is 0")
		case 1:
			err = addEmployee(organization)
			if err != nil {
				log.Fatal("error occurred: ", err)
			}
			fmt.Println("--- employee successfully added to given manager")
		case 2:
			err = findNearestManager(organization)
			if err != nil {
				log.Fatal("error occurred: ", err)
			}
		case 3:
			fmt.Println("--Organization details-")
			org.PrintOrgDetails(organization)
		case 4:
			fmt.Println("Exiting....")
			return
		default:
			fmt.Println("wrong choice, please select again.")
			continue
		}
	}
}

func findNearestManager(organization *tree.Tree) error {
	var emp1, emp2 int
	fmt.Println("Enter 2 employee ids seperated by space")
	fmt.Scanf("%d %d", &emp1, &emp2)

	manager, err := org.FindNearestManager(emp1, emp2, organization)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Printf("--- Nearest manager for %v and %v is %v\n", emp1, emp2, manager)
	return nil
}

func addEmployee(organization *tree.Tree) error {
	var managerId, empCount int
	empId := []int{}
	fmt.Println("Enter manager id:")
	fmt.Scanf("%d", &managerId)

	fmt.Println("Enter no of employees to be added:")
	fmt.Scanf("%d", &empCount)

	fmt.Println("Enter employee id seperated by space")
	var v int
	for i := 0; i < empCount; i++ {
		fmt.Scanf("%d", &v)
		empId = append(empId, v)
	}

	err := org.AddEmployeesToManager(managerId, empId, organization)
	if err != nil {
		log.Println("error occurred: ", err)
		return err
	}
	return nil
}

func printMenu() {
	fmt.Println("Menu")
	fmt.Println("0. Create organization with CEO(call only once)")
	fmt.Println("1. Add employees to manager")
	fmt.Println("2. Find nearest manager")
	fmt.Println("3. Print organization details")
	fmt.Println("4. Exit")
	fmt.Println("Enter your choice:")
}
