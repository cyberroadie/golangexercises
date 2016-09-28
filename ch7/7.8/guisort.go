package guisort

// ColumnData struct containing column title and data
// for simplicity we assume all columns have the same length
type ColumnData struct {
	Columns map[string][]string // column name with slice for data data
	Order   []string            // order of columns last clicked
	Length  int                 // length of column
}

func (cd *ColumnData) getTitles() (titles []string) {
	for v := range cd.Columns {
		titles = append(titles, v)
	}
	return
}

func (cd *ColumnData) Len() int {
	titles := cd.getTitles()
	rows := cd.Columns[titles[1]]
	return len(rows)
}

func (cd *ColumnData) Less(i, j int) bool {
	return false
}

func (cd *ColumnData) Swap(i, j int) {

}
