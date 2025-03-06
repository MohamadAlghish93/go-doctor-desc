package controllers

import (
	"bytes"
	"html/template"
	"net/http"
	"os"
	"os/exec"

	"doc-desc/database"
	"doc-desc/models"
	"github.com/gin-gonic/gin"
)

func AddReceipt(c *gin.Context) {
	var receipt models.Receipt
	if err := c.ShouldBindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&receipt)
	c.JSON(http.StatusOK, gin.H{"message": "Receipt created successfully", "receipt": receipt})
}

func GenerateReceiptPDF(c *gin.Context) {
	receiptID := c.Param("id")
	var receipt models.Receipt
	database.DB.Preload("Medicines").First(&receipt, receiptID)
	if receipt.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}
	tmpl, err := template.New("receipt").Parse(`
		<html>
		<head><style>
		body { font-family: Arial, sans-serif; }
		table { width: 100%; border-collapse: collapse; }
		th, td { border: 1px solid black; padding: 10px; text-align: left; }
		</style></head>
		<body>
		<h2>Medical Receipt</h2>
		<p>Patient: {{.PatientName}}</p>
		<p>Date: {{.Date}}</p>
		<h3>Medicines</h3>
		<table>
		<tr><th>Name</th><th>How to Use</th></tr>
		{{range .Medicines}}
		<tr><td>{{.MedicineName}}</td><td>{{.HowToUse}}</td></tr>
		{{end}}
		</table>
		</body>
		</html>
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Template parsing error"})
		return
	}

	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, receipt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Template execution error"})
		return
	}

	tempHTMLFile := "receipt.html"
	os.WriteFile(tempHTMLFile, htmlBuffer.Bytes(), 0644)
	tempPDFFile := "receipt.pdf"
	cmd := exec.Command("wkhtmltopdf", tempHTMLFile, tempPDFFile)
	err = cmd.Run()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate PDF"})
		return
	}

	c.File(tempPDFFile)
}
