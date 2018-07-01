package markup

import "github.com/pkg/errors"

const (
	// FullSync indicates that sync should replace the full node.
	FullSync SyncScope = iota

	// AttrSync indicates that sync should replace only the attributes of the
	// node.
	AttrSync
)

// Sync is a struct which defines how a driver should handle a synchronisation
// of a node on the native side.
type Sync struct {
	Scope      SyncScope
	Index      int
	Node       *Node
	Attributes AttributeMap
}

// SyncScope defines the scope of a sync.
type SyncScope uint8

// Synchronize synchronize a whole component.
// Compares the newer state with the live state of the component.
func Synchronize(c Componer) (syncs []Sync, err error) {
	live := Root(c)

	r, err := render(c)
	if err != nil {
		err = errors.Errorf("unable to render %T: %v\n%v", c, err, c.Render())
		return
	}

	new, err := stringToNode(r)
	if err != nil {
		err = errors.Errorf("%T markup returned by Render() has a %v\n%v", c, err, r)
		return
	}

	if new.Type != HTMLNode {
		err = errors.Errorf("%T markup returned by Render() has a syntax error: root node is not a HTMLNode\n%v", c, r)
		return
	}

	syncs, _, err = syncNodes(live, new)
	return
}

func syncNodes(live *Node, new *Node) (syncs []Sync, parentShouldFullSync bool, err error) {
	if live.Type != new.Type {
		replaceNode(live, new)
		parentShouldFullSync = true
		return
	}

	switch live.Type {
	case TextNode:
		parentShouldFullSync = syncTextNodes(live, new)

	case ComponentNode:
		syncs, parentShouldFullSync, err = syncComponentNodes(live, new)

	case HTMLNode:
		syncs, parentShouldFullSync, err = syncHTMLNodes(live, new)
	}
	return
}

func syncTextNodes(live *Node, new *Node) (changed bool) {
	if live.Text == new.Text {
		return
	}

	live.Text = new.Text
	changed = true
	return
}

func syncComponentNodes(live *Node, new *Node) (syncs []Sync, parentShouldFullSync bool, err error) {
	if live.Tag != new.Tag {
		if err = replaceNode(live, new); err != nil {
			return
		}

		parentShouldFullSync = true
		return
	}

	attrDiff := live.Attributes.diff(new.Attributes)

	if len(attrDiff) == 0 {
		return
	}

	live.Attributes = new.Attributes
	decodeAttributeMap(new.Attributes, live.Component)
	syncs, err = Synchronize(live.Component)
	return
}

func syncHTMLNodes(live *Node, new *Node) (syncs []Sync, parentShouldFullSync bool, err error) {
	if live.Tag != new.Tag || len(live.Children) != len(new.Children) {
		if err = mergeHTMLNodes(live, new); err != nil {
			return
		}

		s := Sync{
			Scope: FullSync,
			Node:  live,
		}
		syncs = []Sync{s}
		return
	}

	shouldFullSync := false

	for i := 0; i < len(live.Children); i++ {
		childSyncs, requireFullSync, err := syncNodes(live.Children[i], new.Children[i])
		if err != nil {
			return nil, false, err
		}

		if requireFullSync && !shouldFullSync {
			shouldFullSync = true
		}

		if shouldFullSync {
			continue
		}

		syncs = append(syncs, childSyncs...)
	}

	if shouldFullSync {
		s := Sync{
			Scope: FullSync,
			Node:  live,
		}
		syncs = []Sync{s}
		return
	}

	if attrDiff := live.Attributes.diff(new.Attributes); len(attrDiff) != 0 {
		live.Attributes = new.Attributes
		s := Sync{
			Scope:      AttrSync,
			Node:       live,
			Attributes: attrDiff,
		}
		syncs = append([]Sync{s}, syncs...)
	}
	return
}

func replaceNode(live *Node, new *Node) error {
	if live.Type == ComponentNode {
		Dismount(live.Component)
	}

	live.Tag = new.Tag
	live.Type = new.Type
	live.Text = new.Text
	live.Attributes = new.Attributes
	live.Children = new.Children

	for _, c := range live.Children {
		c.Parent = live
	}
	return mountNode(live, live.Mount, live.ContextID)
}

func mergeHTMLNodes(live *Node, new *Node) error {
	live.Tag = new.Tag
	live.Attributes = new.Attributes

	for _, c := range live.Children {
		dismountNode(c)
	}

	live.Children = new.Children

	for _, c := range live.Children {
		c.Parent = live

		if err := mountNode(c, live.Mount, live.ContextID); err != nil {
			return err
		}
	}
	return nil
}
