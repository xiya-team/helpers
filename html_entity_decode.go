package helpers

import "html"

func HtmlEntityDecode(str string)  string{

	return html.UnescapeString(str)
	
}

// HTMLEntityDecode html_entity_decode()
func HTMLEntityDecode(str string) string {
	return html.UnescapeString(str)
}