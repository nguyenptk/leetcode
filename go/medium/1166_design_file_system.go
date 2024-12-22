// https://leetcode.com/problems/design-file-system
package medium

import "strings"

type TrieFile struct {
	val  int
	path map[string]*TrieFile
}

type FileSystem struct {
	root *TrieFile
}

func ConstructorFileSystem() FileSystem {
	return FileSystem{root: &TrieFile{
		val:  0,
		path: make(map[string]*TrieFile),
	}}
}

func (f *FileSystem) CreatePath(path string, value int) bool {
	comps := strings.Split(path, "/")
	curr := f.root
	for i := 1; i < len(comps); i++ {
		if next, ok := curr.path[comps[i]]; ok {
			if i == len(comps)-1 { // the path already exists
				return false
			} else {
				curr = next
			}
		} else {
			if i == len(comps)-1 { // add a new path by the end of the trie
				curr.path[comps[i]] = &TrieFile{
					val:  value,
					path: make(map[string]*TrieFile),
				}
				return true
			} else { // the parent path doesn't exist
				return false
			}
		}
	}

	return true
}

func (f *FileSystem) Get(path string) int {
	comps := strings.Split(path, "/") // ["leet", "code"]
	curr := f.root
	for i := 1; i < len(comps); i++ {
		if next, ok := curr.path[comps[i]]; !ok {
			return -1
		} else {
			curr = next
		}
	}
	return curr.val
}

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.CreatePath(path,value);
 * param_2 := obj.Get(path);
 */
