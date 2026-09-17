package utility

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// XMLParamEngine provides XML parameter injection utilities.
type XMLParamEngine struct{}

// NewXMLParamEngine creates a new XMLParamEngine.
func NewXMLParamEngine() *XMLParamEngine {
	return &XMLParamEngine{}
}

// InjectAttribute injects an attribute into an XML element.
func (e *XMLParamEngine) InjectAttribute(xmlStr string, element string, attr string, value string) (string, error) {
	lines := strings.Split(xmlStr, "\n")
	result := make([]string, 0, len(lines))
	for _, line := range lines {
		result = append(result, line)
		if strings.Contains(line, "<"+element) && !strings.Contains(line, attr+"=") {
			idx := strings.Index(line, ">")
			if idx != -1 {
				line = line[:idx] + fmt.Sprintf(` %s="%s"`, attr, value) + line[idx:]
				result[len(result)-1] = line
			}
		}
	}
	return strings.Join(result, "\n"), nil
}

// InjectElement injects a child element into an XML structure.
func (e *XMLParamEngine) InjectElement(xmlStr string, parent string, child string, value string) (string, error) {
	closeTag := fmt.Sprintf("</%s>", parent)
	insertion := fmt.Sprintf("\n  <%s>%s</%s>", child, value, child)

	idx := strings.Index(xmlStr, closeTag)
	if idx == -1 {
		return "", fmt.Errorf("parent element %s not found", parent)
	}

	result := xmlStr[:idx] + insertion + xmlStr[idx:]
	return result, nil
}

// WrapElement wraps content in an XML element.
func (e *XMLParamEngine) WrapElement(content string, wrapper string) string {
	return fmt.Sprintf("<%s>%s</%s>", wrapper, content, wrapper)
}

// ParseAndExtract extracts values from XML by element name.
func (e *XMLParamEngine) ParseAndExtract(xmlStr string, element string) ([]string, error) {
	var values []string

	type generic struct {
		XMLName xml.Name
		Content string `xml:",chardata"`
	}

	decoder := xml.NewDecoder(strings.NewReader(xmlStr))
	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == element {
				var g generic
				_ = decoder.DecodeElement(&g, &se)
				values = append(values, g.Content)
			}
		}
	}

	return values, nil
}

// XXEInject simulates XML External Entity injection vector construction.
func (e *XMLParamEngine) XXEInject(doctype string, entity string, value string) string {
	return fmt.Sprintf(`<!DOCTYPE %s [<!ENTITY %s "%s">]>`, doctype, entity, value)
}

// BuildXML builds a simple XML document from key-value pairs.
func (e *XMLParamEngine) BuildXML(root string, pairs map[string]string) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "<%s>\n", root)
	for key, value := range pairs {
		fmt.Fprintf(&sb, "  <%s>%s</%s>\n", key, value, key)
	}
	fmt.Fprintf(&sb, "</%s>", root)
	return sb.String()
}
