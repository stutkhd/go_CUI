//Package convert is the special one
package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

//MyImage is image type
type MyImage struct {
	image.Image
}

//Do is Change Ext
func Do(dst, src string) error {
	//読み込み用ファイル
	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("画像ファイルが開けませんでした。%s", src)
	}
	defer sf.Close()

	//書き込み用ファイル
	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("画像ファイルを書き出せませんでした。%s", dst)
	}
	defer df.Close()

	_img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	img := &MyImage{_img} //埋め込む

	//dstが出力したい画像ファイル
	switch strings.ToLower(filepath.Ext(dst)) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	}

	if err != nil { //Encodeで返ってくるError
		return fmt.Errorf("画像ファイルを書き出せませんでした。 %s", dst)
	}

	return nil
}
