package task08

import (
	"reflect"
	"strings"
	"testing"
)

func TestRevenueByProduct(t *testing.T) {
	in := "product,quantity,unit_price\ncoffee,2,10.10\ncoffee,1,0.95\n\"tea, green\",3,2\n"
	got, err := RevenueByProduct(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}
	want := map[string]int64{"coffee": 2115, "tea, green": 600}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("RevenueByProduct = %#v, want %#v", got, want)
	}
}

func TestRevenueByProductRejectsBadRows(t *testing.T) {
	inputs := []string{
		"name,quantity,unit_price\na,1,1.00\n",
		"product,quantity,unit_price\n,1,1.00\n",
		"product,quantity,unit_price\na,0,1.00\n",
		"product,quantity,unit_price\na,1,1.001\n",
		"product,quantity,unit_price\na,nope,1.00\n",
	}
	for _, in := range inputs {
		if _, err := RevenueByProduct(strings.NewReader(in)); err == nil {
			t.Errorf("expected error for %q", in)
		}
	}
}
