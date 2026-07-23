package learntesting

import (
	"testing"
)

func TestReturnIntRole(t *testing.T) {
	var myRole Role = RoleGuest
	myRoleInt := ReturnIntOfRole(myRole)

	if myRoleInt != 0 {
		t.Errorf("Expected 0, got %d", myRoleInt)
	}
}

func TestIsAdmin(t *testing.T) {
	var myRole Role = RoleAdmin
	if !ReturnIsAdmin(myRole) {
		t.Errorf("User is admin - but function returns false!")
	}
}

func TestRequiresImmediateAttention(t *testing.T) {
	var p Priority = High
	if !RequiresImmediateAttention(p) {
		t.Errorf("Requires attention - return false")
	}
}

func TestRequiresImmediateAttentionFalse(t *testing.T) {
	var p Priority = Low
	if RequiresImmediateAttention(p) {
		t.Errorf("Does not require attention - returned true")
	}
}
