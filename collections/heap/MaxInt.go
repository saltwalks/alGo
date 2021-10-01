package heap

import "errors"

/* ----- PUBLIC ----- */
func NewMaxInt(size int) *MaxInt {
	return &MaxInt{0, make([]int, size)}
}

/**
 * Push one element to min heap
 */
func (h *MaxInt) Push(el int) error {
	// if heap is full, returns error.
	if h.Tail >= len(h.Tree) { return errors.New("heap: full") }

	// add element in tail of Complete Binary Tree.
	h.Tree[h.Tail] = el

	// if two or more element is in CBT, rearrange it from bottom to top.
	if h.Tail > 0 { h.sortAfterPush(h.Tail) }

	// move tail
	h.Tail++
	return nil
}

func (h *MaxInt) Pop() (int, error) {
	// if heap is empty, returns error.
	if h.Tail == 0 { return 0, errors.New("heap: empty") }

	// save head of Complete Binary Tree.
	target := h.Tree[0]

	// allocate last element of CBT to head.
	h.Tree[0] = h.Tree[h.Tail - 1]

	// move tail
	h.Tail--

	// if two or more element is in CBT, rearrange it from top to bottom.
	if h.Tail > 1 { h.sortAfterPop(0) }
	return target, nil
}

/* ----- PRIVATE ----- */
/**
 * Rearrange Complete Binary Tree from bottom to top.
 */
func (h *MaxInt) sortAfterPush(cur int) {
	// Notice that cur means index of current node and (cur - 1) / 2 means index of parent node.
	//     if current node is not root and value of parent node is smaller than current node, swap value.
	//     and check validity of parent node recursively.
	if cur > 0 && h.Tree[cur] > h.Tree[(cur - 1) / 2] {
		h.Tree[cur], h.Tree[(cur - 1) / 2] = h.Tree[(cur - 1) / 2], h.Tree[cur]
		h.sortAfterPush((cur - 1) / 2)
	}
}

func (h *MaxInt) sortAfterPop(cur int) {
	// Notice that cur means index of current node and cur * 2 + 1 means index of left-child node.
	//     if left-child is not tail and value of left-child node is bigger than current node, swap value.
	//     and check validity of left-child node recursively.
	if cur * 2 + 1 < h.Tail && h.Tree[cur * 2 + 1] > h.Tree[cur] {
		h.Tree[cur * 2 + 1], h.Tree[cur] = h.Tree[cur], h.Tree[cur * 2 + 1]
		h.sortAfterPop(cur * 2 + 1)
	}

	// Notice that cur means index of current node and cur * 2 + 2 means index of right-child node.
	//     if right-child is not tail and value of right-child node is bigger than current node, swap value.
	//     and check validity of right-child node recursively.
	if cur * 2 + 2 < h.Tail && h.Tree[cur * 2 + 2] > h.Tree[cur] {
		h.Tree[cur * 2 + 2], h.Tree[cur] = h.Tree[cur], h.Tree[cur * 2 + 2]
		h.sortAfterPop(cur * 2 + 2)
	}
}