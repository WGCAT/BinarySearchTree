package main

import "fmt"

// 노드 타입 정의
type Node struct {
	key   int
	left  *Node
	right *Node
}

// 루트 노드 정의
type Bst struct {
	root *Node
}

// 1. 탐색
// 찾을 키를 입력했을 때 해당 키를 갖는 노드를 반환
func (b *Bst) Search(findKey int) *Node {
	return search(b.root, findKey)
}

// 현재노드와 찾을 키를 매개변수로 받는 재귀함수
func search(now *Node, findKey int) *Node {
	if now == nil || now.key == findKey {
		return now
	}
	if findKey < now.key {
		return search(now.left, findKey)
	}
	return search(now.right, findKey)
}

// 1-1. 최댓값 탐색
func max(now *Node) *Node {
	for {
		if now != nil && now.right != nil {
			now = now.right
		} else {
			break
		}
	}
	return now
}

// 1-2. 최소값 탐색
func min(now *Node) *Node {
	for {
		if now != nil && now.left != nil {
			now = now.left
		} else {
			break
		}
	}
	return now
}

// 2. 삽입
func (b *Bst) Insert(newKey int) {
	new := &Node{
		key:   newKey,
		left:  nil,
		right: nil,
	}
	if b.root == nil {
		b.root = new
	} else {
		insert(b.root, new)
	}
}

// 현재노드와 새노드를 매개변수로 받는 재귀함수
func insert(now *Node, new *Node) {
	// 새 키가 현재 키보다 작다면
	if new.key < now.key {
		// 왼쪽 자녀노드 없다면
		if now.left == nil {
			now.left = new
			// 왼쪽 자녀노드 있다면
		} else {
			insert(now.left, new)
		}
	} else if new.key > now.key {
		if now.right == nil {
			now.right = new
		} else {
			insert(now.right, new)
		}
	}
	// n.key == new.key 암것도 안함
}

// 3. 삭제
func (b *Bst) Delete(deleteKey int) {
	delete(b.root, deleteKey)
}

// 현재노드와 삭제할 키를 매개변수로 받는 재귀함수
func delete(now *Node, deleteKey int) *Node {
	if now == nil {
		return nil
	} else if deleteKey < now.key {
		now.left = delete(now.left, deleteKey)
		return now
	} else if deleteKey > now.key {
		now.right = delete(now.right, deleteKey)
		return now
	} else {
		// 3-1. 삭제할 노드가 리프 노드인 경우
		if now.left == nil && now.right == nil {
			now = nil
			return nil
		}
		// 3-2. 삭제할 노드가 1개의 자녀 노드만 가지고 있는 경우
		// 오른자녀만 지닐경우
		if now.left == nil {
			now = now.right
			return now
		}
		// 왼자녀만 지닐경우
		if now.right == nil {
			now = now.left
			return now
		}

		// 3-3. 삭제할 노드가 2개 이상의 자녀 노드를 가지고 있는 경우
		now = min(now.right)
		delete(now.right, now.key)
	}
	return now
}

func main() {
	tree := &Bst{
		root: &Node{
			key:   7,
			left:  nil,
			right: nil,
		},
	}
	tree.Insert(4)
	tree.Insert(9)
	fmt.Println(tree.Search(9))

	// tree.Insert(12)
	// tree.Insert(11)
	tree.Insert(8)
	fmt.Println(tree.Search(9))
	tree.Delete(4)
	fmt.Println(tree.Search(4))

}
