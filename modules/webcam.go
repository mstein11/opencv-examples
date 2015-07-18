package modules

import (
    "fmt"
    "os"
    "github.com/lazywei/go-opencv/opencv"

    "log"
)

func ProcessImageTest() {

    log.Print("Try to open webcam.")
    cap := opencv.NewCameraCapture(0)
    if cap == nil {
        panic("can not open camera")
    }

    log.Print("webcame opened.")


    var outputFileName = "testvideo.mpeg"
    var width = 640
    var height = 480
    var fps = 20.0;
    var fourcc = opencv.FOURCC('m','p','4','v')
    var noOfCapturedFrames = 100

    cap.SetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH, float64(width))
    cap.SetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT, float64(height))
    cap.SetProperty(opencv.CV_CAP_PROP_FPS, fps)

    log.Print("capturing ", noOfCapturedFrames, " frames and save them to the following file: ", outputFileName)
    log.Print("width: ", width)
    log.Print("height: ", height)
    log.Print("fps: ", fps)
    log.Print("fourcc: ",fourcc)

    var frames []*opencv.IplImage

    for ;len(frames) < noOfCapturedFrames;{
        if cap.GrabFrame() {
            img := cap.RetrieveFrame(1)
            if img != nil {

                var frame = cap.RetrieveFrame(1)
                var copy = opencv.CreateImage(int(width), int(height), frame.Depth(), frame.Channels())
                defer copy.Release()
                opencv.Copy(frame, copy, nil)

                frames = append(frames,copy)

            } else {
                fmt.Println("Image ins nil")
            }
        }
    }

    cap.Release()
    win := opencv.NewWindow("Go-OpenCV Webcam")
    defer win.Destroy()
    for i := 0; i < len(frames);i++ {
        log.Print("displaying frame no: ", i)
        win.ShowImage(frames[i])
        key := opencv.WaitKey(10)
        if (key != 0) {
            log.Print(key)
        }
    }

    log.Print("start writing frames to file.",)
    var videoWritter = NewVideoWriter(fourcc,float32(fps), opencv.Size{int(width), int(height)})
    videoWritter.SaveAsVideo(frames, outputFileName)

    log.Print("created file with name: ",outputFileName)
    os.Exit(0)
}
