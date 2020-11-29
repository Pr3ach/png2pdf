package main

import (
	"fmt"
	"os"
	_ "image/png"
	"github.com/jung-kurt/gofpdf"
	re "regexp"
	"sort"
	"strconv"
	"io/ioutil"
	"path/filepath"
)

const VERSION = "1.0.0"

type pngs []string

func main() {
	var dir string
	var l []string

	if len(os.Args) != 2 {
		usage(os.Args[0])
		os.Exit(-1)
	}

	dir = os.Args[1]

	if !dir_exists(dir) {
		fmt.Println("[!] No such directory")
		os.Exit(-1)
	}

	list_png(dir, &l)

	if len(l) < 1 {
		fmt.Printf("[!] No PNG files found in '%s'\n", dir)
		os.Exit(0)
	}

	if is_ordered(l) {
		sort.Sort(pngs(l))
	}

	if e := make_pdf_from_png(dir, l); e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}

	os.Exit(0)
}

/*
 * Display basic usage info.
 *
 */
func usage(self string) {
	fmt.Printf("png2pdf v%s\n\n", VERSION)
	fmt.Printf("[*] Usage: png2pdf <png_dir>\n")
}

/*
 * Merge PNGs into a single pdf file
 * FIXME: handle error
 *
 */
func make_pdf_from_png(dir string, png_list []string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	for i := 0; i < len(png_list); i++ {
		pdf.AddPage()
		fullpath := filepath.Join(dir, png_list[i])
		pdf.ImageOptions(fullpath, 0, 0, -1, -1, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "",)
	}

	pdf.OutputFileAndClose(filepath.Join(dir, "a.pdf"))

	return nil
}

/*
* Tell if a file exists.
* Return true/false
*
*/
func file_exists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

/*
* Tell if a directory exists.
* Return true/false
*
*/
func dir_exists(dir string) bool {
	info, err := os.Stat(dir)

	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

/*
 * Tell if the specified string array is only composed
 * of png files with digits as name
 * Return true/false
 *
 */
func is_ordered(l []string) bool {
	for i := 0; i < len(l); i++ {
		if  m, _ := re.MatchString("^[0-9]+\\.(?i)png$", l[i]); !m {
			return false;
		}
	}

	return true
}

/*
 * Custom implementation for the sort interface

 */
func (p pngs) Len() int {
	return len(p)
}

func (p pngs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p pngs) Less(i, j int) bool {
	// strip .png ext
	a, _ := strconv.Atoi(p[i][0:len(p[i])-4])
	b, _ := strconv.Atoi(p[j][0:len(p[j])-4])

	return a < b
}

/*
 * dir: directory
 * l: ptr to slice of string
 *
 * List all png files in specified dir
 *
 */
func list_png(dir string, l *[]string) error {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return err
	}

	for _, f := range files {
		fullpath := filepath.Join(dir, f.Name())

		if file_exists(fullpath) {
			if m, _:= re.MatchString(".*\\.(?i)png$", f.Name()); m {
				*l = append(*l, f.Name())
			}
		}
	}

	return nil
}
