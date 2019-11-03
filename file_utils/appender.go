package file_utils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"io"
	"strings"
)

func findTextOnLine(path string, substr string ) (int, error){
	log.Println(path, " PATH ", substr, " ")
	f, err := os.Open(path)
	if err != nil {
	return 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 1

	for scanner.Scan() {
	if strings.Contains(scanner.Text(), substr) {
	return line, nil
	}

	line++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return line, nil
}

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

/**
 * Insert sting to n-th line of file.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFile(path string , str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func AppendStringToLine(path string, substr string, textToAdd string ) (error){
	log.Println("APPEND STRING TO FILE")
	lineNumber, err := findTextOnLine(path, substr)
	log.Print(path, substr, textToAdd)

	log.Println(err, lineNumber, "LINE NUMBER")
	err = InsertStringToFile(path, textToAdd, lineNumber)
	if err != nil {
		return err
	}

	return nil
}
