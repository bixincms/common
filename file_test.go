package common

import "testing"

func TestPathIsNoDirCreate(t *testing.T) {
	PathIsNoDirCreate("a/sad/as/asf/")
}
