package main

import (
    "github.com/signintech/gopdf"
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
	
	pdf.AddPage()

	template := pdf.ImportPage("./assets/sample.pdf", 1, "/MediaBox")
	//　文字を書き込む
	pdf.UseImportedTemplate(template, 0, 0, size.W, size.H)
	pdf.SetX(8)
	pdf.SetY(13)
	pdf.Cell(nil, "iketa")

	pdf.AddPage()
	pdf.SetX(8)
	pdf.SetY(13)
	pdf.Cell(nil, "iketa?")


	// PDFをファイルに書き出す
	pdf.WritePdf("output/output.pdf")
}
