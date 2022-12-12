package main

type Part2 struct {
	answer int
}

func (alg *Part2) Process(data []string) (error, interface{}) {
	maxSpace := 70000000
	targetSpace := 30000000
	unusedSpace := 0
	needToPurgeTarget := 0
	_, fileSystem := parse(data)
	if fileSystem != nil {
		usedSpace := fileSystem.root.sizeOf()
		unusedSpace = maxSpace - usedSpace
		needToPurgeTarget = targetSpace - unusedSpace
	}

	if needToPurgeTarget > 0 {
		alg.FindSmallestDirectoryGreaterInSizeThan(fileSystem.root, needToPurgeTarget)
	}

	return nil, alg.answer

}

func (alg *Part2) FindSmallestDirectoryGreaterInSizeThan(aFileToConsider *File, targetSize int) {

	thisSize := 0
	for _, aFile := range aFileToConsider.files {
		if aFile.kind == DIRECTORY {
			alg.FindSmallestDirectoryGreaterInSizeThan(aFile, targetSize)
			thisSize += aFile.sizeOf()
		} else {
			thisSize += aFile.sizeOf()
		}
	}
	if thisSize >= targetSize {
		if alg.answer == 0 || thisSize < alg.answer {
			alg.answer = thisSize
		}

	}

}
