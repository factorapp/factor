//go:generate go run generate.go

// Package tags defines markup to create DOM elements.
//
// Generated from "HTML element reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under
// CC-BY-SA 2.5.
package tags

import (
	"github.com/bketelsen/factor/markup"
	"github.com/gobuffalo/packr"
)

var box packr.Box

func init() {
	box = packr.NewBox("./templates")
	markup.Register(&Heading1{})
}

// Anchor (or anchor element) creates a hyperlink to other web pages, files,
// locations within the same page, email addresses, or any other URL.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
type Anchor struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Anchor) Render() string {
	return box.String("a.html")
}

// The HTML Abbreviation element (<abbr>) represents an abbreviation or
// acronym; the optional title attribute can provide an expansion or
// description for the abbreviation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
type Abbreviation struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Abbreviation) Render() string {
	return box.String("abbr.html")
}

// Address indicates that the enclosed HTML provides contact information for a
// person or people, or for an organization.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
type Address struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Address) Render() string {
	return box.String("address.html")
}

// Area defines a hot-spot region on an image, and optionally associates it
// with a hypertext link. This element is used only within a <map> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
type Area struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Area) Render() string {
	return box.String("area.html")
}

// Article represents a self-contained composition in a document, page,
// application, or site, which is intended to be independently distributable or
// reusable (e.g., in syndication). Examples include: a forum post, a magazine
// or newspaper article, or a blog entry.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
type Article struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Article) Render() string {
	return box.String("article.html")
}

// Aside represents a portion of a document whose content is only indirectly
// related to the document's main content.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
type Aside struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Aside) Render() string {
	return box.String("aside.html")
}

// Audio is used to embed sound content in documents. It may contain one or
// more audio sources, represented using the src attribute or the <source>
// element: the browser will choose the most suitable one. It can also be the
// destination for streamed media, using a MediaStream.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
type Audio struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Audio) Render() string {
	return box.String("audio.html")
}

// The HTML Bring Attention To element (<b>) is used to draw the reader's
// attention to the element's contents, which are not otherwise granted special
// importance.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
type Bold struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Bold) Render() string {
	return box.String("b.html")
}

// Base specifies the base URL to use for all relative URLs contained within a
// document. There can be only one <base> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
type Base struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Base) Render() string {
	return box.String("base.html")
}

// The HTML BiDirectional Isolation element (<bdi>) is used to indicate spans
// of text which might need to be rendered in the opposite direction than the
// surrounding text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
type BidirectionalIsolation struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *BidirectionalIsolation) Render() string {
	return box.String("bdi.html")
}

// The HTML Bidirectional Text Override element (<bdo>) overrides the current
// directionality of text, so that the text within is rendered in a different
// direction.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
type BidirectionalOverride struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *BidirectionalOverride) Render() string {
	return box.String("bdo.html")
}

// BlockQuote (or HTML Block Quotation Element) indicates that the enclosed
// text is an extended quotation. Usually, this is rendered visually by
// indentation (see Notes for how to change it). A URL for the source of the
// quotation may be given using the cite attribute, while a text representation
// of the source can be given using the <cite> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
type BlockQuote struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *BlockQuote) Render() string {
	return box.String("blockquote.html")
}

// Body represents the content of an HTML document. There can be only one
// <body> element in a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/body
type Body struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Body) Render() string {
	return box.String("body.html")
}

// Break produces a line break in text (carriage-return). It is useful for
// writing a poem or an address, where the division of lines is significant.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
type Break struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Break) Render() string {
	return box.String("br.html")
}

// Button represents a clickable button, which can be used in forms, or
// anywhere in a document that needs simple, standard button functionality.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
type Button struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Button) Render() string {
	return box.String("button.html")
}

// Use the HTML <canvas> element with either the canvas scripting API or the
// WebGL API to draw graphics and animations.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
type Canvas struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Canvas) Render() string {
	return box.String("canvas.html")
}

// The HTML Table Caption element (<caption>) specifies the caption (or title)
// of a table, and if used is always the first child of a <table>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
type Caption struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Caption) Render() string {
	return box.String("caption.html")
}

// The HTML Citation element (<cite>) is used to describe a reference to a
// cited creative work, and must include either the title or the URL of that
// work.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
type Citation struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Citation) Render() string {
	return box.String("cite.html")
}

// Code displays its contents styled in a fashion intended to indicate that the
// text is a short fragment of computer code.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
type Code struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Code) Render() string {
	return box.String("code.html")
}

// Column defines a column within a table and is used for defining common
// semantics on all common cells. It is generally found within a <colgroup>
// element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
type Column struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Column) Render() string {
	return box.String("col.html")
}

