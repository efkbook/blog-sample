package scenario

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type shakespeareRecord struct {
	LineID       int    `json:"line_id"`
	PlayName     string `json:"play_name"`
	SpeechNumber int    `json:"speech_number"`
	Speaker      string `json:"speaker"`
	Text         string `json:"text_entry"`
}

func Shakespeare() error {
	if err := loginOrSignup("shakespeare@example.com", "shakespeare", "password"); err != nil {
		return err
	}
	if err := create("begin", fmt.Sprintf("agent started at %v", time.Now())); err != nil {
		return err
	}

	f, err := os.OpenFile("shakespeare.json", os.O_RDONLY, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	title := ""
	text := []string{}
	for {
		line, isPrefix, err := rd.ReadLine()
		if isPrefix {
			continue
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		rec := &shakespeareRecord{}
		err = json.Unmarshal(line, &rec)
		if err != nil {
			if _, ok := err.(*json.UnmarshalTypeError); ok {
				continue
			}
			return err
		}

		recTitle := fmt.Sprintf("%s:%d:%s", rec.PlayName, rec.SpeechNumber, rec.Speaker)
		if title != recTitle {
			if err := create(title, strings.Join(text, "\n")); err != nil {
				return err
			}
			title = recTitle
			text = text[:0]

			time.Sleep(time.Second)
		}
		text = append(text, rec.Text)
	}
}
