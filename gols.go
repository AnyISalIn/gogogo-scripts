package main

import (
	"io/ioutil"
	"os"
	"syscall"
	"strings"
	"github.com/urfave/cli"
	"fmt"
)

func getFileList(dir string) ([]os.FileInfo) {
	files, _ := ioutil.ReadDir(dir)
	return files
}

func getFileInfo(file os.FileInfo) (map[string]interface{}) {
	s, ok := file.Sys().(*syscall.Stat_t)
	m := map[string]interface{}{
		"name": file.Name(),
		"size": file.Size(),
		"mode": file.Mode()}
	if ok {
		m["nlink"] = uint16(s.Nlink)
		m["uid"] = uint32(s.Uid)
		m["gid"] = uint32(s.Gid)
	}
	return m
}

func getDirList(files []os.FileInfo) []os.FileInfo {
	results := []os.FileInfo{}
	for _, file := range files {
		if file.IsDir() {
			results = append(results, file)
		}
	}
	return results
}

func filter(files []os.FileInfo, c *cli.Context) []os.FileInfo {
	results := []os.FileInfo{}
	if c.Bool("dir") {
		files = getDirList(files)
	}
	if c.Bool("all") {
		return files
	}
	for _, file := range files {
		if ! strings.HasPrefix(file.Name(), ".") {
			results = append(results, file)
		}
	}
	return results
}

func parse(c *cli.Context) {
	dirs := c.Args()
	for _, dir := range dirs {
		files := getFileList(dir)
		fmt.Printf("%v:\n", dir)
		for _, file := range filter(files, c) {
			fileinfo := getFileInfo(file)
			if c.Bool("list") {
				fmt.Printf("%-2v  %-2v  %-2v  %-2v  %-2v  %-2v\n",
					fileinfo["mode"], fileinfo["nlink"], fileinfo["uid"], fileinfo["gid"],
					fileinfo["size"], fileinfo["name"])
			} else {
				fmt.Println(fileinfo["name"])
			}
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "gols"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name: "l, list",
			Usage: "list format",
		},
		cli.BoolFlag{
			Name: "a, all",
			Usage: "list all file",
		},
		cli.BoolFlag{
			Name: "d, dir",
			Usage: "list dir",
		},
	}
	app.Action = func(c *cli.Context) error {
		parse(c)
		return nil
	}
	app.Run(os.Args)
}
