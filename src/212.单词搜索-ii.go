/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */
package leetcode

// @lc code=start
type _Trie struct {
	children [26]*_Trie
	word string
}

func (t *_Trie) Insert(word string) {
	node := t
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &_Trie{}
		}
		node = node.children[ch]
	}
	// node.isEnd = true
	node.word = word
}

func findWords(board [][]byte, words []string) (ans []string) {
	t := &_Trie{}

	for _, word := range words {
		t.Insert(word)
	}

	n, m := len(board), len(board[0])

	directins := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	seen := map[string]bool{}

	var (
		backtrack func(x, y int, node *_Trie)
	)

	backtrack = func(x,y int, node *_Trie) {
		ch := board[x][y]
		node = node.children[ch - 'a']

		if node == nil {
			return
		}

		if node.word != "" {
			seen[node.word] = true
		}

		board[x][y] = '#'
		for _, v := range directins {
			newX := x + v[0]
			newY := y + v[1]
		
			if newX >= 0 && newX < n && newY >= 0 && newY < m && board[newX][newY] != '#' {
				backtrack(newX, newY, node)
			}
		}
		board[x][y] = ch
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			backtrack(i, j, t)
		}
	}

	for s := range seen {
		ans = append(ans, s)
	}

	return
}
// @lc code=end

