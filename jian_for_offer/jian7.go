package jian_for_offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	in := make(map[int]int, len(inorder))
	for i := 0; i < len(inorder); i++ {
		in[inorder[i]] = i
	}
	return help(preorder, 0, len(preorder)-1, in, 0, len(inorder)-1)
}

func help(preorder []int, startPre int, endPre int, in map[int]int, startIn int, endIn int) *TreeNode {
	if startPre > endPre || startIn > endIn {
		return nil
	}
	root := &TreeNode{
		Val: preorder[startPre],
	}
	index, ok := in[preorder[startPre]]
	if !ok {
		return nil
	}
	root.Left = help(preorder, startPre+1, index-startIn+startPre, in, startIn, index-1)
	root.Right = help(preorder, index-startIn+startPre+1, endPre, in, index+1, endIn)
	return root
}
