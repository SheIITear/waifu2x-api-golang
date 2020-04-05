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

func form(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, `<!DOCTYPE html>
	<html>
		<!-- By ShellTear -->
			<title>waifu upgrader</title>
				<style>
					body {
						background-image: url('https://images.wallpapersden.com/image/download/blonde-blue-eye-anime-girl_65536_2560x1440.jpg');
						background-attachment: fixed;
						background-size: cover;
						background-position: center center;
						background-repeat: no-repeat;
					}
				</style>
					<center>
						<h1><color style="color:blue">By ShellTear. Webp files have high chance of failing, so please avoid using them.</color></h1>
							<form action="/convert" method="post" enctype="multipart/form-data">
							  <input type="file" name="file" accept="image/*">
	 							 <input type="submit">
							</form>
					</center>
					</body>
	</html>
	`)
}

func checkSubstrings(str string, subs ...string) (bool, int) {

	matches := 0
	isCompleteMatch := true

	gg.Blue.Printf("security check, comparing: \"%s\" to: %s\n", str, subs)

	for _, sub := range subs {

		if strings.Contains(str, sub) {
			matches++
		} else {
			isCompleteMatch = false
		}
	}

	return isCompleteMatch, matches
}

func convert(w http.ResponseWriter, r *http.Request) {
	var maxSize int64 = 4
	var sizembs = r.ContentLength / 1024 / 1024

	if r.ContentLength <= 0 {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Only fixed-size files are allowed")
		return
	}

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
	ext := filepath.Ext(header.Filename)
	inFile := fmt.Sprintf("/data/%v_%v", rand.Int63(), header.Filename) // TODO: fix security
	isCompleteMatch1, matches1 := checkSubstrings(ext, ".jpeg", ".webp", ".jpg", ".png")
	outFile := fmt.Sprintf("/data/%v_out.png", rand.Int63())

	if matches1 == 1 {
		gg.Green.Println("filetype accepted\n" + "filename: " + header.Filename)
		gg.Blue.Println("complete match:", isCompleteMatch1)

		// if size is smaller than 1mb
		if sizembs == 0 {
			gg.Blue.Println("file size: <", sizembs, "mb")
		} else {
			gg.Blue.Println("file size:", sizembs, "mb")
		}

		gg.Blue.Println("enhancing... (this might take a while depending on the file size)")
		err = ioutil.WriteFile(inFile, data, 7777)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "write file: %v", err)
			return
		}

	} else {
		gg.Red.Println("filetype not accepted\n" + "filename: " + header.Filename)

		// if size is smaller than 1mb
		if sizembs == 0 {
			gg.Blue.Println("file size: <", sizembs, "mb")
		} else {
			gg.Blue.Println("file size:", sizembs, "mb")
		}

		w.WriteHeader(400)
		fmt.Fprintf(w, "%v <- not an image file", header.Filename)
		return
	}

	cmd := exec.Command("/opt/waifu2x-cpp/waifu2x-converter-cpp", fmt.Sprintf("-i %v", inFile), fmt.Sprintf("-o %v", outFile))
	out, err := cmd.CombinedOutput()

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "program err: %v %v", err, string(out))
		return
	}

	// log.Print("waifu: ", string(out))
	gg.Green.Println("successfully enhanced -> " + header.Filename)
	gg.Blue.Println("converting to jpg and uploading...")
	data, err = ioutil.ReadFile(outFile)

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "read file: %v", err)
		return
	}

	imgSrc, err := png.Decode(bytes.NewBuffer(data))

	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "png decode: %v", err)
		return
	}

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

func main() {
	rand.Seed(time.Now().Unix())
	r := mux.NewRouter()
	r.HandleFunc("/", form)
	r.HandleFunc("/convert", convert)
	print("\033[H\033[2J")
	gg.Blue.Println("waifus waiting on port 33457")
	err := http.ListenAndServe(":"+"33457", r)
	if err != nil {
		log.Fatal("listen: ", err)
	}
}
