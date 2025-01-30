package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("could not open file")
	}

	// defer is a special statement that executes
	// the function after the outer function completes.
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	// .Scan() returns False when the scan stops.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		// file.Close() // use `defer` instead
		return nil, errors.New("failed to scan the file")
	}

	// file.Close()
	return lines, nil
}

// interface{} any value is allowed
func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	// simulate a long process
	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		// file.Close() // use `defer` instead
		return errors.New("failed to convert data to json")
	}
	// file.Close()
	return nil
}

// constructor
func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
