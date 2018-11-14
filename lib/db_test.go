package lib

import "testing"

func TestMigrationType_String(t *testing.T) {

	if Upgrade.String() != UpgradeString {
		t.Error("expected upgrade type")
	}

	if Downgrade.String() != DowngradeString {
		t.Error("expected downgrade type")
	}

}