// ColumnGroup defines a group of columns within a table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
type ColumnGroup struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *ColumnGroup) Render() string {
	return box.String("colgroup.html")
}

// Data links a given content with a machine-readable translation. If the
// content is time- or date-related, the <time> element must be used.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
type Data struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Data) Render() string {
	return box.String("data.html")
}

// DataList contains a set of <option> elements that represent the values
// available for other controls.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
type DataList struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *DataList) Render() string {
	return box.String("datalist.html")
}

// Description provides the details about or the definition of the preceding
// term (<dt>) in a description list (<dl>).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
type Description struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Description) Render() string {
	return box.String("dd.html")
}

// DeletedText represents a range of text that has been deleted from a
// document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
type DeletedText struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *DeletedText) Render() string {
	return box.String("del.html")
}

// The HTML Details Element (<details>) creates a disclosure widget in which
// information is visible only when the widget is toggled into an "open" state.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
type Details struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Details) Render() string {
	return box.String("details.html")
}

// The HTML Definition element (<dfn>) is used to indicate the term being
// defined within the context of a definition phrase or sentence.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
type Definition struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Definition) Render() string {
	return box.String("dfn.html")
}

// Dialog represents a dialog box or other interactive component, such as an
// inspector or window.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
type Dialog struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Dialog) Render() string {
	return box.String("dialog.html")
}

// The HTML Content Division element (<div>) is the generic container for flow
// content. It has no effect on the content or layout until styled using CSS.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
type Div struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Div) Render() string {
	return box.String("div.html")
}

// DescriptionList represents a description list. The element encloses a list
// of groups of terms (specified using the <dt> element) and descriptions
// (provided by <dd> elements). Common uses for this element are to implement a
// glossary or to display metadata (a list of key-value pairs).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
type DescriptionList struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *DescriptionList) Render() string {
	return box.String("dl.html")
}

// DefinitionTerm specifies a term in a description or definition list, and as
// such must be used inside a <dl> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
type DefinitionTerm struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *DefinitionTerm) Render() string {
	return box.String("dt.html")
}

// Emphasis marks text that has stress emphasis. The <em> element can be
// nested, with each level of nesting indicating a greater degree of emphasis.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
type Emphasis struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Emphasis) Render() string {
	return box.String("em.html")
}

// Embed embeds external content at the specified point in the document. This
// content is provided by an external application or other source of
// interactive content such as a browser plug-in.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
type Embed struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Embed) Render() string {
	return box.String("embed.html")
}

// FieldSet is used to group several controls as well as labels (<label>)
// within a web form.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
type FieldSet struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *FieldSet) Render() string {
	return box.String("fieldset.html")
}

// FigureCaption represents a caption or a legend associated with a figure or
// an illustration described by the rest of the data of the <figure> element
// which is its immediate ancestor.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
type FigureCaption struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *FigureCaption) Render() string {
	return box.String("figcaption.html")
}

// Figure represents self-contained content, frequently with a caption
// (<figcaption>), and is typically referenced as a single unit.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
type Figure struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Figure) Render() string {
	return box.String("figure.html")
}

// Footer represents a footer for its nearest sectioning content or sectioning
// root element. A footer typically contains information about the author of
// the section, copyright data or links to related documents.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
type Footer struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Footer) Render() string {
	return box.String("footer.html")
}

// Form represents a document section that contains interactive controls for
// submitting information to a web server.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
type Form struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Form) Render() string {
	return box.String("form.html")
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
type Heading1 struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Heading1) Render() string {
	return box.String("h1.html")
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
type Heading2 struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Heading2) Render() string {
	return box.String("h2.html")
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
type Heading3 struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Heading3) Render() string {
	return box.String("h3.html")
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
type Heading4 struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Heading4) Render() string {
	return box.String("h4.html")
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
type Heading5 struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Heading5) Render() string {
	return box.String("h5.html")
}

// The HTML <h1>–<h6> elements represent six levels of section headings. <h1>
// is the highest section level and <h6> is the lowest.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
type Heading6 struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Heading6) Render() string {
	return box.String("h6.html")
}

// Header represents introductory content, typically a group of introductory or
// navigational aids. It may contain some heading elements but also other
// elements like a logo, a search form, an author name, and so on.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
type Header struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Header) Render() string {
	return box.String("header.html")
}

// HeadingsGroup represents a multi-level heading for a section of a document.
// It groups a set of <h1>–<h6> elements.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
type HeadingsGroup struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *HeadingsGroup) Render() string {
	return box.String("hgroup.html")
}

