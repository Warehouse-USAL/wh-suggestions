package store

import "testing"

func TestKeys(t *testing.T) {
	if got := UserSKUs("u1"); got != "user:u1:skus" {
		t.Fatal(got)
	}
	if got := Cooccur("TEC-002"); got != "cooccur:TEC-002" {
		t.Fatal(got)
	}
	if HotGlobal != "hot:global" {
		t.Fatal(HotGlobal)
	}
}
