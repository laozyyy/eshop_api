package handler

import (
	"eshop_api/common/constant"
	"eshop_api/log"
	"eshop_api/model/resp"
	"github.com/gin-gonic/gin"
	"github.com/go-zookeeper/zk"
	"net/http"
	"strconv"
	"time"
)

func HandleGetTags(ctx *gin.Context) {
	// 定义配置结构体，这里假设配置包含服务器地址、端口和超时时间，可根据实际需求修改
	type Config struct {
		ID   string `json:"tag_id"`
		Name string `json:"tag_name"`
	}
	// Zookeeper服务器地址列表，可配置多个地址以实现高可用，这里示例只写一个
	hosts := []string{"117.72.72.114:2181"}
	// 配置数据在Zookeeper中的节点路径，需替换为实际存储配置的路径
	configNodePath := "/frontend/menu"

	// 创建Zookeeper连接
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		log.Errorf("error: %v", err)
		ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
	}
	defer conn.Close()

	// 获取配置节点的数据
	data, _, err := conn.Get(configNodePath)
	if err != nil {
		log.Errorf("error: %v", err)
		ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
	}
	// 将获取到的数据（字节切片类型）解析为Config结构体
	var config []*Config
	//err = sonic.Unmarshal(data, &config)
	config = []*Config{
		{"T73SNA", "电子产品"},
		{"j8yi6g", "书籍教材"},
		{"4GYSCA", "运动器材"},
		{"FCMIjQ", "生活用品"},
		{"SUF3zA", "服装鞋帽"},
		{"DNc25A", "文具用品"},
		{"nvzRAw", "交通工具"},
		{"_u3XMw", "数码配件"},
	}
	if err != nil {
		log.Errorf("error: %v", err)
		ctx.JSON(http.StatusInternalServerError, "服务器内部错误")
	}
	tags := make([]resp.Tag, 0)
	for _, t := range config {
		tmp := resp.Tag{
			ID:   t.ID,
			Name: t.Name,
		}
		tags = append(tags, tmp)
	}
	res := resp.TagRespDTO{
		Code:   strconv.Itoa(constant.Success),
		Msg:    "success",
		Result: tags,
	}
	ctx.JSON(http.StatusOK, res)
}
