package modules
import (
    "github.com/lazywei/go-opencv/opencv"
    "log"
    "image"
    "image/color"
)


func CaptureFrames() {
    var cam = opencv.NewCameraCapture(0)
    defer cam.Release()

    frames := []*opencv.IplImage{}
    //var writter = opencv.NewVideoWriter("test.avi", int(opencv.FOURCC('D','I','V','3')),20, 50, 50, 1)

    win := opencv.NewWindow("Go-OpenCV Webcam123")
defer win.Destroy()
    for framesCounter:=0;framesCounter <= 25;framesCounter++{
        log.Print(framesCounter)
        if (cam.GrabFrame() || framesCounter > 25) {
//            var frame = cam.RetrieveFrame(1)
//            var img = ProcessImage(frame, 50)
//            frames = append(frames, img)



            var frame = cam.RetrieveFrame(1)
            win.ShowImage(frame)
//            var img3 = frame.ToImage()


            var imageHeader = opencv.CreateImage(frame.Width(),frame.Height(), frame.Depth(), frame.Channels())

            for i := 0; i < frame.Width(); i++ {
                for j := 0; j < frame.Height(); j++ {

                    scalar := frame.Get2D(i,j)
                    var bgra = scalar.Val()
                    var newscalar = opencv.NewScalar(bgra[0],bgra[1],bgra[2],1)
                    imageHeader.Set2D(i,j, newscalar)

                }
            }
//            var img3Opencv = opencv.FromImage(img3)

            var copy = opencv.CreateImage(frame.Width(),frame.Height(),frame.Depth(),frame.Channels())
            opencv.Copy(frame,copy,nil)
//            frames = append(frames, img3Opencv)

            //win.ShowImage(imageHeader)


        }
    }

    var videoWritter = NewVideoWriter(opencv.FOURCC('D','I','V','3'), 20, opencv.Size{frames[0].Width(), frames[0].Height()})
    videoWritter.SaveAsVideo(frames, "test.avi")




}

func ProcessImage(img *opencv.IplImage, gradeOfCanny int) *opencv.IplImage {
    w := img.Width()
    h := img.Height()

    // Create the output image
    cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)
    //defer cedge.Release()

    // Convert to grayscale
    gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
    edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
    //newimage := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
    defer gray.Release()
    defer edge.Release()

    opencv.CvtColor(img, gray, opencv.CV_BGR2GRAY)

    opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
    opencv.Not(gray, edge)

    // Run the edge detector on grayscale
    opencv.Canny(gray, edge, float64(gradeOfCanny), float64(gradeOfCanny*3), 3)

    opencv.Zero(cedge)
    // copy edge points
    opencv.Copy(img, cedge, edge)

    return cedge
}

func ProcessRoi(img *opencv.IplImage, gradeOfCanny int) *opencv.IplImage {

    var roi = img.GetROI()




    w := roi.Width()
    h := roi.Height()

    // Create the output image
    cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)
    //defer cedge.Release()

    // Convert to grayscale
    gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
    edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
    //newimage := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
    defer gray.Release()
    defer edge.Release()

    opencv.CvtColor(img, gray, opencv.CV_BGR2GRAY)

    opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
    opencv.Not(gray, edge)

    // Run the edge detector on grayscale
    opencv.Canny(gray, edge, float64(gradeOfCanny), float64(gradeOfCanny*3), 3)

    opencv.Zero(cedge)
    // copy edge points
    opencv.Copy(img, cedge, edge)

    return cedge




}

func HoughCircles1(img *opencv.IplImage) []opencv.CircleStruct{
    var gray = opencv.CreateImage(img.Width(),img.Height(),opencv.IPL_DEPTH_8U,1)
    opencv.CvtColor(img, gray, opencv.CV_BGR2GRAY);
    return opencv.HoughCircles(gray, 15, 100,75,5,25,500)
}

func HoughCirclesWithParams(img *opencv.IplImage, dp float64, min_dist float64, param_1 float64, param_2 float64, min_radius int, maxradius int) []opencv.CircleStruct {
    var gray = opencv.CreateImage(img.Width(),img.Height(),opencv.IPL_DEPTH_8U,1)
    opencv.CvtColor(img, gray, opencv.CV_BGR2GRAY);
    return opencv.HoughCircles(gray, dp,min_dist, param_1,param_2,1,500)
}


func ToImage(img *opencv.IplImage) image.Image {
    out := image.NewNRGBA(image.Rect(0, 0, img.Width(), img.Height()))
    if img.Depth() != opencv.IPL_DEPTH_8U {
        return nil // TODO return error
    }

    for y := 0; y < img.Height(); y++ {
        for x := 0; x < img.Width(); x++ {
            s := img.Get2D(x, y).Val()
            b, g, r, a := s[2], s[1], s[0], s[3]

            c := color.NRGBA{uint8(b), uint8(g), uint8(r), uint8(a)}
            out.Set(x, y, c)
        }
    }

    return out
}