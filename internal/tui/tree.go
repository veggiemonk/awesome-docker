package tui

import (
	"sort"
	"strings"

	"github.com/veggiemonk/awesome-docker/internal/cache"
)

// TreeNode represents a node in the category tree.
type TreeNode struct {
	Name     string       // display name (leaf segment, e.g. "Networking")
	Path     string       // full path (e.g. "Container Operations > Networking")
	Children []*TreeNode
	Expanded bool
	Entries  []cache.HealthEntry
}

// FlatNode is a visible tree node with its indentation depth.
type FlatNode struct {
	Node  *TreeNode
	Depth int
}

// HasChildren returns true if this node has child categories.
func (n *TreeNode) HasChildren() bool {
	return len(n.Children) > 0
}

// TotalEntries returns the count of entries in this node and all descendants.
func (n *TreeNode) TotalEntries() int {
	count := len(n.Entries)
	for _, c := range n.Children {
		count += c.TotalEntries()
	}
	return count
}

// AllEntries returns entries from this node and all descendants.
func (n *TreeNode) AllEntries() []cache.HealthEntry {
	result := make([]cache.HealthEntry, 0, n.TotalEntries())
	result = append(result, n.Entries...)
	for _, c := range n.Children {
		result = append(result, c.AllEntries()...)
	}
	return result
}

// BuildTree constructs a tree from flat HealthEntry slice, grouping by Category.
func BuildTree(entries []cache.HealthEntry) []*TreeNode {
	root := &TreeNode{Name: "root"}
	nodeMap := map[string]*TreeNode{}

	for _, e := range entries {
		cat := e.Category
		if cat == "" {
			cat = "Uncategorized"
		}

		node := ensureNode(root, nodeMap, cat)
		node.Entries = append(node.Entries, e)
	}

	// Sort children at every level
	sortTree(root)
	return root.Children
}

func ensureNode(root *TreeNode, nodeMap map[string]*TreeNode, path string) *TreeNode {
	if n, ok := nodeMap[path]; ok {
		return n
	}

	parts := strings.Split(path, " > ")
	current := root
	for i, part := range parts {
		subpath := strings.Join(parts[:i+1], " > ")
		if n, ok := nodeMap[subpath]; ok {
			current = n
			continue
		}
		child := &TreeNode{
			Name: part,
			Path: subpath,
		}
		current.Children = append(current.Children, child)
		nodeMap[subpath] = child
		current = child
	}
	return current
}

func sortTree(node *TreeNode) {
	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].Name < node.Children[j].Name
	})
	for _, c := range node.Children {
		sortTree(c)
	}
}

// FlattenVisible returns visible nodes in depth-first order for rendering.
func FlattenVisible(roots []*TreeNode) []FlatNode {
	var result []FlatNode
	for _, r := range roots {
		flattenNode(r, 0, &result)
	}
	return result
}

func flattenNode(node *TreeNode, depth int, result *[]FlatNode) {
	*result = append(*result, FlatNode{Node: node, Depth: depth})
	if node.Expanded {
		for _, c := range node.Children {
			flattenNode(c, depth+1, result)
		}
	}
}
