package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	_Domain "go-back/domain"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CallAPIUploadFile(c *gin.Context) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	docTypeId := c.PostForm("docTypeId")
	groupId := c.PostForm("groupId")

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error cannot get multi part file from request")
		return
	}
	files := form.File["file"]

	url := "http://localhost:8080/upload-api/uploadFile/" + docTypeId + "/" + groupId

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	/*
	   	fw, err := writer.CreateFormField("name")
	       if err != nil {
	       }
	       _, err = io.Copy(fw, strings.NewReader("John"))
	       if err != nil {
	           return err
	       }
	*/

	for _, file := range files {
		if file != nil && file.Size > 0 {
			src, err := file.Open()
			if src != nil {
				defer func(src multipart.File) {
					err := src.Close()
					if err != nil {
						fmt.Println("Src Close Error.")
					}
				}(src)
			}
			if err != nil {
				fmt.Println("Error open src file process")
				c.JSON(http.StatusOK, "Error to open src file")
				return
			}

			fw, err := writer.CreateFormFile("file", file.Filename)
			if err != nil {
				c.JSON(http.StatusOK, "Error to write mutipart request")
				return
			}
			_, err = io.Copy(fw, src)
			if err != nil {
				c.JSON(http.StatusOK, "Error to copy src file to request")
				return
			}
		}
	}
	// Close multipart writer.
	writer.Close()
	req, err := http.NewRequest("POST", url, bytes.NewReader(body.Bytes()))
	if err != nil {
		c.JSON(http.StatusOK, "Error to create new request")
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, err := client.Do(req)
	if err != nil {
		c.JSON(rsp.StatusCode, "Error when client try to do request")
		return
	}
	defer rsp.Body.Close()

	finalRes := new(_Domain.ResponseUploadFile)

	json.NewDecoder(rsp.Body).Decode(finalRes)

	c.JSON(rsp.StatusCode, finalRes)
}
