package components

import (
	"github.com/bketelsen/factor/markup"
)

type Index struct {
}

var IndexTemplate = `<main><h1>
	Great success!</h1>

<figure>
	<img alt='Borat' src='/assets/great-success.png'></img>
	<figcaption>HIGH FIVE!</figcaption>
</figure>

<p>
	<strong>Try editing this file (routes/index.html) to test hot module reloading.</strong>
</p></main>`
var IndexStyles = `
	h1,
	figure,
	p {
		text-align: center;
		margin: 0 auto;
	}

	h1 {
		font-size: 2.8em;
		text-transform: uppercase;
		font-weight: 700;
		margin: 0 0 0.5em 0;
	}

	figure {
		margin: 0 0 1em 0;
	}

	img {
		width: 100%;
		max-width: 400px;
		margin: 0 0 1em 0;
	}

	p {
		margin: 1em auto;
	}

	@media (min-width: 480px) {
		h1 {
			font-size: 4em;
		}
	}
`

func (t *Index) Render() string {
	return IndexTemplate
}

func init() {
	markup.Register(&Index{})
}
