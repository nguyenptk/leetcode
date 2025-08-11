// https://leetcode.com/problems/design-in-memory-file-system/
package hard

import (
	"sort"
	"strings"
)

type NodeFS struct {
	paths         map[string]*NodeFS
	name, content string
	isFile        bool
}

type FileSystem struct {
	root *NodeFS
}

func ConstructorFileSystem() FileSystem {
	return FileSystem{
		root: &NodeFS{
			paths:  make(map[string]*NodeFS),
			name:   "",
			isFile: false,
		},
	}
}

func (f *FileSystem) Ls(path string) []string {
	curr := f.root
	ds := splitPath(path)
	for _, dir := range ds {
		curr = curr.paths[dir]
	}
	if curr.isFile {
		return []string{ds[len(ds)-1]}
	}
	paths := make([]string, 0, len(curr.paths))
	for name := range curr.paths {
		paths = append(paths, name)
	}
	sort.Strings(paths)
	return paths
}

func (f *FileSystem) Mkdir(path string) {
	curr := f.root
	for _, dir := range splitPath(path) {
		if _, ok := curr.paths[dir]; !ok {
			curr.paths[dir] = &NodeFS{
				paths:  make(map[string]*NodeFS),
				name:   dir,
				isFile: false,
			}
		}
		curr = curr.paths[dir]
	}
}

func (f *FileSystem) AddContentToFile(filePath string, content string) {
	dirs := splitPath(filePath)
	// We don't go all the way to the end of the path because the problem
	// description says we can assume valid input. In that case, the last
	// element must be the file.
	curr := f.root
	for _, dir := range dirs[:len(dirs)-1] {
		curr = curr.paths[dir]
	}
	name := dirs[len(dirs)-1]
	file, ok := curr.paths[name]
	if !ok {
		file = &NodeFS{
			paths:  make(map[string]*NodeFS),
			name:   name,
			isFile: true,
		}
		curr.paths[name] = file
	}
	file.content += content
}

func (f *FileSystem) ReadContentFromFile(filePath string) string {
	curr := f.root
	for _, dir := range splitPath(filePath) {
		curr = curr.paths[dir]
	}
	return curr.content
}

func splitPath(path string) []string {
	if path == "/" {
		return []string{}
	}
	return strings.Split(path, "/")[1:]
}

// | Method                | Time Complexity | Space Complexity |
// | --------------------- | --------------- | ---------------- |
// | `Constructor`         | O(1)            | O(1)             |
// | `Mkdir`               | O(L)            | O(L)             |
// | `AddContentToFile`    | O(L + C)        | O(C)             |
// | `ReadContentFromFile` | O(L)            | O(1)             |
// | `Ls`                  | O(L + D log D)  | O(D)             |
// | `splitPath`           | O(N)            | O(L)             |
// L is the number of directories in the path.
// C is the length of the content to append.
// D is the number of items in the target directory (for sorting).
// N is the length of the path string.

/**
 * Your FileSystem object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ls(path);
 * obj.Mkdir(path);
 * obj.AddContentToFile(filePath,content);
 * param_4 := obj.ReadContentFromFile(filePath);
 */
