package wikimedia

import "testing"
import "fmt"

func TestGetPageHTML(t *testing.T) {
	html, _ := GetPageHTML("cats")
	fmt.Print(html)
}
