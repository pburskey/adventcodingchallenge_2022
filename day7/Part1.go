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
		anExistingFile := fs.getCurrentDirectory().fileOfKindNamed(DIRECTORY, name)
		if anExistingFile == nil {
			// this is a directory
			fs.getCurrentDirectory().addFile(&File{
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
		anExistingFile := fs.getCurrentDirectory().fileOfKindNamed(DIRECTORY, name)
		if anExistingFile == nil {
			fs.getCurrentDirectory().addFile(&File{
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
	//path             []*File
	currentDirectory *File
	root             *File
	offset           int
}

func (fs *FileSystem) addToPath(aFile *File) *FileSystem {
	//fs.path = append(fs.path, aFile)
	fs.currentDirectory = aFile

	//if fs.getCurrentDirectory().name != fs.currentDirectory.name {
	//	panic("Current directory out of sync")
	//}
	fs.offset++
	return fs
}

func (fs *FileSystem) changeToRoot() *FileSystem {

	fs.currentDirectory = fs.root
	//fs.path = fs.path[0:1]
	//if fs.getCurrentDirectory().name != fs.currentDirectory.name {
	//	panic("Current directory out of sync")
	//}
	fs.offset = 0

	return fs
}

func (fs *FileSystem) goUpOneDirectory() *FileSystem {
	owner := fs.currentDirectory.owner
	fs.currentDirectory = owner
	fs.offset--
	//fs.path = fs.path[:len(fs.path)-1]
	//if fs.getCurrentDirectory().name != fs.currentDirectory.name {
	//	panic("Current directory out of sync")
	//}

	return fs
}

func (fs *FileSystem) getCurrentDirectory() *File {
	//currentDirectory := fs.path[len(fs.path)-1]
	//if currentDirectory.name != fs.currentDirectory.name {
	//	panic("Current directory out of sync")
	//}
	return fs.currentDirectory
}

func (fs *FileSystem) changeDirectory(name string) *FileSystem {

	startingOffSet := fs.offset
	expectation := 0
	currentDirectory := fs.getCurrentDirectory()
	if currentDirectory != nil {

		if strings.EqualFold("..", name) {
			fs.goUpOneDirectory()
			expectation = startingOffSet - 1
		} else if strings.EqualFold("/", name) {
			fs.changeToRoot()
			expectation = 0

		} else {
			directoryToMoveTo := currentDirectory.fileOfKindNamed(DIRECTORY, name)

			fs.addToPath(directoryToMoveTo)
			expectation = startingOffSet + 1
		}

	}

	if expectation != fs.offset {
		panic("Offset difference")
	}

	return fs
}

func (fs *FileSystem) executeCommand(command CommandType, word string) *FileSystem {
	if command == CD {
		fs.changeDirectory(word)
	} else if command == LS {
		fs.list(word)
	}
	return fs
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
	fileSystem := &FileSystem{root: rootFile, currentDirectory: rootFile}

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

					} else if lastCommand == UNKNOWN && strings.EqualFold(aWord, "cd") {
						lastCommand = CD
					} else if lastCommand == UNKNOWN && strings.EqualFold(aWord, "ls") {
						lastCommand = LS
						lastListCommandCounter = 0
					} else {
						fileSystem.executeCommand(lastCommand, aWord)
					}
				}
			} else if lastCommand == LS {
				fileSystem.executeCommand(lastCommand, row)
				lastListCommandCounter++
				if len(fileSystem.getCurrentDirectory().files) != lastListCommandCounter {
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
		rootFile := fileSystem.changeDirectory("/").getCurrentDirectory()
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
