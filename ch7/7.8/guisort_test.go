package guisort

import (
	"sort"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {

	cd := ColumnData{
		Columns: map[string][]string{
			"greet": {"hi", "hello", "hello", "haai"},
			"name":  {"horse", "sheep", "pig", "pigeon"},
		},
	}

	t.Logf("Length: %d\n", cd.Len())
	t.Logf("%v", cd)
	cd.ClickColumn("greet")
	sort.Sort(&cd)

	for k, v := range cd.Columns["greet"][1:] {
		if strings.Compare(cd.Columns["greet"][k], v) > 0 {
			t.Errorf("wrong order: %s > %s", cd.Columns["greet"][k], v)
		}
	}

	t.Logf("%v", cd)

	cd.ClickColumn("name")
	sort.Sort(&cd)
	for k, v := range cd.Columns["name"][1:] {
		if strings.Compare(cd.Columns["name"][k], v) > 0 {
			t.Errorf("wrong order: %s > %s", cd.Columns["name"][k], v)
		}
	}
	t.Logf("%v", cd)

}
