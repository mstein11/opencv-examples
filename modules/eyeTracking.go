package modules
import (
    "github.com/lazywei/go-opencv/opencv"
    "fmt"
    "os"
    "log"
    "path"
    "runtime"
//    "local/imageProcessing/go-opencv/opencv"
)


func TrackEyes(srcFile, destFile string) {
    var video = opencv.NewFileCapture(srcFile)


    var fourcc = opencv.FOURCC('m','p','4','v')
    var fps = float32(video.GetProperty(opencv.CV_CAP_PROP_FPS))
    var size = opencv.Size{int(video.GetProperty(opencv.CV_CAP_PROP_FRAME_WIDTH)),
                    int(video.GetProperty(opencv.CV_CAP_PROP_FRAME_HEIGHT))};



    if video == nil {
        panic("no file found with name: " + srcFile)
    }
    defer video.Release()

    win := opencv.NewWindow("Go-OpenCV Webcam")
    defer win.Destroy()

    _, currentfile, _, _ := runtime.Caller(0)
    var cascade = path.Join(path.Dir(currentfile), "/../cascades/haarcascade_eye_tree_eyeglasses.xml")
    var eyeMarker = NewEyeMarker(cascade)

    var imgs = []*opencv.IplImage{}
    var imgs2 = []*opencv.IplImage{}

    for {
        if video.GrabFrame() {
            img := video.RetrieveFrame(1)
            if img != nil {

                var copy = opencv.CreateImage(img.Width(),img.Height(),img.Depth(),img.Channels())
                opencv.Copy(img.Clone(),copy,nil)

                var markedImg = eyeMarker.MarkEyes(copy)

                imgs = append(imgs, markedImg)
                //win.ShowImage(imgs[len(imgs)-1])
            } else {
                fmt.Println("Image ins nil")
            }
        } else {
            break
        }
        key := opencv.WaitKey(10)

        if key == 27 {
            os.Exit(0)
        }
    }

    for i := 0; i < len(imgs2);i++ {
        win.ShowImage(imgs2[i])
    }
    log.Print("before write")
    NewVideoWriter(uint32(fourcc), fps, size).SaveAsVideo(imgs, "test.avi")
    os.Exit(0)
}

func NewEyeMarker(classifierPath string) *eyeMarker {
    //_, currentfile, _, _ := runtime.Caller(0)
    var classifierFile = opencv.LoadHaarClassifierCascade(classifierPath)

    if (classifierFile == nil) {
        panic("Classifier must not be nil")
    }
    return &eyeMarker{classifierFile, 0, []int{},[]int{}}
}



type eyeMarker struct {
    classifier *opencv.HaarCascade
    counter int
    cachedPointsCenterY, cachedPointsCenterX []int
}
func (e *eyeMarker) MarkEyes(img *opencv.IplImage) *opencv.IplImage{

    if (e.counter == 1) {
        //find eye, cache them and then mark them
        e.counter =0;
        e.cachedPointsCenterY = []int{}
        e.cachedPointsCenterX = []int{}

        rects := e.classifier.DetectObjects(img)


        for _, value := range rects {

            var pointRect1 = opencv.Point{value.X() + value.Width(), value.Y()}
            var pointRect2 = opencv.Point{value.X(), value.Y() + value.Height()}
            var centerPointX = int((pointRect1.X + pointRect2.X) / 2)
            var centerPointY = int((pointRect1.Y + pointRect2.Y) / 2)

            var center = opencv.Point{centerPointX, centerPointY}
            e.cachedPointsCenterX = append(e.cachedPointsCenterX,center.X)
            e.cachedPointsCenterY = append(e.cachedPointsCenterY, center.Y)


            //draw eye marker
            e.drawEyeMarker(img, center)
        }
    } else {
        //if we have a cached value
        for i := 0; i < len(e.cachedPointsCenterY); i++ {
            e.drawEyeMarker(img, opencv.Point{e.cachedPointsCenterX[i], e.cachedPointsCenterY[i]})
        }

    }
    //counter for cache
    e.counter = e.counter+ 1;

    return img
}

//find eyes in image. return rect as found area.
func (e *eyeMarker) MarkEyesRects(img *opencv.IplImage) []*opencv.Rect {
    rects := e.classifier.DetectObjects(img)
    return rects
}

func (e eyeMarker) drawEyeMarker(img *opencv.IplImage, p opencv.Point) {
    opencv.Circle(img, p, 5, opencv.ScalarAll(175.0), 1, 1, 0)
}

