
package components

import (
	
	"github.com/bketelsen/factor/markup"
)



var NavTemplate =  `<nav>
    <ul>
        <li>This is an item</li>
        <li>This is another item</li>
    </ul>
</nav>` 
var NavStyles = `
    nav {
        border-bottom: 1px solid rgba(170, 30, 30, 0.1);
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
        background-color: rgb(170, 30, 30);
        display: block;
        bottom: -1px;
    }
    
    a {
        text-decoration: none;
        padding: 1em 0.5em;
        display: block;
    }
`


func (t *Nav) Render() string {
	return NavTemplate
}

func init() {
	markup.Register(&Nav{})
}
