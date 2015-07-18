package modules
import (
    "os"
    "fmt"
    "runtime"
    "path"
    "log"
    "local/imageProcessing/go-opencv/opencv"
)

func PrintVideosMetadata() {
    filename := "../data/???.avi"
    if len(os.Args) == 2 {
        _, currentfile, _, _ := runtime.Caller(0)

        filename = os.Args[1]
        filename = path.Join(path.Dir(currentfile), filename)
    } else {
        fmt.Printf("Usage: go run player.go videoname\n")
        os.Exit(0)
    }

    cap := opencv.NewFileCapture(filename)
    if cap == nil {
        panic("can not open video")
    }
    defer cap.Release()


    log.Print(" ")
    log.Print("Getting Metadata from video: ", filename)
    log.Print("Metadata start")
    log.Print("fourcc: ",cap.GetProperty(opencv.CV_CAP_PROP_FOURCC))
    log.Print("fps: ",cap.GetProperty(opencv.CV_CAP_PROP_FPS))
    log.Print("brightness: ",cap.GetProperty(opencv.CV_CAP_PROP_BRIGHTNESS))
    log.Print("contrast: ",cap.GetProperty(opencv.CV_CAP_PROP_CONTRAST))
    log.Print("convet_RGB: ",cap.GetProperty(opencv.CV_CAP_PROP_CONVERT_RGB))
    log.Print("exposure: ",cap.GetProperty(opencv.CV_CAP_PROP_EXPOSURE))
    log.Print("format: ",cap.GetProperty(opencv.CV_CAP_PROP_FORMAT))
    log.Print("framecount: ",cap.GetProperty(opencv.CV_CAP_PROP_FRAME_COUNT))
    log.Print("frameHeight: ",cap.GetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT))
    log.Print("frameWidth: ",cap.GetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH))
    log.Print("Gain: ",cap.GetProperty(opencv.CV_CAP_PROP_GAIN))
    log.Print("Hue: ",cap.GetProperty(opencv.CV_CAP_PROP_HUE))
    log.Print("Mode: ",cap.GetProperty(opencv.CV_CAP_PROP_MODE))
    log.Print("PosAviRatio: ",cap.GetProperty(opencv.CV_CAP_PROP_POS_AVI_RATIO))
    log.Print("PosFrames: ",cap.GetProperty(opencv.CV_CAP_PROP_POS_FRAMES))
    log.Print("PosMsec: ",cap.GetProperty(opencv.CV_CAP_PROP_POS_MSEC))
    log.Print("Rectification: ",cap.GetProperty(opencv.CV_CAP_PROP_RECTIFICATION))
    log.Print("Saturation: ",cap.GetProperty(opencv.CV_CAP_PROP_SATURATION))
    log.Print("Metadata end")
    os.Exit(0)
}