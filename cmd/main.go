package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"convert"
)

var (
	srcFormat string
	dstFormat string
)

func init() {
	flag.StringVar(&srcFormat, "srcFormat", ".jpg", "ディレクトリ指定の場合に、変換する画像ファイルのフォーマット（`png|jpeg|jpg`）")
	flag.StringVar(&dstFormat, "dstFormat", ".png", "ディレクトリ指定の場合に、変換する画像ファイルのフォーマット（`png|jpeg|jpg`）")
	flag.Parse()
}

func run() error {
	args := flag.Args()
	if len(args) > 1 {
		return fmt.Errorf("入力が多すぎます")
	}
	if f, err := os.Stat(args[0]); os.IsNotExist(err) || !f.IsDir() {
		return fmt.Errorf("ディレクトリは存在しません！")
	}

	err := filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == srcFormat {
			srcAbs, err := filepath.Abs(path)
			if err != nil {
				return fmt.Errorf("ファイルを読み込めませんでした")
			}
			dst := path[:len(path)-len(srcFormat)] + dstFormat //拡張子変換
			dstAbs, err := filepath.Abs(dst)
			if err != nil {
				return fmt.Errorf("ファイルを読み込めませんでした")
			}
			//変換処理
			//第１引数:出力する画像ファイル,第２引数:読み込みたい画像ファイル
			convert.Do(dstAbs, srcAbs)
			//元ファイル削除
			if err := os.Remove(srcAbs); err != nil {
				fmt.Println(err)
			}
			return nil
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
