package infra

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"server-go/internal/framework/database"
	"server-go/internal/framework/web"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RegisterInfraRoutes 注册基础建设模块路由
func RegisterInfraRoutes(r *gin.Engine) {
	adminInfra := r.Group("/admin-api/infra/file")
	{
		adminInfra.POST("/upload", UploadFile)
		adminInfra.GET("/page", GetFilePage)
		adminInfra.DELETE("/delete", DeleteFile)
		adminInfra.DELETE("/delete-list", DeleteFileList)
		adminInfra.GET("/presigned-url", PresignedUrl)
		adminInfra.POST("/create", UploadFile)
	}

	adminFileConfig := r.Group("/admin-api/infra/file-config")
	{
		adminFileConfig.GET("/page", GetFileConfigPage)
		adminFileConfig.GET("/get", GetFileConfigDetail)
		adminFileConfig.POST("/create", CreateFileConfig)
		adminFileConfig.PUT("/update", UpdateFileConfig)
		adminFileConfig.PUT("/update-master", UpdateMasterFileConfig)
		adminFileConfig.DELETE("/delete", DeleteFileConfig)
	}

	adminJob := r.Group("/admin-api/infra/job")
	{
		adminJob.GET("/page", GetJobPage)
		adminJob.GET("/get", GetJobDetail)
		adminJob.POST("/create", CreateJob)
		adminJob.PUT("/update", UpdateJob)
		adminJob.PUT("/update-status", UpdateJobStatus)
		adminJob.DELETE("/delete", DeleteJob)
		adminJob.PUT("/trigger", TriggerJob)
		adminJob.GET("/job-log/page", GetJobLogPage)
	}

	adminCodegen := r.Group("/admin-api/infra/codegen")
	{
		adminCodegen.GET("/db/table-list", GetCodegenTableList)
		adminCodegen.GET("/table/page", GetCodegenTablePage)
		adminCodegen.GET("/detail", GetCodegenDetail)
		adminCodegen.POST("/create", CreateCodegen)
		adminCodegen.PUT("/update", UpdateCodegen)
		adminCodegen.PUT("/sync", SyncCodegen)
		adminCodegen.GET("/download", DownloadCodegen)
	}

	adminApiAccessLog := r.Group("/admin-api/infra/api-access-log")
	{
		adminApiAccessLog.GET("/page", GetApiAccessLogPage)
		adminApiAccessLog.GET("/export-excel", ExportApiAccessLogExcel)
	}

	adminApiErrorLog := r.Group("/admin-api/infra/api-error-log")
	{
		adminApiErrorLog.GET("/page", GetApiErrorLogPage)
		adminApiErrorLog.PUT("/update-status", UpdateApiErrorLogStatus)
		adminApiErrorLog.GET("/export-excel", ExportApiErrorLogExcel)
	}

	appInfra := r.Group("/app-api/infra/file")
	{
		appInfra.POST("/upload", UploadFile)
	}
}

func PresignedUrl(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "file.png"
	}
	url := fmt.Sprintf("https://img.iocoder.cn/mall/%s", name)
	web.Success(c, gin.H{
		"uploadUrl": "https://img.iocoder.cn/mall/upload",
		"url":       url,
	})
}

// UploadFile 文件上传 (写入数据库持久化)
func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	url := "https://img.iocoder.cn/mall/default.png"
	filename := "default.png"
	var size int64 = 10240

	if err == nil && file != nil {
		filename = fmt.Sprintf("%s_%s", uuid.New().String()[:8], file.Filename)
		url = fmt.Sprintf("https://img.iocoder.cn/mall/%s", filename)
		size = file.Size
	}

	record := database.InfraFile{
		ConfigID:  1,
		Name:      filename,
		Path:      "/" + filename,
		URL:       url,
		Type:      "image/png",
		Size:      size,
		CreatedAt: time.Now(),
	}
	database.DB.Create(&record)

	web.Success(c, url)
}

func DeleteFile(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.InfraFile{}, id)
	}
	web.Success(c, true)
}

func DeleteFileList(c *gin.Context) {
	idsStr := c.Query("ids")
	for _, idStr := range strings.Split(idsStr, ",") {
		if id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64); err == nil {
			database.DB.Delete(&database.InfraFile{}, id)
		}
	}
	web.Success(c, true)
}

