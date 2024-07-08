package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func handleFoundFile(ext, fileName string) error {
if ext == "mp4" {
	return handleVideoFiles(fileName)
}

return nil

}

func handleVideoFiles(fileName string) error {
			cmdToCopy := fmt.Sprintf("mv ./%s ./Videos", fileName)
			 err := exec.Command("/bin/sh", "-c", cmdToCopy).Run()
			 if err != nil {
				return err
			 }
			 return nil

}

func main (){
	// get user home directory
	home , err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	downloadsDir := fmt.Sprintf("%s/Downloads", home)
	// switch to downloads folder
	err = os.Chdir(downloadsDir)
	if err != nil {
		panic(err)
	}
	// get a list of all files in the downloads dir
	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		info , err := os.Stat(file.Name())
		if err != nil {
		fmt.Printf("\nThe error is: %+v",err)
			}
			brkdwn := strings.SplitAfter(info.Name(), ".")
			fileExtension :=  brkdwn[len(brkdwn)-1]
			handleFoundFile(fileExtension, info.Name())
	// fmt.Printf("\n%+v", file.Info())
	}
}
