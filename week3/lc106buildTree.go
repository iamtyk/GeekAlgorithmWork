package week3

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
中、后查找就是后续遍历逐个从右子树种过去，如果中序中已经没有右边了，就说明右边到头了，该找左子树了
*/
func buildTree(inorder []int, postorder []int) *TreeNode {
	mp := map[int]int{}
	for i, val := range inorder {
		mp[val] = i
	}

	var build func(int, int) *TreeNode

	build = func(iLeft, iRight int) *TreeNode {
		if iLeft > iRight {
			return nil
		}
		root := &TreeNode{}
		//后续遍历逐个减少
		root.Val = postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		inorderRootIndex := mp[root.Val]
		//中序遍历找到根节点后，右子树就是根节点的右边到头
		root.Right = build(inorderRootIndex+1, iRight)
		//左子树就是左边到根节点的左边
		root.Left = build(iLeft, inorderRootIndex-1)
		return root
	}

	return build(0, len(inorder)-1)
}
