package main

import (
   "os"
   "fmt"
   "io/ioutil"
   "gopkg.in/gographics/imagick.v3/imagick"
	"os/exec"
)
func main(){

   fileway:="/home/mky/Desktop/a.pdf"
   file, _ := os.Open(fileway)
   file.Close()
	ConvertToFile(GetFile())
	MultiTiff()
   fmt.Println("\n We are creating a new multi tiff format dokümant.It name's new.tiff !!! \n")
 }

func ConvertToFile (Fileway string) {

	imagick.Initialize()
	defer imagick.Terminate()

	aw:=imagick.NewMagickWand()
	mw :=imagick.NewMagickWand()

	defer mw.Destroy()
	defer aw.Destroy()

	aw.SetImageFormat("tiff")
	if err := mw.ReadImage(Fileway); err != nil {
		panic(err)
	}
	for i := 0; i < int(mw.GetNumberImages()); i++ {

		mw.SetIteratorIndex(i)
		tw := mw.GetImage()
		aw.AddImage(tw)
		aw.SetInterlaceScheme(aw.GetImageInterlaceScheme())
		aw.WriteImage("%d " + "test.tiff")
		tw.Destroy()
	}
}

func GetFile() string {

	var fileway string
	fmt.Println("Can you enter file document way")

	fmt.Scanf("%s", &fileway)

	if fileway=="" {
		fmt.Println("Documan way is empty!")
	}
	return fileway
}

func MultiTiff(){
	cmd := exec.Command("convert", "*.tiff", "new.tiff")
	cmd.CombinedOutput()
	cmd.Run()
}

func generateFaxFile(File *os.File, FaxId string)  {

    fileway:="/home/mky/Desktop/test_file.tiff"
    bytesread,_ := ioutil.ReadAll(File)
    fmt.Println("File:",File)
    f, _ := os.Create(fileway)
    n2, _ := f.Write(bytesread)
    fmt.Println("n2:",n2)
    defer f.Close()

}