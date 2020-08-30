package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

// ReadFile read a file line by line, and returns a slice.
func ReadFile(filePath string) []string {
	lines := make([]string, 0)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

// SliceAtoi convert string slice to int64 slice.
func SliceAtoi(texts []string) ([]int64, error) {
	numbers := make([]int64, 0, len(texts))
	for _, text := range texts {
		num, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, int64(num))
	}
	return numbers, nil
}

// FakeStdin emulates stdin. See https://eli.thegreenplace.net/2020/faking-stdin-and-stdout-in-go/
type FakeStdin struct {
	stdinWriter *os.File
	origStdin   *os.File
}

// NewFakeStdin is constructor of Fakestdin
func NewFakeStdin(stdinText string) (*FakeStdin, error) {
	stdinReader, stdinWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	origStdin := os.Stdin
	os.Stdin = stdinReader

	_, err = stdinWriter.Write([]byte(stdinText))
	if err != nil {
		stdinWriter.Close()
		os.Stdin = origStdin
		return nil, err
	}

	return &FakeStdin{
		stdinWriter: stdinWriter,
		origStdin:   origStdin,
	}, nil
}

// Close stops the emulation.
func (fakestdin *FakeStdin) Close() error {
	if fakestdin.origStdin == nil {
		return errors.New("instance has already been closed")
	}
	fakestdin.stdinWriter.Close()
	os.Stdin.Close()
	os.Stdin = fakestdin.origStdin
	fakestdin.origStdin = nil
	return nil
}

type Fakestdout struct {
	stdoutReader *os.File
	origStdout   *os.File

	outCh chan []byte
}

func NewFakeStdout() (*Fakestdout, error) {
	stdoutReader, stdoutWriter, err := os.Pipe()
	if err != nil {
		return nil, err
	}
	originStdout := os.Stdout
	os.Stdout = stdoutWriter

	outCh := make(chan []byte)

	// This goroutine reads stdout into a buffer in the background.
	go func() {
		var b bytes.Buffer
		if _, err := io.Copy(&b, stdoutReader); err != nil {
			log.Println(err)
		}
		outCh <- b.Bytes()
	}()

	return &Fakestdout{
		origStdout:   originStdout,
		stdoutReader: stdoutReader,
		outCh:        outCh,
	}, nil
}

func (fakestdout *Fakestdout) ReadandClose() string {
	if fakestdout.origStdout == nil {
		panic("instance has already been closed")
	}
	os.Stdout.Close()
	storedChars := string((<-fakestdout.outCh)[:])
	fakestdout.Close()
	return storedChars
}

func (fakestdout *Fakestdout) Close() {
	fakestdout.stdoutReader.Close()
	os.Stdout = fakestdout.origStdout
}

// Timeout timeouts a channel of int64.
func Timeout(channel <-chan int64, d time.Duration) (int64, error) {
	timeCh := time.After(d)
	select {
	case d := <-channel:
		return d, nil
	case <-timeCh:
		return -1, errors.New("timeout")
	}
}
