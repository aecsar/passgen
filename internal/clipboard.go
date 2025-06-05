package internal

import "golang.design/x/clipboard"

var err = clipboard.Init()

func WriteToClipboard(content string) {
	if err != nil {
		panic("Error while trying to write to clipboard")
	}

	clipboard.Write(clipboard.FmtText, []byte(content))
}
