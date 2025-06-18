package tree

import "context"

type Node struct {
	Value    int
	Children []*Node
}

func Walk(ctx context.Context, n *Node, visit func(int)) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}
	if n == nil {
		return nil
	}
	visit(n.Value)
	for _, child := range n.Children {
		if err := Walk(ctx, child, visit); err != nil {
			return err
		}
	}
	return nil
}
