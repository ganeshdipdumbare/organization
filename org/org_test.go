package org

import (
	"organization/tree"
	"testing"
)

func TestCreateOrganization(t *testing.T) {
	org, err := CreateOrganization()
	if err != nil {
		t.Errorf("expected %v, got %v", "no error", err)
	}

	if org == nil {
		t.Errorf("expected %v, got %v", "organization with non nil", nil)
	}
}

func TestAddEmployeesToManager(t *testing.T) {
	root, err := tree.NewNode(1, []*tree.Node{})
	if err != nil {
		t.Errorf("expected %v, got %v for error while calling NewNode", "nil", "not nil")
	}
	org := tree.CreateTree(root)

	err = AddEmployeesToManager(1, []int{2, 3, 4}, org)
	if err != nil {
		t.Errorf("expected %v, got %v for error while calling AddEmployeesToManager", "nil", "not nil")
	}

	if len(org.Root.Child) != 3 {
		t.Errorf("expected %v, got %v for no of child to emp 1", 3, len(org.Root.Child))
	}
}

func TestFindNearestManager(t *testing.T) {
	root, err := tree.NewNode(1, []*tree.Node{})
	if err != nil {
		t.Errorf("expected %v, got %v for error while calling NewNode", "nil", "not nil")
	}
	org := tree.CreateTree(root)

	err = AddEmployeesToManager(1, []int{2, 3, 4}, org)
	if err != nil {
		t.Errorf("expected %v, got %v for error while calling AddEmployeesToManager", "nil", "not nil")
	}

	manager, err := FindNearestManager(2, 3, org)
	if err != nil {
		t.Errorf("expected %v, got %v for error while calling FindNearestManager", "nil", "not nil")
	}
	if manager != 0 {
		t.Errorf("expected %v, got %v for manager value", 0, manager)
	}
}
