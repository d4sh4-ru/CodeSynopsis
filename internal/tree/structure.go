package tree

import (
	"os"
	"strings"
)

type node struct {
	name     string
	children []*node
}

// Рекурсивное построение дерева
func BuildTreeStructure(paths []string) string {
	root := &node{name: ""}
	for _, path := range paths {
		components := strings.Split(path, string(os.PathSeparator))
		current := root
		for _, comp := range components {
			if comp == "" {
				continue
			}
			found := false
			for _, child := range current.children {
				if child.name == comp {
					current = child
					found = true
					break
				}
			}
			if !found {
				newNode := &node{name: comp}
				current.children = append(current.children, newNode)
				current = newNode
			}
		}
	}

	var sb strings.Builder
	var build func(n *node, prefix string, isLast bool)
	build = func(n *node, prefix string, isLast bool) {
		if n.name == "" {
			for i, child := range n.children {
				build(child, "", i == len(n.children)-1)
			}
			return
		}

		conn := "├── "
		if isLast {
			conn = "└── "
		}
		sb.WriteString(prefix + conn + n.name + "\n")

		newPrefix := prefix
		if isLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}

		for i, child := range n.children {
			build(child, newPrefix, i == len(n.children)-1)
		}
	}
	build(root, "", true)
	return sb.String()
}
