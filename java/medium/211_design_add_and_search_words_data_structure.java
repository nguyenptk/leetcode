// https://leetcode.com/problems/design-add-and-search-words-data-structure/

package java.medium;

class WordDictionary {

    // Init the Node class
    class Node {
        private Node[] children = new Node[26];
        boolean isEnd = false;
        public Node() {}
    }

    Node root;

    public WordDictionary() {
        root = new Node();
    }

    public void addWord(String word) {
        Node curr = this.root;
        for (int i = 0; i < word.length(); i++) {
            char ch = word.charAt(i);

            if (curr.children[ch-'a'] == null) {
                curr.children[ch-'a'] = new Node();
            }

            curr = curr.children[ch-'a'];
        }

        curr.isEnd = true;
    }

    public boolean search(String word) {
        return dfsSearch(0, this.root, word);
    }

    private boolean dfsSearch(int idx, Node curr, String word) {
        for (int i = idx; i < word.length(); i++) {
            char ch = word.charAt(i);
            if (ch == '.') {
                for (Node tmp : curr.children) {
                    if (tmp != null && dfsSearch(i+1, tmp, word)) {
                        return true;
                    }
                }
                return false;
            }
            if (curr.children[ch-'a'] == null) {
                return false;
            }
            curr = curr.children[ch-'a'];
        }

        return curr.isEnd;
    }
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary obj = new WordDictionary();
 * obj.addWord(word);
 * boolean param_2 = obj.search(word);
 */