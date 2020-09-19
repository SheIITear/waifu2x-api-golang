package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	gg "github.com/gookit/color"

	"github.com/gorilla/mux"
)

// check the string aka files
func checkSubstrings(str string, subs ...string) int {

	matches := 0
	gg.Blue.Printf("security check, comparing: \""+str+"\" to: %s\n", subs)

	for _, sub := range subs {

		if strings.Contains(str, sub) {
			matches++
		}
	}

	return matches
}

// do the actual job
func convert(w http.ResponseWriter, r *http.Request) {

	// set max size and get the size of uploaded file
	var maxSize int64 = 4
	var sizembs = r.ContentLength / 1024 / 1024

	// check that the size is smaller than provided on upper variable
	if r.ContentLength <= 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Only fixed-size files are allowed")
		return
	}

	// if size is bigger than allowed
	if r.ContentLength > maxSize*1024*1024 {
		gg.Blue.Println("file size is too big:", sizembs)
		w.WriteHeader(400)
		fmt.Fprintf(w, "file size is > %vmb", maxSize)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxSize*1024*1024)
	err := r.ParseMultipartForm(maxSize * 1024 * 1024)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "multipart read: %v", err)
		return
	}

	file, header, err := r.FormFile("file")

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "%v <- 400 bad request, did you forget to select a file before trying to upload?", err)
		return
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "read: %v", err)
		return
	}

	// check that the file extension is allowed
	ext := filepath.Ext(header.Filename)
	matches1 := checkSubstrings(ext, ".jpeg", ".webp", ".jpg", ".png")

	// get file size
	if sizembs == 0 {
		gg.Blue.Println("file size: <", sizembs, "mb")
	} else {
		gg.Blue.Println("file size:", sizembs, "mb")
	}

	// if no matches or too many, don't allow the file
	if matches1 != 1 {

		gg.Red.Println("filetype not accepted\n" + "filename: " + header.Filename)

		w.WriteHeader(400)
		fmt.Fprintf(w, "%v <- not an image file", header.Filename)
		return
	}

	gg.Green.Println("filetype accepted\n" + "filename: " + header.Filename)

	// create files
	inFile := fmt.Sprintf("/data/%v_%v", rand.Int63(), header.Filename)
	outFile := fmt.Sprintf("/data/%v_out.png", rand.Int63())

	gg.Blue.Println("enhancing... (this might take a while depending on the file size)")

	// write given file to server
	err = ioutil.WriteFile(inFile, data, 7777)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "write file: %v", err)
		return
	}

	// run the upscaling command
	command := "/opt/waifu2x-cpp/waifu2x-converter-cpp" + " -i " + inFile + " -o " + outFile
	cmdString := strings.TrimSuffix(command, "\n")
	cmdString2 := strings.Fields(cmdString)
	cmdOutput, err := exec.Command(cmdString2[0], cmdString2[1:]...).Output()

	// get output
	realout := string(cmdOutput[:])
	realout = strings.TrimSuffix(realout, "\n")

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "program err: %v %v", err, realout)
		return
	}

	// print result
	gg.Green.Println("successfully enhanced -> " + header.Filename)
	gg.Blue.Println("converting to jpg and uploading...")
	data, err = ioutil.ReadFile(outFile)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "read file: %v", err)
		return
	}

	// decode png
	imgSrc, err := png.Decode(bytes.NewBuffer(data))

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "png decode: %v", err)
		return
	}

	// create jpg from upscaled (save space etc)
	newImg := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(newImg, newImg.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(newImg, newImg.Bounds(), imgSrc, imgSrc.Bounds().Min, draw.Over)
	var opt jpeg.Options
	opt.Quality = 95

	var extension = filepath.Ext(header.Filename)
	var name = header.Filename[0 : len(header.Filename)-len(extension)]
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%v_2x.jpg"`, name))
	err = jpeg.Encode(w, newImg, &opt)

	if err != nil {
		log.Printf("jpeg encode err: %v", err)
	}

	gg.Green.Println("successfully converted and uploaded -> " + name + "_2x.jpg")
}

// start the server
func main() {
	rand.Seed(time.Now().Unix())
	r := mux.NewRouter()

	// main
	fs := http.FileServer(http.Dir("./frontend"))
	r.Handle("/", fs)

	// converting
	r.HandleFunc("/convert", convert)

	// clear terminal
	print("\033[H\033[2J")

	// start listening
	gg.Blue.Println("waifus waiting on port 33457")
	err := http.ListenAndServe(":"+"33457", r)

	if err != nil {
		log.Fatal("listen: ", err)
	}
}
