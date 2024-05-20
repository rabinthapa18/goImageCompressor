package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/disintegration/imaging"
)

func main() {
	// Create a new local server
	fmt.Println("Starting the server on localhost:3000")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Read the size to compress the image
		size, err := strconv.Atoi(r.URL.Query().Get("size"))
		if err != nil {
			fmt.Println(err)
			fmt.Println("Error in reading size")
			return
		}

		// Read the image file
		file, _, err := r.FormFile("image")
		if err != nil {
			fmt.Println("Error in reading image file")
			return
		}

		// Create a buffer to store the image
		var imageFile []byte
		buf := new(bytes.Buffer)
		io.Copy(buf, file)
		imageFile = buf.Bytes()

		// Compress the image
		compressedImage, err := compressImage(imageFile, size)
		if err != nil {
			fmt.Println("Error in compressing image")
			return
		}

		// Write the compressed image to the response
		w.Header().Set("Content-Type", "image/jpeg")
		w.Header().Set("Content-Length", fmt.Sprint(len(compressedImage)))
		w.Write(compressedImage)
	})

	// Start the local server
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error in starting server", err)
	}
}

func compressImage(file []byte, size int) ([]byte, error) {
	// Decode the image
	img, err := imaging.Decode(bytes.NewReader(file))
	if err != nil {
		fmt.Println("Error in decoding image")
		return nil, fmt.Errorf("error in decoding image")
	}

	// Resize the image to width provided size keeping the aspect ratio
	img = imaging.Resize(img, size, 0, imaging.Lanczos)

	// Encode the image to JPEG
	buf := new(bytes.Buffer)
	err = imaging.Encode(buf, img, imaging.JPEG)
	if err != nil {
		fmt.Println("Error in encoding image")
		return nil, fmt.Errorf("error in encoding image")
	}

	return buf.Bytes(), nil
}
