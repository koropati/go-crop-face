package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "opencv/highgui"
    "opencv/objdetect"
)

func main() {
    // Load the Haar cascade classifier.
    faceCascade := objdetect.NewHaarCascadeClassifier("haarcascade_frontalface_alt2.xml")

    // Open the image file.
    img, err := highgui.ReadImage("image.jpg")
    if err != nil {
        fmt.Println(err)
        return
    }

    // Convert the image to grayscale.
    gray := image.NewGray(img.Bounds())
    draw.Draw(gray, gray.Bounds(), img, image.ZP, draw.Src)

    // Detect faces in the image.
    faces := faceCascade.DetectObjects(gray)

    // Crop the faces from the image.
    for _, face := range faces {
        rect := face.Rect()
        faceImage := image.NewRGBA(rect)
        draw.Draw(faceImage, faceImage.Bounds(), img, rect, draw.Src)

        // Save the cropped face image.
        highgui.WriteImage("face.jpg", faceImage)
    }
}
