package main

import (
	"fmt"
	"strconv"
	"strings"
)

type FileKind int

const (
	FILE FileKind = iota
	DIRECTORY
)

type File struct {
	name  string
	size  int
	kind  FileKind
	owner *File
	files []*File
}

func (f *File) addFile(newFile *File) {
	newFile.owner = f
	f.files = append(f.files, newFile)
}

func (f *File) fileOfKindNamed(kind FileKind, name string) (foundFile *File) {
	for _, aFile := range f.files {
		if aFile.kind == kind && strings.EqualFold(aFile.name, name) {
			foundFile = aFile
			break
		}
	}
	return

}

func (f *File) collectDirectoriesLimitedInSizeTo(maxSize int, directories *[]File) {
	thisSize := 0
	for _, aFile := range f.files {
		if aFile.kind == DIRECTORY {
			aFile.collectDirectoriesLimitedInSizeTo(maxSize, directories)
			thisSize += aFile.sizeOf()
		} else {
			thisSize += aFile.sizeOf()
		}
	}
	if thisSize <= maxSize {
		*directories = append(*directories, *f)
	}
}

func (f *File) sizeOf() (aSize int) {

	for _, aFile := range f.files {
		if aFile.kind == DIRECTORY {
			aSize += aFile.sizeOf()
		} else if aFile.kind == FILE {
			aSize += aFile.size
		}
	}
	aSize += f.size
	return aSize
}

func (fs *FileSystem) list(name string) {

	words := strings.Split(name, " ")
	if strings.EqualFold(words[0], "dir") {
		name := words[1]
		anExistingFile := fs.currentDirectory().fileOfKindNamed(DIRECTORY, name)
		if anExistingFile == nil {
			// this is a directory
			fs.currentDirectory().addFile(&File{
				name:  name,
				size:  0,
				kind:  DIRECTORY,
				files: []*File{},
			})
		} else {
			panic("Already exists")
		}

	} else {
		// this is a file
		fileSize, _ := strconv.Atoi(words[0])

		name := words[1]
		anExistingFile := fs.currentDirectory().fileOfKindNamed(DIRECTORY, name)
		if anExistingFile == nil {
			fs.currentDirectory().addFile(&File{
				name:  words[1],
				size:  fileSize,
				kind:  FILE,
				files: []*File{},
			})
		} else {
			panic("Already exists")
		}

	}

}

type FileSystem struct {
	path []*File
}

func (fs *FileSystem) addToPath(aFile *File) *FileSystem {
	fs.path = append(fs.path, aFile)
	return fs
}

func (fs *FileSystem) goUpOneDirectory() *FileSystem {
	fs.path = fs.path[:len(fs.path)-1]
	return fs
}

func (fs *FileSystem) currentDirectory() *File {
	currentDirectory := fs.path[len(fs.path)-1]
	return currentDirectory
}

func (fs *FileSystem) changeDirectory(name string) *FileSystem {

	currentDirectory := fs.currentDirectory()
	if currentDirectory != nil {
		if strings.EqualFold(currentDirectory.name, name) {
			// no need to change
		} else {
			if strings.EqualFold("..", name) {
				fs.goUpOneDirectory()
			} else if strings.EqualFold("/", name) {
				fs.path = fs.path[0:1]

			} else {
				directoryToMoveTo := currentDirectory.fileOfKindNamed(DIRECTORY, name)

				if directoryToMoveTo == nil {
					panic(fmt.Sprintf("Unable to find directory: %s", name))
				} else if directoryToMoveTo.owner.name != currentDirectory.name {
					panic(fmt.Sprintf("Directory name mismatch: %s", name))
				}
				fs.addToPath(directoryToMoveTo)
			}

		}
	}
	return fs
}

func (fs *FileSystem) executeCommand(command CommandType, word string) {
	if command == CD {
		fs.changeDirectory(word)
	} else if command == LS {
		fs.list(word)
	}
}

type Part1 struct {
	answer int
}

type CommandType int

const (
	CD CommandType = iota
	LS
	UNKNOWN
)

func parse(data []string) (error, *FileSystem) {

	rootFile := &File{name: "/", kind: DIRECTORY, size: 0, owner: nil}
	fileSystem := &FileSystem{path: []*File{rootFile}}

	var lastCommand CommandType
	lastListCommandCounter := 0
	for _, row := range data {
		println(fmt.Sprintf("%s", row))
		words := strings.Split(row, " ")
		if words != nil && len(words) > 0 {
			if strings.EqualFold(words[0], "$") {
				for _, aWord := range words {
					if aWord == "$" {
						lastCommand = UNKNOWN

					} else if lastCommand == UNKNOWN && aWord == "cd" {
						lastCommand = CD
					} else if lastCommand == UNKNOWN && aWord == "ls" {
						lastCommand = LS
						lastListCommandCounter = 0
					} else {
						fileSystem.executeCommand(lastCommand, aWord)
					}
				}
			} else if lastCommand == LS {
				fileSystem.executeCommand(lastCommand, row)
				lastListCommandCounter++
				if len(fileSystem.currentDirectory().files) != lastListCommandCounter {
					panic("Files out of sync")
				}

			}

		}
	}

	return nil, fileSystem
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	_, fileSystem := parse(data)
	if fileSystem != nil {
		directories := make([]File, 0)
		rootFile := fileSystem.changeDirectory("/").currentDirectory()
		rootFile.collectDirectoriesLimitedInSizeTo(100000, &directories)
		totalSize := 0
		if directories != nil {
			for _, aDirectory := range directories {
				totalSize += aDirectory.sizeOf()
			}
		}
		alg.answer = totalSize
	}
	return nil, alg.answer
}
