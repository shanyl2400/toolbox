package main

import (
	"toolbox/internal/config"
	"toolbox/internal/logs"
	"toolbox/internal/repository"
	"toolbox/internal/repository/boltdb"
	"toolbox/internal/repository/model"
	"toolbox/internal/utils"

	"go.uber.org/zap"
)

var (
	users = []*model.User{
		{
			UserName:     "admin",
			PasswordHash: utils.Hash("admin"),
		},
	}

	columns = []*model.Column{
		{
			Name: "信控中心",
		},
		{
			Name: "GOMSS",
		},
		{
			Name: "其他",
		},
	}

	tools = []*model.Tool{
		//研发环境
		{
			Name:        "Grafana",
			URL:         "http://10.0.2.247:3000/?orgId=1",
			Profile:     "grafana.jpg",
			Description: "gomsg系统监控工具",

			Environment: "研发环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "gomsg-debug",
			URL:         "http://10.0.10.104:8081/",
			Profile:     "gomsg_debug.png",
			Description: "gomsg研发环境debug页面",

			Environment: "研发环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "Prometheus",
			URL:         "http://10.0.2.247:9090/graph?g0.expr=&g0.tab=1&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h",
			Profile:     "prometheus.jpg",
			Description: "gomsg系统监控数据收集",

			Environment: "研发环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "GOMSS发版工具",
			URL:         "http://10.0.2.247:8080/publisher/",
			Profile:     "publisher.png",
			Description: "GOMSS与zrtc编译及镜像打包工具",

			Environment: "研发环境",
			ColumnName:  "GOMSS",
		},
		{
			Name:        "WebRTC page",
			URL:         "http://10.0.2.247:8080/webrtc_page/index.html",
			Profile:     "webrtc.png",
			Description: "WebRTC测试页",

			Environment: "研发环境",
			ColumnName:  "其他",
		},

		//预发环境
		{
			Name:        "Grafana",
			URL:         "http://10.0.2.247:3000/?orgId=1",
			Profile:     "grafana.jpg",
			Description: "gomsg系统监控工具",

			Environment: "预发环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "gomsg-debug",
			URL:         "http://10.0.2.62:8081/",
			Profile:     "gomsg_debug.png",
			Description: "gomsg预发环境debug页面",

			Environment: "预发环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "Prometheus",
			URL:         "http://10.0.2.247:9090/graph?g0.expr=&g0.tab=1&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h",
			Profile:     "prometheus.jpg",
			Description: "gomsg系统监控数据收集",

			Environment: "预发环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "GOMSS发版工具",
			URL:         "http://10.0.2.247:8080/publisher/",
			Profile:     "publisher.png",
			Description: "GOMSS与zrtc编译及镜像打包工具",

			Environment: "预发环境",
			ColumnName:  "GOMSS",
		},
		{
			Name:        "WebRTC page",
			URL:         "http://10.0.2.247:8080/webrtc_page/index.html",
			Profile:     "webrtc.png",
			Description: "WebRTC测试页",

			Environment: "预发环境",
			ColumnName:  "其他",
		},

		//生产环境
		{
			Name:        "Grafana",
			URL:         "http://10.0.2.247:3000/?orgId=1",
			Profile:     "grafana.jpg",
			Description: "gomsg系统监控工具",

			Environment: "生产环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "gomsg-debug",
			URL:         "http://172.20.20.161:8081/",
			Profile:     "gomsg_debug.png",
			Description: "gomsg生产环境debug页面",

			Environment: "生产环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "Prometheus",
			URL:         "http://10.0.2.247:9090/graph?g0.expr=&g0.tab=1&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h",
			Profile:     "prometheus.jpg",
			Description: "gomsg系统监控数据收集",

			Environment: "生产环境",
			ColumnName:  "信控中心",
		},
		{
			Name:        "GOMSS发版工具",
			URL:         "http://10.0.2.247:8080/publisher/",
			Profile:     "publisher.png",
			Description: "GOMSS与zrtc编译及镜像打包工具",

			Environment: "生产环境",
			ColumnName:  "GOMSS",
		},
		{
			Name:        "WebRTC page",
			URL:         "http://10.0.2.247:8080/webrtc_page/index.html",
			Profile:     "webrtc.png",
			Description: "WebRTC测试页",

			Environment: "生产环境",
			ColumnName:  "其他",
		},
	}
)

func main() {
	config.Set(&config.Config{
		BoltDBPath: "./data.db",
	})

	err := boltdb.GetClient().Open()
	if err != nil {
		panic(err)
	}
	defer boltdb.GetClient().Close()

	// create users
	usersRepo := repository.GetUsersRepository()
	for _, user := range users {
		err := usersRepo.Set(user)
		if err != nil {
			logs.Error("create users repository failed",
				zap.Error(err))
		}
	}

	// create columns
	columnRepo := repository.GetColumnsRepository()
	for _, column := range columns {
		err := columnRepo.Set(column)
		if err != nil {
			logs.Error("create column repository failed",
				zap.Error(err))
		}
	}

	// create tools
	toolsRepo := repository.GetToolsRepository()
	for _, tool := range tools {
		err := toolsRepo.Set(tool)
		if err != nil {
			logs.Error("create tools repository failed",
				zap.Error(err))
		}
	}

	logs.Info("init data success")
}
