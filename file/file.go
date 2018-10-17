package file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func IsExist(dstFilePath string) (out bool, err error) {
	dstFilePath, err = filepath.Abs(dstFilePath)
	if err != nil {
		log.Println("Error while finding absolute path: ", dstFilePath, " - ", err)
		return false, err
	}
	if _, err := os.Stat(dstFilePath); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func ArrayToCsv(dstFilePath string, data [][]string) (err error) {
	dstFilePath, err = filepath.Abs(dstFilePath)
	if err != nil {
		log.Println("Error while finding absolute path: ", dstFilePath, " - ", err)
		return err
	}
	/*
		if _, err := os.Stat(dstFilePath); os.IsNotExist(err) {
			os.Mkdir(dstFilePath, 0777)
		}
	*/

	// Equal to: file, err := os.Create(dstFilePath)
	// Reference: https://golang.org/src/os/file.go?s=8039:8078#L263
	file, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("Error while creating file: ", dstFilePath, " - ", err)
		return err
	}
	defer file.Close()

	// 建立一個寫檔緩衝。
	fileWriter := csv.NewWriter(file)

	for _, record := range data {
		if err := fileWriter.Write(record); err != nil {
			log.Println("Error writing record to csv: ", dstFilePath, " - ", err)
			return err
		}
	}
	// Write any buffered data to the underlying writer (standard output).
	fileWriter.Flush()
	if err := fileWriter.Error(); err != nil {
		log.Println("Error while flushing writer: ", dstFilePath, " - ", err)
		return err
	}

	// 同步檔案。
	if err := file.Sync(); err != nil {
		log.Println("Error while syncing file: ", dstFilePath, " - ", err)
		return err
	}

	return nil
}

func ArrayAppendCsv(dstFilePath string, data [][]string) (err error) {
	dstFilePath, err = filepath.Abs(dstFilePath)
	if err != nil {
		log.Println("Error while finding absolute path: ", dstFilePath, " - ", err)
		return err
	}

	file, err := os.OpenFile(dstFilePath, os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Println("Error while creating fiel: ", dstFilePath, " - ", err)
		return err
	}
	defer file.Close()

	// 建立一個寫檔緩衝。
	fileWriter := csv.NewWriter(file)

	for _, record := range data {
		if err := fileWriter.Write(record); err != nil {
			log.Println("Error writing record to csv: ", dstFilePath, " - ", err)
			return err
		}
	}
	// Write any buffered data to the underlying writer (standard output).
	fileWriter.Flush()
	if err := fileWriter.Error(); err != nil {
		log.Println("Error while flushing writer: ", dstFilePath, " - ", err)
		return err
	}

	// 同步檔案。
	if err := file.Sync(); err != nil {
		log.Println("Error while syncing file: ", dstFilePath, " - ", err)
		return err
	}

	return nil
}

func CsvToArray(srcFilePath string) (data [][]string, err error) {
	srcFilePath, err = filepath.Abs(srcFilePath)
	if err != nil {
		log.Println("Error while finding absolute path: ", srcFilePath, " - ", err)
		return nil, err
	}

	file, err := os.Open(srcFilePath)
	if err != nil {
		log.Println("Error while creating file: ", srcFilePath, " - ", err)
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true
	// Read in the records one by one.
	for {
		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error while reading file: ", srcFilePath, " - ", err)
			return nil, err
		}

		// Append the record to our data set.
		data = append(data, record)
	}

	return data, nil
}

func StringsToFile(dstFilePath string, data []string) (err error) {
	dstFilePath, err = filepath.Abs(dstFilePath)
	if err != nil {
		log.Println("Error while finding absolute path", dstFilePath, "-", err)
		return err
	}

	// Equal to: file, err := os.Create(dstFilePath)
	// Reference: https://golang.org/src/os/file.go?s=8039:8078#L263
	file, err := os.OpenFile(dstFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Println("Error while creating file: ", dstFilePath, " - ", err)
		return err
	}
	defer file.Close()
	/*
		_, err = file.WriteString(data)
		if err != nil {
			fmt.Println("Error while writing ", dstFilePath, "-", err)
			return
		}
	*/
	// 建立一個寫檔緩衝。
	fileWriter := bufio.NewWriter(file)

	for i := 0; i < len(data); i++ {
		// 寫入緩衝內。
		fmt.Fprintf(fileWriter, "%s", data[i])
		// 將下一行分隔符號寫入緩衝內。
		fmt.Fprintln(fileWriter, "")
	}

	// Write any buffered data to the underlying writer (standard output).
	// 寫入檔案。
	if err := fileWriter.Flush(); err != nil {
		log.Println("Error while flushing file: ", dstFilePath, " - ", err)
		return err
	}

	// 同步檔案。
	if err := file.Sync(); err != nil {
		log.Println("Error while syncing file: ", dstFilePath, " - ", err)
		return err
	}

	return nil
}

func StringsAppendFile(dstFilePath string, data []string) (err error) {
	dstFilePath, err = filepath.Abs(dstFilePath)
	if err != nil {
		log.Println("Error while finding absolute path: ", dstFilePath, " - ", err)
		return err
	}

	file, err := os.OpenFile(dstFilePath, os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Println("Error while creating file: ", dstFilePath, " - ", err)
		return err
	}
	defer file.Close()

	// 建立一個寫檔緩衝。
	fileWriter := bufio.NewWriter(file)

	for i := 0; i < len(data); i++ {
		// 寫入緩衝內。
		fmt.Fprintf(fileWriter, "%s", data[i])
		// 將下一行分隔符號寫入緩衝內。
		fmt.Fprintln(fileWriter, "")
	}

	// Write any buffered data to the underlying writer (standard output).
	// 寫入檔案。
	if err := fileWriter.Flush(); err != nil {
		log.Println("Error while flushing file: ", dstFilePath, " - ", err)
		return err
	}

	// 同步檔案。
	if err := file.Sync(); err != nil {
		log.Println("Error while syncing file: ", dstFilePath, " - ", err)
		return err
	}

	return nil
}

// https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
// https://stackoverflow.com/questions/8757389/reading-file-line-by-line-in-go
func FileToStrings(srcFilePath string) (data []string, err error) {

	srcFilePath, err = filepath.Abs(srcFilePath)
	if err != nil {
		log.Println("Error while finding absolute path: ", srcFilePath, " - ", err)
		return nil, err
	}

	file, err := os.Open(srcFilePath)
	if err != nil {
		log.Println("Error while creating file: ", srcFilePath, " - ", err)
		return nil, err
	}
	defer file.Close()

	// 建立一個讀檔緩衝。
	//fileReader := bufio.NewReader(file)
	fileReader := bufio.NewScanner(file)

	for fileReader.Scan() {
		data = append(data, fileReader.Text())
	}
	if err := fileReader.Err(); err != nil {
		log.Println("Error while reading file: ", srcFilePath, " - ", err)
		return nil, err
	}

	return data, nil
}
