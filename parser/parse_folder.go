package parser

import (
	"fmt"
	"os"
	"strings"
)

type ParsedFolder struct {
	Path []string
	Name string

	Folders []ParsedFolder
	Files   []ParsedFile
}

func ParseFolder(rawPath string) (*ParsedFolder, error) {
	parts := strings.Split(rawPath, "/")

	parsed := &ParsedFolder{
		Name: parts[len(parts)-1],
		Path: parts[0 : len(parts)-1],
	}

	entries, err := os.ReadDir(rawPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			childFolder, err := ParseFolder(fmt.Sprintf("%s/%s",
				rawPath, entry.Name()))

			if err != nil {
				return nil, err
			}

			parsed.Folders = append(parsed.Folders, *childFolder)
		} else if strings.HasSuffix(entry.Name(), ".py") {
			fileBytes, err := os.ReadFile(fmt.Sprintf("%s/%s",
				rawPath, entry.Name()))

			if err != nil {
				return nil, err
			}

			childFile, err := ParseFile(string(fileBytes))
			if err != nil {
				return nil, err
			}

			childFile.Dir = rawPath
			childFile.Name = entry.Name()

			parsed.Files = append(parsed.Files, *childFile)
		}
	}

	return parsed, nil
}
