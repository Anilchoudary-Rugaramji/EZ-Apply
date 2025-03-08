package parser

// import (
// 	"bytes"
// 	"fmt"

// 	"rsc.io/pdf"
// )

// type ResumeMetaData struct {
// 	Name             string
// 	Email            string
// 	PhoneNumber      int
// 	Desgination      string
// 	Experience       int
// 	HighestEducation string
// 	Location         string
// 	Skiils           []string
// }

// func ExtractResumeMetaDataFromPDF(resume []byte) (*ResumeMetaData, error) {
// 	// Open the PDF
// 	reader, err := pdf.NewReader(bytes.NewReader(resume), int64(len(resume)))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to open PDF: %v", err)
// 	}

// 	// Extract text from all pages
// 	var text string
// 	numPages := reader.NumPage()
// 	for i := 1; i <= numPages; i++ {
// 		page := reader.Page(i)
// 		if page.V.IsNull() {
// 			continue
// 		}
// 		content, err := page.(nil)
// 		if err == nil {
// 			text += content + "\n"
// 		}
// 	}

// 	// Parse extracted text into structured data
// 	parsedResume := parseResume(text)

// 	return &parsedResume, nil
// }