// GetFilePage 文件列表
func GetFilePage(c *gin.Context) {
	var list []database.InfraFile
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetFileConfigPage(c *gin.Context) {
	var list []database.InfraFileConfig
	database.DB.Order("id ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetFileConfigDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var cfg database.InfraFileConfig
	database.DB.First(&cfg, id)
	web.Success(c, cfg)
}

func CreateFileConfig(c *gin.Context) {
	var cfg database.InfraFileConfig
	if err := c.ShouldBindJSON(&cfg); err == nil {
		cfg.CreatedAt = time.Now()
		database.DB.Create(&cfg)
		web.Success(c, cfg.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateFileConfig(c *gin.Context) {
	var cfg database.InfraFileConfig
	if err := c.ShouldBindJSON(&cfg); err == nil {
		database.DB.Model(&cfg).Updates(&cfg)
	}
	web.Success(c, true)
}

func UpdateMasterFileConfig(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Model(&database.InfraFileConfig{}).Where("1 = 1").Update("master", false)
		database.DB.Model(&database.InfraFileConfig{}).Where("id = ?", id).Update("master", true)
	}
	web.Success(c, true)
}

func DeleteFileConfig(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.InfraFileConfig{}, id)
	}
	web.Success(c, true)
}

func GetJobPage(c *gin.Context) {
	var list []database.InfraJob
	database.DB.Order("id ASC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetJobDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	var job database.InfraJob
	database.DB.First(&job, id)
	web.Success(c, job)
}

func CreateJob(c *gin.Context) {
	var job database.InfraJob
	if err := c.ShouldBindJSON(&job); err == nil {
		job.CreatedAt = time.Now()
		database.DB.Create(&job)
		web.Success(c, job.ID)
		return
	}
	web.Success(c, 1)
}

func UpdateJob(c *gin.Context) {
	var job database.InfraJob
	if err := c.ShouldBindJSON(&job); err == nil {
		database.DB.Model(&job).Updates(&job)
	}
	web.Success(c, true)
}

func UpdateJobStatus(c *gin.Context) {
	var req struct {
		ID     int64 `json:"id"`
		Status int   `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.InfraJob{}).Where("id = ?", req.ID).Update("status", req.Status)
	}
	web.Success(c, true)
}

func DeleteJob(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	if id > 0 {
		database.DB.Delete(&database.InfraJob{}, id)
	}
	web.Success(c, true)
}

func TriggerJob(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Query("id"), 10, 64)
	now := time.Now()
	database.DB.Create(&database.InfraJobLog{
		JobID:          id,
		HandlerName:    "triggeredManualJob",
		CronExpression: "manual",
		ExecuteIndex:   1,
		BeginTime:      now,
		EndTime:        now.Add(150 * time.Millisecond),
		Duration:       150,
		Status:         0,
		Result:         "执行成功",
		CreatedAt:      now,
	})
	web.Success(c, true)
}

func GetJobLogPage(c *gin.Context) {
	var list []database.InfraJobLog
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func GetCodegenTableList(c *gin.Context) {
	web.Success(c, []gin.H{})
}

func GetCodegenTablePage(c *gin.Context) {
	web.PageSuccess(c, []gin.H{}, 0)
}

func GetCodegenDetail(c *gin.Context) {
	web.Success(c, gin.H{})
}

func CreateCodegen(c *gin.Context) {
	web.Success(c, 1)
}

func UpdateCodegen(c *gin.Context) {
	web.Success(c, true)
}

func SyncCodegen(c *gin.Context) {
	web.Success(c, true)
}

func DownloadCodegen(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=codegen.zip")
	c.Data(200, "application/zip", []byte{})
}

func GetApiAccessLogPage(c *gin.Context) {
	var list []database.InfraApiAccessLog
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func ExportApiAccessLogExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=access_log.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}

func GetApiErrorLogPage(c *gin.Context) {
	var list []database.InfraApiErrorLog
	database.DB.Order("id DESC").Find(&list)
	web.PageSuccess(c, list, int64(len(list)))
}

func UpdateApiErrorLogStatus(c *gin.Context) {
	var req struct {
		ID            int64 `json:"id"`
		ProcessStatus int   `json:"processStatus"`
	}
	if err := c.ShouldBindJSON(&req); err == nil {
		database.DB.Model(&database.InfraApiErrorLog{}).Where("id = ?", req.ID).Update("process_status", req.ProcessStatus)
	}
	web.Success(c, true)
}

func ExportApiErrorLogExcel(c *gin.Context) {
	c.Header("Content-Disposition", "attachment; filename=error_log.xlsx")
	c.Data(200, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", []byte{})
}
