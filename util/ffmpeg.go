package util

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"strings"
)

// GetImage 获取视频封面图并储存
func GetImage(finalName string) string {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input("public/video/"+finalName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 3)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	img, err := imaging.Decode(buf)
	if err != nil {
		fmt.Println(err)
	}
	// 去掉字符后缀
	imgName := strings.Split(finalName, ".")[0]
	err = imaging.Save(img, "./public/image/"+imgName+".jpeg")
	if err != nil {
		fmt.Println(err)
	}
	return imgName
}
