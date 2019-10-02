package composite

// Client client that need to process the leaf and the non leaf node
type Client struct{}

// Run process the tree
// Tree:
//      1
//     / \
//    2   3
//   / \
//  4   5
func (Client) Run() {
	four := NewLeaf(4)
	five := NewLeaf(5)
	two := NewComposite(2)
	two.Add(four, five)
	three := NewLeaf(3)
	one := NewComposite(1)
	one.Add(two, three)
	one.Operation()
}
