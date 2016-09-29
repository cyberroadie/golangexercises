package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/cyberroadie/golangexercises/ch7/7.9/tablesort"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	const tpl = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="UTF-8">
		
	</head>
    <style>
    table, th, td {
        border: 1px solid black;
    }
    </style>
	<body>
        <table>
        <form method="GET" action="">
        <tr>
        {{ range $key, $value := .Columns }}
        <td>
            
            <table>
                <thead>
                    <tr>
                        <td><button name="title" value="{{ $key }}" onClick="click(this);" >{{ $key }}</button></td>
                    </tr>
                </thead>
                <tbody>
                    {{ range $value }}
                    <tr>
                        <td>{{.}}</td>
                    </tr>
                    {{ end }}
                </tbody>
            </table>
            
        </td>
        {{ end }} 
        </tr>
        </form>
        </table>
	</body>
</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := tablesort.ColumnData{
		Columns: map[string][]string{
			"greet": {"hi", "hello", "hello", "haai"},
			"name":  {"horse", "sheep", "pig", "pigeon"},
		},
	}

	title := r.URL.Query().Get("title")
	if title != "" {
		data.ClickColumn(title)
		sort.Sort(&data)
	}
	err = t.Execute(w, data)
	check(err)

}
