package core

import (
	"errors"
	"io/fs"
	"os"
)

// Open a file with the default settings
func OpenFile(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Delete a file
func DeleteFile(path string) error {
	err := os.Remove(path)
	return err
}

// Read the length of a file
func LengthOfFile(file *os.File) (int64, error) {
	stat, err := os.Stat(file.Name())
	if err != nil {
		return 0, err
	}

	return stat.Size(), nil
}

// Read a file by seeking to an offset from the origin and reading up to len bytes
func ReadFile(file *os.File, offset int64, len int) ([]byte, error) {
	off, err := file.Seek(offset, 0)
	if err != nil {
		return nil, err
	} else if off != offset {
		return nil, errors.New("non-matching file offsets")
	}

	buf := make([]byte, len)
	i, err := file.Read(buf)
	if err != nil {
		return nil, err
	} else if i != len {
		return nil, errors.New("invalid file length read")
	}

	return buf, nil
}

// Write data to a section of a file
// Note: Does not flush the changes to the file system
func WriteToFile(file *os.File, offset int64, data []byte) error {
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	off, err := file.Seek(offset, 0)
	if err != nil {
		return err
	} else if off != offset {
		return errors.New("non-matching file offsets")
	}

	res, err := file.Write(data)
	if err != nil {
		return err
	} else if res != len(data) {
		return errors.New("invalid file length written")
	}

	return nil
}

// Append data to the end of a file
// Note: Does not flush the changes to the file system
func AppendToFile(file *os.File, data []byte) error {
	size, err := LengthOfFile(file)
	if err != nil {
		return err
	}

	err = WriteToFile(file, size, data)
	if err != nil {
		return err
	}

	return nil
}

// Create a new directory with the default settings
func Mkdir(path string) error {
	return os.Mkdir(path, 0755)
}

func ReadDir(path string) ([]fs.DirEntry, error) {
	return os.ReadDir(path)
}

// Delete an empty directory
func Rmdir(path string) error {
	return os.Remove(path)
}
