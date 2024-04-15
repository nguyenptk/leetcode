// https://leetcode.com/problems/implement-trie-prefix-tree/

package java.medium;

class Trie {

    Node root;

    public Trie() {
        root = new Node('\0'); // dummy node
    }

    public void insert(String word) {
        Node curr = root;
        for (char x : word.toCharArray()) {
            if (curr.children[x-'a'] == null) {
                curr.children[x-'a'] = new Node(x);
            }
            curr = curr.children[x-'a'];
        }
        curr.isEnd = true;
    }

    public boolean search(String word) {
        Node res = getLast(word);
        return (res != null && res.isEnd);
    }

    public Node getLast(String word) {
        Node curr = root;
        for (char x : word.toCharArray()) {
            if (curr.children[x-'a'] == null) {
                return null;
            }
            curr = curr.children[x-'a'];
        }
        return curr;
    }

    public boolean startsWith(String prefix) {
        Node res = getLast(prefix);
        return res != null;
    }

    class Node {
        private Node[] children;
        private boolean isEnd;

        public Node(char val) {
            this.isEnd = false;
            this.children = new Node[26];
        }
    }
}