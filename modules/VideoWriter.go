package modules
import (
    "github.com/lazywei/go-opencv/opencv"
    "log"
)

type VideoWriter struct {
    outputDirectory string
    fourcc uint32
    fps float32
    frame_width, frame_height int
    is_color int

}
func (v *VideoWriter) SaveAsVideo (frames []*opencv.IplImage, name string) {
    if (len(frames) == 0 || &frames[0] == nil) {
        panic("no frames")
    }

    log.Print(int(v.fourcc), v.fps, v.frame_width,v.frame_height, v.is_color)
    var writter = opencv.NewVideoWriter(v.outputDirectory + name, int(v.fourcc), v.fps, v.frame_width,v.frame_height, v.is_color)
    defer writter.Release()

    log.Print("No. of frames: ", len(frames))
    for i := 0; i < len(frames);i++ {
        if (frames[i] == nil) {
            log.Print("img == nil")
        } else {
            writter.WriteFrame(frames[i])
        }
    }



}

func NewVideoWriter(fourCC uint32, fps float32, size opencv.Size) *VideoWriter {

    var writer = VideoWriter{"out/", fourCC, fps, size.Width, size.Height, 1}
    return &writer
}