package main

import (
	"fmt"
	"os"
	"strings"
	"errors"
	_ "image/png"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	var png_list []string

	if len(os.Args) < 2 {
		usage(os.Args[0])
		os.Exit(-1)
	}

	png_list = os.Args[1:]

	if e := validate_args(png_list); e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}

	if e := make_pdf_from_png(png_list); e != nil {
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
	fmt.Printf("[!] Usage: png2pdf <png1> <png2> ...\n")
}

/*
 * Check if all files are actually .png && file exist
 * Return nil/error
 *
 */
func validate_args(png_list []string) error {
	for i := 0; i < len(png_list); i ++ {
		if !strings.HasSuffix(strings.ToLower(png_list[i]), ".png") {
			return errors.New("[-] File list must contain .png images only")
		} else if !file_exists(png_list[i]) {
			return errors.New(fmt.Sprintf("[-] %s: No such file", png_list[i]))
		}
	}

	return nil
}

/*
 * Merge PNGs into a single pdf file
 * FIXME: handle error
 *
 */
func make_pdf_from_png(png_list []string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")

	for i := 0; i < len(png_list); i++ {
		pdf.AddPage()
		pdf.ImageOptions(png_list[i], 0, 0, -1, -1, false, gofpdf.ImageOptions{ImageType: "PNG", ReadDpi: true}, 0, "",)
	}

	pdf.OutputFileAndClose("a.pdf")

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
