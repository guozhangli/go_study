package reader

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"time"
)

type CommonInformationItem struct {
	/**
	 * Title of the item
	 */
	title string

	/**
	 * Date of the item in text format
	 */
	txtDate string

	/**
	 * Date of the item as a Date object
	 */
	date time.Time

	/**
	 * Link of the item
	 */
	link string

	/**
	 * Description of the item
	 */
	description string

	/**
	 * Id of the item
	 */
	id string

	/**
	 * Source (the rss name) of the item
	 */
	source string
	/**
	作者
	*/
	author string
}

func (item *CommonInformationItem) toXML() string {
	var buf bytes.Buffer
	buf.WriteString("<item>\n")
	buf.WriteString("<id>\n")
	buf.WriteString(item.id)
	buf.WriteString("\n</id>\n")
	buf.WriteString("\n<title>\n")
	buf.WriteString(item.title)
	buf.WriteString("\n</title>\n")
	buf.WriteString("\n<date>\n")
	buf.WriteString(item.txtDate)
	buf.WriteString("\n</date>\n")
	buf.WriteString("\n<link>\n")
	buf.WriteString(item.link)
	buf.WriteString("\n</link>\n")
	buf.WriteString("\n<description>\n")
	buf.WriteString(item.description)
	buf.WriteString("\n</description>\n")
	buf.WriteString("\n</item>\n")
	return buf.String()
}

func (item *CommonInformationItem) getFileName() string {
	var buf bytes.Buffer
	buf.WriteString(item.source)
	buf.WriteString("_")
	hashCode := md5.Sum([]byte(item.description))
	buf.WriteString(fmt.Sprintf("%x", hashCode))
	buf.WriteString(".xml")
	return buf.String()
}