// HorizontalRule represents a thematic break between paragraph-level elements
// (for example, a change of scene in a story, or a shift of topic with a
// section); historically, this has been presented as a horizontal rule or
// line.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
type HorizontalRule struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *HorizontalRule) Render() string {
	return box.String("hr.html")
}

// Italic represents a range of text that is set off from the normal text for
// some reason. Some examples include technical terms, foreign language
// phrases, or fictional character thoughts. It is typically displayed in
// italic type.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
type Italic struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Italic) Render() string {
	return box.String("i.html")
}

// The HTML Inline Frame element (<iframe>) represents a nested browsing
// context, effectively embedding another HTML page into the current page.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
type InlineFrame struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *InlineFrame) Render() string {
	return box.String("iframe.html")
}

// Image embeds an image into the document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
type Image struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Image) Render() string {
	return box.String("img.html")
}

// Input is used to create interactive controls for web-based forms in order to
// accept data from the user.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
type Input struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Input) Render() string {
	return box.String("input.html")
}

// InsertedText represents a range of text that has been added to a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
type InsertedText struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *InsertedText) Render() string {
	return box.String("ins.html")
}

// The HTML Keyboard Input element (<kbd>) represents a span of inline text
// denoting textual user input from a keyboard, voice input, or any other text
// entry device.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
type KeyboardInput struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *KeyboardInput) Render() string {
	return box.String("kbd.html")
}

// Label represents a caption for an item in a user interface.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
type Label struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Label) Render() string {
	return box.String("label.html")
}

// Legend represents a caption for the content of its parent <fieldset>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
type Legend struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Legend) Render() string {
	return box.String("legend.html")
}

// ListItem is used to represent an item in a list. It must be contained in a
// parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu
// (<menu>). In menus and unordered lists, list items are usually displayed
// using bullet points. In ordered lists, they are usually displayed with an
// ascending counter on the left, such as a number or letter.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
type ListItem struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *ListItem) Render() string {
	return box.String("li.html")
}

// Link specifies relationships between the current document and an external
// resource. This element is most commonly used to link to stylesheets.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
type Link struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Link) Render() string {
	return box.String("link.html")
}

// Main represents the dominant content of the <body> of a document, portion of
// a document or application. The main content area consists of content that is
// directly related to or expands upon the central topic of a document, or the
// central functionality of an application.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
type Main struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Main) Render() string {
	return box.String("main.html")
}

// Map is used with <area> elements to define an image map (a clickable link
// area).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
type Map struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Map) Render() string {
	return box.String("map.html")
}

// The HTML Mark Text element (<mark>) represents text which is marked or
// highlighted for reference or notation purposes, due to the marked passage's
// relevance or importance in the enclosing context.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
type Mark struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Mark) Render() string {
	return box.String("mark.html")
}

// Menu represents a group of commands that a user can perform or activate.
// This includes both list menus, which might appear across the top of a
// screen, as well as context menus, such as those that might appear underneath
// a button after it has been clicked.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
type Menu struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Menu) Render() string {
	return box.String("menu.html")
}

// Meta represents metadata that cannot be represented by other HTML
// meta-related elements, like <base>, <link>, <script>, <style> or <title>.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
type Meta struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Meta) Render() string {
	return box.String("meta.html")
}

// Meter represents either a scalar value within a known range or a fractional
// value.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
type Meter struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Meter) Render() string {
	return box.String("meter.html")
}

// Navigation represents a section of a page whose purpose is to provide
// navigation links, either within the current document or to other documents.
// Common examples of navigation sections are menus, tables of contents, and
// indexes.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
type Navigation struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Navigation) Render() string {
	return box.String("nav.html")
}

// NoScript defines a section of HTML to be inserted if a script type on the
// page is unsupported or if scripting is currently turned off in the browser.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
type NoScript struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *NoScript) Render() string {
	return box.String("noscript.html")
}

// Object represents an external resource, which can be treated as an image, a
// nested browsing context, or a resource to be handled by a plugin.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
type Object struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Object) Render() string {
	return box.String("object.html")
}

// OrderedList represents an ordered list of items, typically rendered as a
// numbered list.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
type OrderedList struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *OrderedList) Render() string {
	return box.String("ol.html")
}

// OptionsGroup creates a grouping of options within a <select> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
type OptionsGroup struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *OptionsGroup) Render() string {
	return box.String("optgroup.html")
}

// Option is used to define an item contained in a <select>, an <optgroup>, or
// a <datalist> element. As such, <option> can represent menu items in popups
// and other lists of items in an HTML document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
type Option struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Option) Render() string {
	return box.String("option.html")
}

