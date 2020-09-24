package main

import (
	"crypto/rand"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	err := createTestFile()
	if err != nil {
		t.Skip(err)
	}

	err = Copy("testFile1", "testFile2", 256, 256)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Open("testFile2")
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}
	file.Close()

	for i := 0; i < 256; i++ {
		if b[i] != '2' {
			t.FailNow()
		}
	}

	err = deleteTestFiles()
	if err != nil {
		t.Skip(err)
	}
}

func createTestFile() error {
	testFile, err := os.OpenFile("testFile1", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer testFile.Close()

	testString := make([]byte, 512)
	for i := 0; i < 256; i++ {
		testString[i] = '1'
	}
	for i := 0; i < 256; i++ {
		testString[i + 256] = '2'
	}

	_, err = testFile.Write(testString)
	if err != nil {
		return err
	}

	reader := io.LimitReader(rand.Reader, 1024)
	_, err = io.Copy(testFile, reader)
	if err != nil {
		return err
	}

	return nil
}

func deleteTestFiles() error {
	err := os.Remove("testFile1")
	if err != nil {
		return err
	}

	err = os.Remove("testFile2")
	if err != nil {
		return err
	}

	return nil
}