package markup

import (
	"testing"

	"github.com/satori/go.uuid"
)

type CompoSync struct {
	TextChange         bool
	HTMLAttrChange     bool
	HTMLTagChange      bool
	HTMLTagChangeError bool
	CompoAttrChange    bool
	CompoChangeError   bool
	CompoChange        bool
	TypeChange         bool
	AddRemove          bool
}

func (c *CompoSync) Render() string {
	return `
<div>
    <!-- TextChange -->
    <p>{{if .TextChange}}Maxoo{{else}}Jonhy{{end}}</p>

    <!-- HTMLAttrChange -->
    <p class="{{if .HTMLAttrChange}}boo{{end}}">Say something</p>

    <!-- HTMLTagChange -->
    {{if .HTMLTagChange}}
        <h1>Hello</h1>
    {{else}}
        <h2>Hello</h2>
    {{end}}

	 <!-- HTMLTagChangeError -->
    {{if .HTMLTagChangeError}}
        <h1>
			<CompoSyncError BadTemplate="true" />
		</h1>
    {{else}}
        <h2>Hello</h2>
    {{end}}

    <!-- CompoAttrChange -->
    <SubCompoSync Name="{{if .CompoAttrChange}}Max{{else}}Maxence{{end}}" />

    <!-- CompoChange -->
    <div>
        {{if .CompoChange}}
            <SubCompoSyncBis />          
        {{else}}
            <SubCompoSync Name="Jonhzy" />
        {{end}}
    </div>

	<!-- CompoChangeError -->
    <div>
        {{if .CompoChangeError}}
            <CompoSyncError BadTemplate="true" />          
        {{else}}
            <SubCompoSync Name="poaa" />
        {{end}}
    </div>

     <!-- TypeChange -->
    <div>
        {{if .TypeChange}}
            <div>
                <h1>I'm changed</h1>
                <SubCompoSyncBis /> 
            </div>           
        {{else}}
            <SubCompoSync Name="Bravo" />
        {{end}}
    </div>


    <!-- AddRemove -->
    <div>
        {{if .AddRemove}}<h1>Plop!</h1>{{end}}
    </div>
</div>
    `
}

type SubCompoSync struct {
	Name string
}

func (c *SubCompoSync) Render() string {
	return `
<div>
    <h1>{{html .Name}}</h1>
    <p>Whoa</p>
</div>
    `
}

type SubCompoSyncBis struct {
}

func (c *SubCompoSyncBis) Render() string {
	return `<p>I'm sexy</p>`
}

type CompoSyncError struct {
	BadTemplate bool
	BadRoot     bool
	BadMarkup   bool
}

func (c *CompoSyncError) Render() string {
	return `
{{if .BadTemplate}}
    {{.Unknown}}
{{else if .BadRoot}}
    <CompoBadRoot />
{{else if .BadMarkup}}
    <div></p>
{{else}}
    <div>Murloks!!!</div>
{{end}}
    `
}

func init() {
	Register(&CompoSync{})
	Register(&SubCompoSync{})
	Register(&SubCompoSyncBis{})
	Register(&CompoSyncError{})
}

func TestSynchronizeTextChange(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.TextChange = true

	syncs, err := Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s := syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != FullSync {
		t.Error("s.Scope should be FullSync")
	}
}

func TestSynchronizeHTMLAttrChange(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.HTMLAttrChange = true

	syncs, err := Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s := syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != AttrSync {
		t.Error("s.Scope should be AttrSync")
	}

	if s.Attributes["class"] != "boo" {
		t.Error(`s.Attributes["class"] should be boo:`, s.Attributes["class"])
	}
}

func TestSynchronizeHTMLTagChange(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.HTMLTagChange = true

	syncs, err := Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s := syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != FullSync {
		t.Error("s.Scope should be FullSync")
	}
}

func TestSynchronizeHTMLTagChangeError(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.HTMLTagChangeError = true

	if _, err := Synchronize(c); err == nil {
		t.Error("error should not be nil")
	}
}

func TestSynchronizeCompoAttrChange(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.CompoAttrChange = true

	syncs, err := Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s := syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != FullSync {
		t.Error("s.Scope should be FullSync")
	}
}

func TestSynchronizeCompoChange(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.CompoChange = true

	syncs, err := Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s := syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != FullSync {
		t.Error("s.Scope should be FullSync")
	}
}

func TestSynchronizeCompoChangeError(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.CompoChangeError = true

	if _, err := Synchronize(c); err == nil {
		t.Error("err should not be nil")
	}
}

func TestSynchronizeTypeChange(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.TypeChange = true

	syncs, err := Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s := syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != FullSync {
		t.Error("s.Scope should be FullSync")
	}
}

func TestAddRemove(t *testing.T) {
	c := &CompoSync{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	// Add.
	c.AddRemove = true

	syncs, err := Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s := syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != FullSync {
		t.Error("s.Scope should be FullSync")
	}

	// Remove.
	c.AddRemove = false

	syncs, err = Synchronize(c)
	if err != nil {
		t.Fatal(err)
	}

	if l := len(syncs); l != 1 {
		t.Error("l should be 1:", l)
	}

	s = syncs[0]
	t.Log(s.Node.Markup())

	if s.Scope != FullSync {
		t.Error("s.Scope should be FullSync")
	}
}

func TestSynchronizeBadTemplate(t *testing.T) {
	c := &CompoSyncError{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.BadTemplate = true
	Synchronize(c)
}

func TestSynchronizeBadRoot(t *testing.T) {
	c := &CompoSyncError{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.BadRoot = true
	Synchronize(c)
}

func TestSynchronizeBadMarkup(t *testing.T) {
	c := &CompoSyncError{}
	ctx := uuid.Must(uuid.NewV1())

	Mount(c, ctx)
	defer Dismount(c)

	c.BadMarkup = true
	Synchronize(c)
}
