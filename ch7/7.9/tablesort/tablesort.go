package tablesort

import "strings"

// ColumnData struct containing column title and data
// for simplicity we assume all columns have the same length
type ColumnData struct {
	Columns map[string][]string // column name with slice for data data
	order   []string            // order of columns last clicked
}

func (cd *ColumnData) setOrder() {

	for k := range cd.Columns {
		cd.order = append(cd.order, k)
	}

}

// ClickColumn sort by column with given titel
func (cd *ColumnData) ClickColumn(title string) {
	if len(cd.order) == 0 {
		cd.setOrder()
	}

	if title == cd.order[0] {
		return
	}

	for i := range cd.order {
		if title == cd.order[i] {
			cd.order[0], cd.order[i] = cd.order[i], cd.order[0]
			break
		}
	}
}

func (cd *ColumnData) getTitles() (titles []string) {
	for v := range cd.Columns {
		titles = append(titles, v)
	}
	return
}

func (cd *ColumnData) Len() int {
	if len(cd.order) == 0 {
		cd.setOrder()
	}
	titles := cd.getTitles()
	rows := cd.Columns[titles[0]]
	return len(rows)
}

func (cd *ColumnData) Less(i, j int) bool {
	switch strings.Compare(cd.Columns[cd.order[0]][i], cd.Columns[cd.order[0]][j]) {
	case 0:
		for _, x := range cd.order[1:] {
			switch strings.Compare(cd.Columns[x][i], cd.Columns[x][j]) {
			case 0:
				continue
			case -1:
				for _, x2 := range cd.order {
					cd.Columns[x2][i], cd.Columns[x2][j] = cd.Columns[x2][j], cd.Columns[x2][i]
				}
				break
			case 1:
				break
			}
		}
		return false
	case -1:
		return true
	default:
		return false
	}
}

func (cd *ColumnData) Swap(i, j int) {
	for _, t := range cd.order {
		cd.Columns[t][i], cd.Columns[t][j] = cd.Columns[t][j], cd.Columns[t][i]
	}
}
