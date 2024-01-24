package visuals

import (
	Month_Package "ExslReaderv2/Month"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateTemplate(months *[]Month_Package.Month) (templates []string) {
	for _, v := range *months {
		fmt.Printf("append template for %v", v)
		templates = append(templates, `<!DOCTYPE html>
		<html>
		
		<head>
			<title>2023</title>
		</head>
		
		<body>
			<h1>Monat {{.Name}}!</h1>
			<p>days: {{.Days}}</p>
			<p>Age: {{.AvgTemp}}</p>
		</body>
		
		</html>
		`)
	}
	return
}
func FillTemplates(months *[]Month_Package.Month, templates *[]string) {
	for i := 0; i < len(*templates)-1; i++ {
		templ, err := template.New("monthTemplate").Parse((*templates)[i])
		if err != nil {
			fmt.Println(err)
		}
		monthData := (*months)[i]
		outputPath := filepath.Join("HtmlMonths", monthData.Name)
		file, err := os.Create(outputPath)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		err = templ.Execute(file, monthData)
		if err != nil {
			fmt.Println(err)
		}
	}

}
