/**
 * @Author: entere
 * @Description:
 * @File:  xml
 * @Version: 1.0.0
 * @Date: 2020/3/19 21:22
 */

package wxutils

import (
	"bytes"
	"encoding/xml"
	"io"
	"strings"
)

//把xml字符转成map[string]string
func XMLToMap(xmlStr string) (m map[string]string, err error) {
	m = make(map[string]string)
	var (
		decoder = xml.NewDecoder(strings.NewReader(xmlStr))
		depth   = 0
		token   xml.Token
		key     string
		value   strings.Builder
	)
	for {
		token, err = decoder.Token()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		switch v := token.(type) {
		case xml.StartElement:
			depth++
			switch depth {
			case 2:
				key = v.Name.Local
				value.Reset()
			case 3:
				if err = decoder.Skip(); err != nil {
					return
				}
				depth--
				key = "" // key == "" indicates that the node with depth==2 has children
			}
		case xml.CharData:
			if depth == 2 && key != "" {
				value.Write(v)
			}
		case xml.EndElement:
			if depth == 2 && key != "" {
				m[key] = value.String()
			}
			depth--
		}
	}
}

func MapToXML(m map[string]string) (xmlStr string) {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range m {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)

	return buf.String()

}