// The HTML Output element (<output>) is a container element into which a site
// or app can inject the results of a calculation or the outcome of a user
// action.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
type Output struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Output) Render() string {
	return box.String("output.html")
}

// Paragraph represents a paragraph of text.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
type Paragraph struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Paragraph) Render() string {
	return box.String("p.html")
}

// Parameter defines parameters for an <object> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
type Parameter struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Parameter) Render() string {
	return box.String("param.html")
}

// Picture serves as a container for zero or more <source> elements and one
// <img> element to provide versions of an image for different display device
// scenarios.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
type Picture struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Picture) Render() string {
	return box.String("picture.html")
}

// Preformatted represents preformatted text which is to be presented exactly
// as written in the HTML file.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
type Preformatted struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Preformatted) Render() string {
	return box.String("pre.html")
}

// Progress displays an indicator showing the completion progress of a task,
// typically displayed as a progress bar.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
type Progress struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Progress) Render() string {
	return box.String("progress.html")
}

// Quote indicates that the enclosed text is a short inline quotation. Most
// modern browsers implement this by surrounding the text in quotation marks.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
type Quote struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Quote) Render() string {
	return box.String("q.html")
}

// The HTML Ruby Fallback Parenthesis (<rp>) element is used to provide
// fall-back parentheses for browsers that do not support display of ruby
// annotations using the <ruby> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
type RubyParenthesis struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *RubyParenthesis) Render() string {
	return box.String("rp.html")
}

// The HTML Ruby Text (<rt>) element specifies the ruby text component of a
// ruby annotation, which is used to provide pronunciation, translation, or
// transliteration information for East Asian typography. The <rt> element must
// always be contained within a <ruby> element.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
type RubyText struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *RubyText) Render() string {
	return box.String("rt.html")
}

// The HTML Ruby Text Container (<rtc>) element embraces semantic annotations
// of characters presented in a ruby of <rb> elements used inside of <ruby>
// element. <rb> elements can have both pronunciation (<rt>) and semantic
// (<rtc>) annotations.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
type RubyTextContainer struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *RubyTextContainer) Render() string {
	return box.String("rtc.html")
}

// Ruby represents a ruby annotation. Ruby annotations are for showing
// pronunciation of East Asian characters.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
type Ruby struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Ruby) Render() string {
	return box.String("ruby.html")
}

// Strikethrough renders text with a strikethrough, or a line through it. Use
// the <s> element to represent things that are no longer relevant or no longer
// accurate. However, <s> is not appropriate when indicating document edits;
// for that, use the <del> and <ins> elements, as appropriate.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
type Strikethrough struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Strikethrough) Render() string {
	return box.String("s.html")
}

// The HTML Sample Element (<samp>) is used to enclose inline text which
// represents sample (or quoted) output from a computer program.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
type Sample struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Sample) Render() string {
	return box.String("samp.html")
}

// Script is used to embed or reference executable code; this is typically used
// to embed or refer to JavaScript code.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
type Script struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Script) Render() string {
	return box.String("script.html")
}

// Section represents a standalone section — which doesn't have a more
// specific semantic element to represent it — contained within an HTML
// document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
type Section struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Section) Render() string {
	return box.String("section.html")
}

// Select represents a control that provides a menu of options:
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
type Select struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Select) Render() string {
	return box.String("select.html")
}

// Slot—part of the Web Components technology suite—is a placeholder inside
// a web component that you can fill with your own markup, which lets you
// create separate DOM trees and present them together.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/slot
type Slot struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Slot) Render() string {
	return box.String("slot.html")
}

// Small makes the text font size one size smaller (for example, from large to
// medium, or from small to x-small) down to the browser's minimum font size.
// In HTML5, this element is repurposed to represent side-comments and small
// print, including copyright and legal text, independent of its styled
// presentation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
type Small struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Small) Render() string {
	return box.String("small.html")
}

// Source specifies multiple media resources for the <picture>, the <audio>
// element, or the <video> element. It is an empty element. It is commonly used
// to serve the same media content in multiple formats supported by different
// browsers.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
type Source struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Source) Render() string {
	return box.String("source.html")
}

// Span is a generic inline container for phrasing content, which does not
// inherently represent anything. It can be used to group elements for styling
// purposes (using the class or id attributes), or because they share attribute
// values, such as lang.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
type Span struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Span) Render() string {
	return box.String("span.html")
}

// The HTML Strong Importance Element (<strong>) indicates that its contents
// have strong importance, seriousness, or urgency. Browsers typically render
// the contents in bold type.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
type Strong struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Strong) Render() string {
	return box.String("strong.html")
}

