package org

import (
	"fmt"
	"log"
	"organization/tree"
)

func CreateOrganization() (*tree.Tree, error) {
	ceo, err := tree.NewNode(0, []*tree.Node{})
	if err != nil {
		return nil, err
	}
	org := tree.CreateTree(ceo)
	return org, nil
}

func AddEmployeesToManager(manager int, employees []int, organization *tree.Tree) error {
	var err error
	var emp *tree.Node
	empnode := []*tree.Node{}

	for _, employee := range employees {
		emp = organization.FindNode(employee)
		if emp == nil {
			emp, err = tree.NewNode(employee, []*tree.Node{})
			if err != nil {
				log.Println("error occurred:", err)
				return err
			}
		}
		empnode = append(empnode, emp)
	}

	mngrnode := organization.FindNode(manager)
	if mngrnode == nil {
		log.Println("manager not found")
		return fmt.Errorf("manager not found")
	}

	mngrnode.AddChildren(empnode...)
	return nil
}

func FindNearestManager(emp1, emp2 int, organization *tree.Tree) (int, error) {
	var managerId int
	found := false

	pathOfEmp1 := organization.FindShortestPath(emp1)
	pathOfEmp2 := organization.FindShortestPath(emp2)
	if len(pathOfEmp1) == 0 || len(pathOfEmp2) == 0 {
		log.Println("given employee not found in the organization")
		return 0, fmt.Errorf("employ not found in the organization")
	}

	minLength := len(pathOfEmp1)
	if minLength > len(pathOfEmp2) {
		minLength = len(pathOfEmp2)
	}

	for i := 0; i < minLength; i++ {
		if pathOfEmp1[i] != pathOfEmp2[i] {
			managerId = pathOfEmp1[i-1]
			found = true
			break
		}
	}

	if !found {
		managerId = pathOfEmp1[minLength-1]
	}
	return managerId, nil
}

func PrintOrgDetails(organization *tree.Tree) {
	if organization == nil {
		return
	}

	organization.PrintPreOrder()
}
