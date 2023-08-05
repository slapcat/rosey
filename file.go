package main

import (
	"fmt"
	"time"
	"os"
	"io"
	"bytes"
	"log"
	"bufio"
	"github.com/robfig/cron/v3"
)

	var position int64
	var lines int

func ReadFile() {

	position = 0
	lines = 0

	c := cron.New()
	c.AddFunc(Timer, func() {

	    file, err := os.Open(Source)
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer file.Close()

		lineCount, err := LineCounter(file)
		if err != nil {
			log.Fatal(err)
		}

		if  lineCount == lines {
			position = 0
			lines = 0
		}

		errr := LineReader(file, position)
		if errr != nil {
			fmt.Println("Scanner error:", err)
		}
	})
	c.Start()

for {
	time.Sleep(time.Second)
}

}

func LineCounter(r io.Reader) (int, error) {

    var count int
    const lineBreak = '\n'

    buf := make([]byte, bufio.MaxScanTokenSize)

    for {
        bufferSize, err := r.Read(buf)
        if err != nil && err != io.EOF {
            return 0, err
        }

        var buffPosition int
        for {
            i := bytes.IndexByte(buf[buffPosition:], lineBreak)
            if i == -1 || bufferSize == buffPosition {
                break
            }
            buffPosition += i + 1
            count++
        }
        if err == io.EOF {
            break
        }
    }

    return count, nil
}


func LineReader(input io.ReadSeeker, start int64) error {
	if _, err := input.Seek(start, 0); err != nil {
		return err
	}
	scanner := bufio.NewScanner(input)

	pos := start
	scanLines := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		pos += int64(advance)
		return
	}
	scanner.Split(scanLines)

	for scanner.Scan() {
			// Toot next line
			Toot(fmt.Sprintf("%s", scanner.Text()))
			// Send to terminal for debug
			//fmt.Printf("%s", scanner.Text())
			position = pos
			lines++
			break
	}
	return scanner.Err()
}
