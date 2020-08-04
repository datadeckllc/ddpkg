package main

import (
"github.com/PuerkitoBio/goquery"
"github.com/360EntSecGroup-Skylar/excelize"
"net/http"
"net/url"
"encoding/csv"
"fmt"
"os"
"strings"
)

func check (e error) {
if e != nil {
panic(e)
}
}
func isAlpha (c int) bool {
return (97 <= c && c <= 122) || (45 <= c && c <= 90)
}

func main () {
/* open file */
args := os.Args[1]
file, err := os.Open(args)
check(err)

/* read CSV */
csvfile := csv.NewReader(file)
records, err := csvfile.ReadAll()
rec_count := len(records)
fmt.Println(records)
records = records[1:rec_count]
check(err)

/* create spreadsheet */
xlsx := excelize.NewFile()
xlsx.NewSheet("SOSCA")
xlsx.SetCellValue("SOSCA", "A1", "Entity #")
xlsx.SetCellValue("SOSCA", "B1", "Registration Date")
xlsx.SetCellValue("SOSCA", "C1", "Status")
xlsx.SetCellValue("SOSCA", "D1", "Entity Name")
xlsx.SetCellValue("SOSCA", "E1", "Jurisdiction")
xlsx.SetCellValue("SOSCA", "F1", "Agent")
xlsx.SetCellValue("SOSCA", "G1", "# Of Entries")
xlsx.SetCellValue("SOSCA", "H1", "Link")
defer xlsx.SaveAs("./SOSCA.xlsx")

/* iterate over records */
var row int = 2
rec_count -= 8
for i := 0; i < rec_count; i++ {
/* fetch page */
url := fmt.Sprintf("https://businesssearch.sos.ca.gov/CBS/SearchResults?filing=&SearchType=LPLLC&SearchCriteria=%s&SearchSubType=Keyword",
url.QueryEscape(records[i][0]))

fmt.Printf("Loading %s\n", url)
res, err := http.Get(url)
check(err)
defer res.Body.Close()

doc, err := goquery.NewDocumentFromReader(res.Body)
check(err)

qs := doc.Find("#enitityTable").Find("td")
if len(qs.Nodes) != 0 {
qs.Each( func(i int, s *goquery.Selection){

xcell := fmt.Sprintf("%c%d", 65 + (i % 6), row)

if i % 6 == 3 {
off := 0
str := s.Text()
nl := strings.IndexRune(str[1:], '\n')
for isAlpha(int(str[off + nl])) == false { off++ }
str = str[nl + off + 1:]
xlsx.SetCellValue("SOSCA", xcell, str)
} else {

xlsx.SetCellValue("SOSCA", xcell, s.Text())

}
})
} else {

xlsx.SetCellValue("SOSCA", fmt.Sprintf("D%d", row), records[i][0])

}

xlsx.SetCellValue("SOSCA", fmt.Sprintf("H%d", row), url)
row++

}
}

	
	
	

