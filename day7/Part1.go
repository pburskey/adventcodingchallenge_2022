package main

import (
	"strings"
)

type FileKind int

const (
	FILE FileKind = iota
	DIRECTORY
)

type File struct {
	name string
	size int
	kind FileKind
}

type FileSystem struct {
	files            []*File
	currentDirectory *File
}

func (fs *FileSystem) changeDirectory(name string) {

}

type Part1 struct {
	answer int
}

type CommandType int

const (
	CD CommandType = iota
	LS
	UNKNONW
)

func parse(data []string) (error, *FileSystem) {

	fileSystem := &FileSystem{}
	var lastCommand CommandType

	for _, row := range data {
		words := strings.Split(row, " ")
		if words != nil && len(words) > 0 {
			if words[0] == "$" {
				lastCommand = UNKNONW
				//parsing a command
				if words[1] == "cd" {
					lastCommand = CD
				} else if words[1] == "ls" {
					lastCommand = LS
				}
			} else {
				if lastCommand == LS {

				} else if lastCommand == CD {

				}
			}
		}
	}
	return nil, fileSystem
}

func (alg *Part1) Process(data []string) (error, interface{}) {
	_, fileSystem := parse(data)
	if fileSystem != nil {

	}
	return nil, alg.answer
}
