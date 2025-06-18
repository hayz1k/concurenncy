package tree

import "context"

// Node описывает вершину дерева со значением и дочерними узлами.
type Node struct {
	Value    int
	Children []*Node
}

// Walk обходит дерево в глубину и прерывается по отмене контекста.
func Walk(ctx context.Context, n *Node, visit func(int)) error {
	// TODO: реализовать рекурсивный обход с проверкой ctx.Done()
	return nil
}
