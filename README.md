## png2pdf
png2pdf is a simple utility written in Go that builds a single pdf file from given PNGs.

## Setup
`$ git clone http://github.com/Pr3ach/png2pdf` <br/>
`$ cd png2pdf` <br/>
`$ go build` <br/>


## Usage
`./png2pdf <png_directory>`<br/><br/>
png2pdf merges PNGs in a `a.pdf` file in the specified directory

**Note**: png2pdf will order PNGs by name in the PDF document if those are all numbered
