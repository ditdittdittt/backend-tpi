package services

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"

	"github.com/ditdittdittt/backend-tpi/constant"
)

func GeneratePdf(header map[string]interface{}, data []map[string]interface{}, pdfType string) ([]byte, error) {
	var filepath string

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	header["table"] = data
	header["downloaded_at"] = time.Now().Format("2 Jan 2006 15:04:05")

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

	outputFile := "./output.pdf"
	switch pdfType {
	case constant.ProductionPdf:
		outputFile = "./output-production.pdf"
	case constant.TransactionPdf:
		outputFile = "./output-transaction.pdf"
	}

	err = pdfg.WriteFile(outputFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Take a pdf file
	outputPDF, err := os.OpenFile(outputFile, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Read a file
	pdf, err := ioutil.ReadAll(outputPDF)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	//// Remove a file
	//err = os.Remove(outputFile)
	//if err != nil {
	//	log.Fatal(err)
	//	return nil, err
	//}
	return pdf, nil
}