// Style contains style information for a document, or part of a document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
type Style struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Style) Render() string {
	return box.String("style.html")
}

// The HTML Subscript element (<sub>) specifies inline text which should be
// displayed as subscript for solely typographical reasons.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
type Subscript struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Subscript) Render() string {
	return box.String("sub.html")
}

// The HTML Disclosure Summary element (<summary>) element specifies a summary,
// caption, or legend for a <details> element's disclosure box.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
type Summary struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Summary) Render() string {
	return box.String("summary.html")
}

// The HTML Superscript element (<sup>) specifies inline text which is to be
// displayed as superscript for solely typographical reasons.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
type Superscript struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Superscript) Render() string {
	return box.String("sup.html")
}

// Table represents tabular data — that is, information presented in a
// two-dimensional table comprised of rows and columns of cells containing
// data.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
type Table struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Table) Render() string {
	return box.String("table.html")
}

// The HTML Table Body element (<tbody>) encapsulates a set of table row (<tr>
// elements, indicating that they comprise the body of the table (<table>).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
type TableBody struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *TableBody) Render() string {
	return box.String("tbody.html")
}

// TableData defines a cell of a table that contains data. It participates in
// the table model.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
type TableData struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *TableData) Render() string {
	return box.String("td.html")
}

// The HTML Content Template (<template>) element is a mechanism for holding
// client-side content that is not to be rendered when a page is loaded but may
// subsequently be instantiated during runtime using JavaScript.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
type Template struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Template) Render() string {
	return box.String("template.html")
}

// TextArea represents a multi-line plain-text editing control, useful when you
// want to allow users to enter a sizeable amount of free-form text, for
// example a comment on a review or feedback form.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
type TextArea struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *TextArea) Render() string {
	return box.String("textarea.html")
}

// TableFoot defines a set of rows summarizing the columns of the table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
type TableFoot struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *TableFoot) Render() string {
	return box.String("tfoot.html")
}

// TableHeader defines a cell as header of a group of table cells. The exact
// nature of this group is defined by the scope and headers attributes.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
type TableHeader struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *TableHeader) Render() string {
	return box.String("th.html")
}

// TableHead defines a set of rows defining the head of the columns of the
// table.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
type TableHead struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *TableHead) Render() string {
	return box.String("thead.html")
}

// Time represents a specific period in time. It may include the datetime
// attribute to translate dates into machine-readable format, allowing for
// better search engine results or custom features such as reminders.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
type Time struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Time) Render() string {
	return box.String("time.html")
}

// The HTML Title element (<title>) defines the document's title that is shown
// in a browser's title bar or a page's tab.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
type Title struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Title) Render() string {
	return box.String("title.html")
}

// TableRow defines a row of cells in a table. The row's cells can then be
// established using a mix of <td> (data cell) and <th> (header cell)
// elements.The HTML <tr> element specifies that the markup contained inside
// the <tr> block comprises one row of a table, inside which the <th> and <td>
// elements create header and data cells, respectively, within the row.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
type TableRow struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *TableRow) Render() string {
	return box.String("tr.html")
}

// Track is used as a child of the media elements <audio> and <video>. It lets
// you specify timed text tracks (or time-based data), for example to
// automatically handle subtitles. The tracks are formatted in WebVTT format
// (.vtt files) — Web Video Text Tracks or Timed Text Markup Language (TTML).
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
type Track struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Track) Render() string {
	return box.String("track.html")
}

// The HTML Unarticulated Annotation element (<u>) represents a span of inline
// text which should be rendered in a way that indicates that it has a
// non-textual annotation.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
type Underline struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Underline) Render() string {
	return box.String("u.html")
}

// UnorderedList represents an unordered list of items, typically rendered as a
// bulleted list.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
type UnorderedList struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *UnorderedList) Render() string {
	return box.String("ul.html")
}

// The HTML Variable element (<var>) represents the name of a variable in a
// mathematical expression or a programming context.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
type Variable struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Variable) Render() string {
	return box.String("var.html")
}

// The HTML Video element (<video>) embeds a media player which supports video
// playback into the document.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
type Video struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *Video) Render() string {
	return box.String("video.html")
}

// WordBreakOpportunity represents a word break opportunity—a position within
// text where the browser may optionally break a line, though its line-breaking
// rules would not otherwise create a break at that location.
//
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
type WordBreakOpportunity struct {
	Tag        string
	Text       string
	Attributes markup.AttributeMap
}

func (t *WordBreakOpportunity) Render() string {
	return box.String("wbr.html")
}
