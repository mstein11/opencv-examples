package main
import (        "./modules"
    "path"
    "os"

    "github.com/lazywei/go-opencv/opencv"
)

const TEST_FILE = "eyevid1.mp4"

func main () {




    //modules.ProcessImageTest()

    //markRois()
    modules.PrintVideosMetadata()

    os.Exit(0)

    //modules.TrackEyes(src, src);



}

func markRois() {
    modules.MarkRoi()
}

func videoProcessor() {
    var src = path.Join(path.Dir(TEST_FILE)) + "/" + TEST_FILE

    var cap = modules.NewVideoCapture(src)
    var processor = modules.NewVideoProcessor(cap, modules.NewVideoWriter(uint32(opencv.FOURCC('m','p','4','v')),10.0, cap.Size))

    processor.TestInit()
    processor.Run()
}