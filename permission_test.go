package vbcore

import "testing"

func TestPermissionItoA(t *testing.T) {
	if PermissionItoA(PermissionBanned) != PermissionBannedString {
		t.Fail()
	}
	if PermissionItoA(PermissionDefault) != PermissionDefaultString {
		t.Fail()
	}
	if PermissionItoA(PermissionVerified) != PermissionVerifiedString {
		t.Fail()
	}
	if PermissionItoA(PermissionTeam) != PermissionTeamString {
		t.Fail()
	}
	if PermissionItoA(PermissionAdmin) != PermissionAdminString {
		t.Fail()
	}
}

func TestPermissionItoA_Lesser(t *testing.T) {
	if PermissionItoA(PermissionBanned-1) != PermissionBannedString {
		t.Fail()
	}
}

func TestPermissionItoA_Bigger(t *testing.T) {
	if PermissionItoA(PermissionAdmin+1) != PermissionAdminString {
		t.Fail()
	}
}

func TestPermissionAtoI(t *testing.T) {
	if PermissionAtoI(PermissionBannedString) != PermissionBanned {
		t.Fail()
	}
	if PermissionAtoI(PermissionDefaultString) != PermissionDefault {
		t.Fail()
	}
	if PermissionAtoI(PermissionVerifiedString) != PermissionVerified {
		t.Fail()
	}
	if PermissionAtoI(PermissionTeamString) != PermissionTeam {
		t.Fail()
	}
	if PermissionAtoI(PermissionAdminString) != PermissionAdmin {
		t.Fail()
	}
}

func TestPermissionAtoI_AdminFallback(t *testing.T) {
	if PermissionAtoI("TRY_FALLBACK") != PermissionAdmin {
		t.Fail()
	}
}
