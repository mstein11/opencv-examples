package modules
import (
    "github.com/lazywei/go-opencv/opencv"
    "runtime"
    "path"
    "os"
    "log"
)
func MarkRoi() {

//
//    var fourcc = opencv.FOURCC('m','p','4','v')
//    var fps = float32(video.GetProperty(opencv.CV_CAP_PROP_FPS))
//    var size = opencv.Size{int(video.GetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH)),
//        int(video.GetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT))};
//
//    log.Print("fourcc: ", fourcc)
//    log.Print("fps: ", fps)
//    log.Print("size: ", size)
//
//    defer video.Release()

    win := opencv.NewWindow("Go-OpenCV Webcam")
    defer win.Destroy()

    win2 := opencv.NewWindow("rect cam: ")
    win2.Move(400,0)
    defer win2.Destroy()

    _, currentfile, _, _ := runtime.Caller(0)
    var cascade = path.Join(path.Dir(currentfile), "/../cascades/haarcascade_eye_tree_eyeglasses.xml")
    var eyeMarker = NewEyeMarker(cascade)

    var cap = NewVideoCapture(path.Join(path.Dir(currentfile), "/../eyevid1.mp4"))
    //var cap = NewVideoCapture("")

    var rois []*opencv.Rect

    for {
        var frame = cap.GetFrame()
        if (frame == nil) {
            break;
        }
            img := frame

                var copy = opencv.CreateImage(img.Width(),img.Height(),opencv.IPL_DEPTH_8U,3)
                opencv.Copy(img,copy,nil)

                    rois = eyeMarker.MarkEyesRects(copy)



                if (len(rois) > 0) {
                    log.Print(rois[0])
                    copy.SetROI(*rois[0])
                    copy = ProcessRoi(copy, 25)
                    //var gray = opencv.CreateImage(copy.Width(),copy.Height(),opencv.IPL_DEPTH_8U,1)
                    //copy.SetROI(*rois[0])


                    var circels = HoughCircles1(copy)
                    log.Print(circels)
                    for i := 0; i < len(circels) ;i++ {


                    var center = opencv.Point{rois[0].X() + int(circels[i].X),rois[0].Y() + int(circels[0].Y)}
                    //var center = opencv.Point{int(circels[i].X),int(circels[0].Y)}
                    opencv.Circle(img ,center, int(circels[i].R),  opencv.NewScalar(255, 0, 0, 1), 1, 1, 0)
                    //src *IplImage, dp float64, min_dist float64, param1 float64, param2 float64, min_radius int, max_radius int
                }
                    copy.ResetROI()
                }

                win.ShowImage(img)
        win2.ShowImage(copy)


        key := opencv.WaitKey(10)

        if key == 27 {
            os.Exit(0)
        }
    }

    os.Exit(0)
}