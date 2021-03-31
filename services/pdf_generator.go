package services

import (
	"bytes"
	"html/template"
	"log"
	"path"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/ditdittdittt/backend-tpi/constant"
)

func GeneratePdf(header map[string]interface{}, data []map[string]interface{}, pdfType string) (*wkhtmltopdf.PDFGenerator, error) {
	var filepath string

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	header["table"] = data

	switch pdfType {
	case constant.ProductionPdf:
		filepath = path.Join("template", "production_report.html")
	case constant.TransactionPdf:
		filepath = path.Join("template", "transaction_report.html")
	}

	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, header)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(&tpl))

	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//outputFile := "./output.pdf"
	//
	//err = pdfg.WriteFile(outputFile)
	//if err != nil {
	//	log.Fatal(err)
	//	return nil, err
	//}

	return pdfg, nil
}
