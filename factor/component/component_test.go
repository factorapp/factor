package component

import (
	"bytes"
	"strings"
	"testing"
)

var goodTpl = `<nav>
<ul>
	<li><a class='{path === "/"  ? "selected" : ""}' href='.'>home</a></li>
	<li><a class='{path === "/about"  ? "selected" : ""}' href='about'>about</a></li>

	<!-- for the blog link, we're using rel=prefetch so that Factor prefetches
		 the blog data when we hover over the link or tap it on a touchscreen -->
	<li><a rel=prefetch class='{path.startsWith("/blog")  ? "selected" : ""}' href='blog'>blog</a></li>
</ul>
</nav>`
var goodStyle = `<style>
	nav {
		border-bottom: 1px solid rgba(170,30,30,0.1);
		font-weight: 300;
		padding: 0 1em;
	}

	ul {
		margin: 0;
		padding: 0;
	}

	/* clearfix */
	ul::after {
		content: '';
		display: block;
		clear: both;
	}

	li {
		display: block;
		float: left;
	}

	.selected {
		position: relative;
		display: inline-block;
	}

	.selected::after {
		position: absolute;
		content: '';
		width: calc(100% - 1em);
		height: 2px;
		background-color: rgb(170,30,30);
		display: block;
		bottom: -1px;
	}

	a {
		text-decoration: none;
		padding: 1em 0.5em;
		display: block;
	}
</style>`

func TestParse(t *testing.T) {
	c := parse(t)
	if c.Template != goodTpl {
		t.Errorf("Template Mismatch")
	}

	if c.Style != removeStyleTags(goodStyle) {
		t.Errorf("Style Mismatch")
	}
}
func TestQuoted(t *testing.T) {
	c := parseComponent(t)
	qs := c.QuotedStyle()
	if strings.Compare(qs[0:1], "`") != 0 {
		t.Errorf("expected style to start with backtick, got: %s", qs[0:1])
	}
	qt := c.QuotedTemplate()
	if strings.Compare(qt[len(qt)-1:len(qt)], "`") != 0 {
		t.Errorf("expected template to start with backtick, got: %s", qt[len(qt)-1:len(qt)])
	}
}
func parseComponent(t *testing.T) *Component {

	good := goodTpl + "\n" + goodStyle
	b := bytes.NewBuffer([]byte(good))

	c, e := Parse(b, "Good")
	if e != nil {
		t.Fatal("failed to parse template")
	}
	return c
}

func TestTransformUnparsed(t *testing.T) {
	c := &Component{}
	b := new(bytes.Buffer)
	err := c.Transform(b)
	if err != ErrComponentNotParsed {
		t.Error("Expected ErrComponentNotParsed on unparsed component")
	}
}

func TestTransform(t *testing.T) {
	c := parseComponent(t)
	c.Package = "mypackage"
	b := new(bytes.Buffer)
	err := c.Transform(b)
	if err != nil {
		t.Error(err)
	}
	// TODO: Compare against golden
}
