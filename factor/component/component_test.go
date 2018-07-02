package component

import (
	"bytes"
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
	good := goodTpl + "\n" + goodStyle
	b := bytes.NewBuffer([]byte(good))

	c, err := Parse(b, "good")
	if c.Template != goodTpl {
		t.Errorf("Template Mismatch")
	}

	if c.Style != removeStyleTags(goodStyle) {
		t.Errorf("Style Mismatch")
	}
	if err != nil {
		t.Error(err)
	}
}
