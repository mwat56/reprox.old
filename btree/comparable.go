/*
   Copyright © 2020 M.Watermann, 10247 Berlin, Germany
               All rights reserved
           EMail : <support@mwat.de>
*/

package btree

type (
	// TNode is a structure representing a leave in the binary tree.
	TNode struct {
		nData  interface{}
		nLeft  *TNode
		nRight *TNode
	}

	// TComparable is an interface requiring a function to return
	// an integer value signalling whether two objects are equal.
	TComparable interface {
		// Compare returns an integer value of `0` (zero) if 'aNode1`
		// equals `aNode2`, or `1` (plus one) if 'aNode1` is greater
		// than is greater than `aNode2`, or `-1` (minus one) if
		// 'aNode1` is less than `aNode2`.
		//
		//	`aNode1` is the tree node to compare.
		//	`aNode2` is the tree node to compare with.
		Compare(aNode1, aNode2 TNode) int
	}

	// TInt is an `int` type with associated methods (see below).
	TInt int
)

func (ti TInt) compare(aNode TNode) int {
	if n2, ok := aNode.nData.(int); ok {
		n1 := int(ti)
		if n1 == n2 {
			return 0
		}
		if n1 < n2 {
			return -1
		}
	}

	// _this_ node is considered greater than `aNode`.
	return 1
} // Compare()

// Data returns the node's data.
func (tn *TNode) Data() interface{} {
	return tn.nData
} // Data()

func eq(aNode1, aNode2 interface{}) bool {
	return aNode1 == aNode2
} // eq()

func (tn *TNode) lt(aNode *TNode) bool {
	switch tn.nData.(type) {
	case int:
		return tn.nData.(int) < aNode.nData.(int)
	case string:
		return tn.nData.(string) < aNode.nData.(string)
	}
	return false
} // lt()

/* _EoF_ */
