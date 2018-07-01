package markup

import (
	"testing"

	"github.com/satori/go.uuid"
)

type CompoMount struct {
	mounted             bool
	EmbedsNonRegistered bool
	EmbedsBadMarkup     bool
}

func (c *CompoMount) OnMount() {
	c.mounted = true
}

func (c *CompoMount) OnDismount() {
	c.mounted = false
}

func (c *CompoMount) Render() string {
	return `
<div>
    CompoMount is mounted
    <CompoEmpty />

    {{if .EmbedsNonRegistered}}
        <CompoNotRegistered />
    {{end}}

	{{if .EmbedsBadMarkup}}
        <CompoBadMarkup />
    {{end}}
</div>
    `
}

type CompoEmpty struct{}

func (c *CompoEmpty) Render() string {
	return `<p>CompoEmpty is mounted</p>`
}

type CompoNotRegistered struct{}

func (c *CompoNotRegistered) Render() string {
	return `<p>CompoNotRegistered</p>`
}

type CompoBadRenderTemplate struct{}

func (c *CompoBadRenderTemplate) Render() string {
	return `<p>CompoBadRender {{.Foo}}</p>`
}

type CompoBadMarkup struct{}

func (c *CompoBadMarkup) Render() string {
	return `<p>CompoBadMarkup</span>`
}

type CompoBadRoot struct{}

func (c *CompoBadRoot) Render() string {
	return `<CompoEmpty />`
}

type compoNotExported struct{}

func (c *compoNotExported) Render() string {
	return `<p>CompoNotExported</p>`
}

type CompoNoPtr struct{}

func (c CompoNoPtr) Render() string {
	return `<p>CompoNoPtr</p>`
}

func init() {
	Register(&CompoMount{})
	Register(&CompoEmpty{})
	Register(&CompoBadRenderTemplate{})
	Register(&CompoBadMarkup{})
	Register(&CompoBadRoot{})
}

func TestRegisterNotExported(t *testing.T) {
	defer func() { recover() }()
	Register(&compoNotExported{})
	t.Error("should panic")
}

func TestRegisterNoPtr(t *testing.T) {
	defer func() { recover() }()
	Register(CompoNoPtr{})
	t.Error("should panic")

}

func TestRootNotMounted(t *testing.T) {
	defer func() { recover() }()
	Root(&CompoEmpty{})
	t.Error("should panic")
}

func TestID(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoEmpty{}
	if _, err := Mount(c, ctx); err != nil {
		t.Fatal(err)
	}
	defer Dismount(c)

	t.Log(ID(c))
}

func TestComponent(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoEmpty{}
	if _, err := Mount(c, ctx); err != nil {
		t.Fatal(err)
	}
	defer Dismount(c)

	id := ID(c)
	if c2 := Component(id); c2 != c {
		t.Error("c and c2 should be the same component")
	}
}

func TestComponentPanic(t *testing.T) {
	defer func() { recover() }()
	Component(uuid.Must(uuid.NewV1()))
	t.Error("should panic")
}

func TestMount(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoMount{}

	root, err := Mount(c, ctx)
	if err != nil {
		t.Error(err)
	}

	t.Log(root)

	if l := len(components); l != 2 {
		t.Error("components len should be 2", l)
	}

	if l := len(nodes); l != 2 {
		t.Error("node len should be 2", l)
	}

	if !c.mounted {
		t.Error("c.mounted should be true:", c.mounted)
	}

	t.Log(Markup(c))

	Dismount(c)

	if l := len(components); l != 0 {
		t.Error("components len should be 0", l)
	}

	if l := len(nodes); l != 0 {
		t.Error("node len should be 0", l)
	}

	if c.mounted {
		t.Error("c.mounted should be false:", c.mounted)
	}

	Dismount(c)
}

func TestMountNotRegistered(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoNotRegistered{}

	if _, err := Mount(c, ctx); err == nil {
		t.Error("err should not be nil")
	}
}

func TestMountEmbedsNotRegistered(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoMount{EmbedsNonRegistered: true}

	if _, err := Mount(c, ctx); err == nil {
		t.Error("err should not be nil")
	}
}

func TestMountEmbedsBadMarkup(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoMount{EmbedsBadMarkup: true}

	if _, err := Mount(c, ctx); err == nil {
		t.Error("err should not be nil")
	}
}

func TestMountAlreadyMounted(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoMount{}

	if _, err := Mount(c, ctx); err != nil {
		t.Error(err)
	}

	if _, err := Mount(c, ctx); err == nil {
		t.Error("err should not be nil")
	}
}

func TestMountBadRenderTemplate(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoBadRenderTemplate{}
	Mount(c, ctx)
}

func TestMountBadMarkup(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoBadMarkup{}
	Mount(c, ctx)
}

func TestMountBadRoot(t *testing.T) {
	ctx := uuid.Must(uuid.NewV1())
	c := &CompoBadRoot{}
	Mount(c, ctx)
}
