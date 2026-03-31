package json

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileType int

const (
	NA FileType = iota
	Controller
	Service
	Repository
)

type FileCounts struct {
	Total        int `json:"total"`
	Controllers  int `json:"controllers"`
	Services     int `json:"services"`
	Repositories int `json:"repositories"`
}
type Summary struct {
	Target   string     `json:"target"`
	Packages string     `json:"packages"`
	Files    FileCounts `json:"files"`
}

type Violation struct {
	FilePath    string `json:"file_path"`
	PackageName string `json:"package_name"`
	ClassName   string `json:"class_name"`
	Violation   string `json:"violation"`
	Description string `json:"description"`
}

type Streamer struct {
	file           *os.File
	startedType    FileType // currently opened array type (NA if none)
	summaryWritten bool     // whether summary has been written
	firstItem      bool     // whether next item in the current array is the first (used to manage commas)
}

func NewJsonStreamer(filename string) (*Streamer, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	if _, err = file.WriteString("{\n"); err != nil {
		_ = file.Close()
		return nil, err
	}

	return &Streamer{file: file, startedType: NA, firstItem: true}, nil
}

func (js *Streamer) WriteSummary(summary Summary) error {
	summaryBytes, err := json.MarshalIndent(summary, "  ", "  ")
	if err != nil {
		return err
	}
	// write summary without a trailing comma; arrays (if any) will insert the comma before their key
	if _, err := js.file.WriteString(fmt.Sprintf("  \"summary\": %s", summaryBytes)); err != nil {
		return err
	}
	js.summaryWritten = true
	return nil
}

func fileTypeName(ft FileType) string {
	switch ft {
	case Controller:
		return "controller"
	case Service:
		return "service"
	case Repository:
		return "repository"
	default:
		return "unknown"
	}
}

func (js *Streamer) AddViolation(fileType FileType, violation Violation) error {
	// validate fileType
	if fileType == NA {
		return fmt.Errorf("invalid file type: NA")
	}

	// If starting a new category
	if fileType != js.startedType {
		// disallow going backward in order (must be non-decreasing: Controller -> Service -> Repository)
		if js.startedType != NA && fileType < js.startedType {
			// return an error to caller (do not record internally)
			return fmt.Errorf("file types out of order: tried to start %s after %s", fileTypeName(fileType), fileTypeName(js.startedType))
		}

		// close the previous array if any (append comma because another array will follow)
		if js.startedType != NA {
			if _, err := js.file.WriteString("\n  ],\n"); err != nil {
				return err
			}
		} else {
			// if no previous array but summary was written, append comma to separate keys
			if js.summaryWritten {
				if _, err := js.file.WriteString(",\n"); err != nil {
					return err
				}
			}
		}

		// start a new array key
		key := fileTypeName(fileType)
		if _, err := js.file.WriteString(fmt.Sprintf("  \"%s\": [\n", key)); err != nil {
			return err
		}

		js.startedType = fileType
		js.firstItem = true
	}

	// marshal violation as compact JSON (one-line) to simplify streaming and comma placement
	violationBytes, err := json.MarshalIndent(violation, "    ", "  ")
	if err != nil {
		return err
	}

	// write item with comma handling
	if js.firstItem {
		if _, err := js.file.WriteString(fmt.Sprintf("    %s", violationBytes)); err != nil {
			return err
		}
		js.firstItem = false
	} else {
		if _, err := js.file.WriteString(fmt.Sprintf(",\n    %s", violationBytes)); err != nil {
			return err
		}
	}

	return nil
}

func (js *Streamer) Close() error {
	// close the currently opened array (no trailing comma)
	if js.startedType != NA {
		if _, err := js.file.WriteString("\n  ]\n"); err != nil {
			return err
		}
	}

	// finish JSON
	if _, err := js.file.WriteString("}\n"); err != nil {
		return err
	}
	return js.file.Close()
}

func (js *Streamer) Delete() {
	if js.file != nil {
		_ = os.Remove(js.file.Name())
	}
}
