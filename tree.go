package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	all := flag.Bool("a", false, "Include all files")
	ftype := flag.Bool("F", false, "Show file type suffix")
	//other potential options:
	//-l follow symlinks, default is to not
	//-P include matches to pattern
	//-I exclude matches to pattern
	flag.Parse()
	args := flag.Args()
	var fileFilter func(os.FileInfo)bool
	if !*all {
		fileFilter = defaultFilter
	}
	filter := func(files []os.FileInfo) []os.FileInfo {
		if fileFilter == nil {
			return files
		}
		result := make([]os.FileInfo, 0)
		for _, info := range files {
			if !fileFilter(info) {
				result = append(result, info)
			}
		}
		return result
	}

	if len(args) > 0 {
		for _, dir := range args {
			tree(dir, filter, *ftype)
		}
	} else {
		tree(".", filter, *ftype)
	}
}

func tree(dir string, filter func([]os.FileInfo) []os.FileInfo, ftype bool) {
	fmt.Printf("%s\n", filepath.Base(dir))
	subtree(dir, filter, ftype, "")
}

func defaultFilter(info os.FileInfo) bool {
	name := info.Name()
	if strings.HasPrefix(name, ".") || strings.HasSuffix(name, "~") {
		return true
	}
	return false
}

func subtree(dir string, filter func([]os.FileInfo) []os.FileInfo, ftype bool, indent string) {
	infos, err := ioutil.ReadDir(dir)
	if err == nil {
		infos = filter(infos)
		i := 0
		last := len(infos) - 1
		spacer := "├──"
		for _, info := range infos {
			name := info.Name()
			if i == last {
				spacer = "└──"
			}
			i++
			if ftype {
				fmt.Printf("%s%s %s%s\n", indent, spacer, name, fileTypeSuffix(info))
			} else {
				fmt.Printf("%s%s %s\n", indent, spacer, name)
			}
			if info.IsDir() {
				subtree(filepath.Join(dir, name), filter, ftype, indent + "    ")
			}
		}
	}
}

func fileTypeSuffix(info os.FileInfo) string {
	switch info.Mode() & os.ModeType {
	case os.ModeDir:
		return "/"
	case os.ModeSymlink:
		return "@"
	case os.ModeSocket:
		return "="
	case os.ModeNamedPipe:
		return "|"
	}
	perms := info.Mode() & os.ModePerm
	if (perms & 0111) != 0 {
		return "*"
	}
	return ""
}
