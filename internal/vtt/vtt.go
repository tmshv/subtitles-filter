package vtt

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type VTTRecord struct {
	time string
	text string
}

func (r *VTTRecord) Time() string {
	return r.time
}

func (r *VTTRecord) Text() string {
	return r.text
}

type VTT struct {
	head    string
	records []VTTRecord
}

func (s *VTT) Head() string {
    return s.head
}

func (s *VTT) Iter() <-chan VTTRecord {
	ch := make(chan VTTRecord)

	go func() {
		defer close(ch)

		for _, rec := range s.records {
			ch <- rec
		}
	}()

	return ch
}

func OpenFile(filename string) (*VTT, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		return nil, err
	}
    return Scan(file)
}

func Scan(buffer  io.Reader) (*VTT, error) {
	scanner := bufio.NewScanner(buffer)

	// Get head
	if !scanner.Scan() {
		return nil, nil
	}
	head := scanner.Text()

	// Empty line after head
	if !scanner.Scan() {
		return nil, nil
	}

	var records []VTTRecord
	for {
		if !scanner.Scan() {
			break
		}
		time := scanner.Text()

		var texts []string
		for scanner.Scan() {
			row := scanner.Text()

			// End of record
			if row == "" {
				break
			}

			texts = append(texts, row)
		}

		records = append(records, VTTRecord{
			time: time,
			text: strings.Join(texts, "\n"),
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &VTT{
		head:    head,
		records: records,
	}, nil
}
