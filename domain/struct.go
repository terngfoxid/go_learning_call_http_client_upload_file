package domain

type ResponseUploadFile struct {
	TransactionId string   `json:"transactionId" binding:"required"`
	Message       []string `json:"msg" binding:"required"`
	Code          []string `json:"code" binding:"required"`
	Filepath      []string `json:"filepath" binding:"required"`
	FileId        []int    `json:"fileId" binding:"required"`
}
