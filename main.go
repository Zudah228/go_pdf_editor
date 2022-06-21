package main

import (
    "github.com/signintech/gopdf"
		"io"
		"net/http"
		"os"
)

func main() {
	// サイズ設定
	size :=* &gopdf.Rect{W: 123, H: 181}


	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: size})

	// フォントの取り込み
	err := pdf.AddTTFFont("ipaexm", "./assets/ipaexm.ttf")
	if err != nil {
			panic(err)
	}
	// フォントサイズを選択
	err = pdf.SetFont("ipaexm", "", 6)
	if err != nil {
			panic(err)
	}
	
	// ページの追加
	pdf.AddPage()

	// PDF を取り込んで、その上に描画する
	template := pdf.ImportPage("./assets/sample.pdf", 1, "/MediaBox")
	pdf.UseImportedTemplate(template, 0, 0, size.W, size.H)

	//　文字を書き込む
	pdf.SetX(8)
	pdf.SetY(13)
	pdf.Cell(nil, "iketa")

	// PDF のダウンロード
	var downloadPath = "./download/example-pdf.pdf"
	fileUrl := "https://tcpdf.org/files/examples/example_012.pdf"
	if err = DownloadFile(downloadPath, fileUrl); err != nil {
			panic(err)
	}

	pdf.AddPage()
	template2 := pdf.ImportPage(downloadPath, 2, "/MediaBox")
	pdf.UseImportedTemplate(template2, 0, 0, size.W, size.H)



	// PDFをファイルに書き出す
	pdf.WritePdf("output/output.pdf")
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
			return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
			return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}