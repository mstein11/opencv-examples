package modules
import (

    "log"
    "github.com/lazywei/go-opencv/opencv"
    "path"
    "runtime"
)




type VideoProcessor struct {
    source VideoCapture
    dest *VideoWriter
    processors []Process
    currentFrameIn *opencv.IplImage
    framesOut []*opencv.IplImage
}

func NewVideoProcessor(capture VideoCapture, dest *VideoWriter) VideoProcessor{
    return VideoProcessor{capture, dest, []Process{}, nil,[]*opencv.IplImage{}}
}

func (v *VideoProcessor) Run () {

    for {
        frame := v.source.GetFrame()
        v.currentFrameIn = frame
        if (frame == nil) {
            break;
        }
        for i := 0; i < len(v.processors); i++ {
            v.processors[i](v)
        }
    }

    v.dest.SaveAsVideo(v.framesOut, "testout.mov")

}

func (v *VideoProcessor) TestInit() {


    //v.processors = append(v.processors, processor1)
    v.processors = append(v.processors, MarkEyesCirclesProcessor)

}

var MarkEyesRectProcessor = func(v *VideoProcessor) {

}

var MarkEyesCirclesProcessor = func (v *VideoProcessor) {
    _, currentfile, _, _ := runtime.Caller(0)
    var cascade = path.Join(path.Dir(currentfile), "/../cascades/haarcascade_eye_tree_eyeglasses.xml")
    var eyeMarker = NewEyeMarker(cascade)

    var copy = opencv.CreateImage(v.currentFrameIn.Width(),v.currentFrameIn.Height(),v.currentFrameIn.Depth(), v.currentFrameIn.Channels())
    var copy2 = opencv.CreateImage(v.currentFrameIn.Width(),v.currentFrameIn.Height(),v.currentFrameIn.Depth(), v.currentFrameIn.Channels())
    opencv.Copy(v.currentFrameIn, copy, nil)
    opencv.Copy(v.currentFrameIn, copy2, nil)

    var rois = eyeMarker.MarkEyesRects(copy)

    if (len(rois) > 0) {
        copy.SetROI(*rois[0])
        copy = ProcessRoi(copy, 75)
        //var gray = opencv.CreateImage(copy.Width(),copy.Height(),opencv.IPL_DEPTH_8U,1)
        //copy.SetROI(*rois[0])


        var circels = HoughCircles1(copy)
        log.Print(circels)
        for i := 0; i < len(circels) ;i++ {


            var center = opencv.Point{rois[0].X() + int(circels[i].X),rois[0].Y() + int(circels[0].Y)}
            //var center = opencv.Point{int(circels[i].X),int(circels[0].Y)}
            opencv.Circle(copy2 ,center, int(circels[i].R),  opencv.NewScalar(255, 0, 0, 1), 1, 1, 0)
            //src *IplImage, dp float64, min_dist float64, param1 float64, param2 float64, min_radius int, max_radius int
        }
        copy.ResetROI()
    }
    v.framesOut = append(v.framesOut, copy2)
}


type Process func (*VideoProcessor)