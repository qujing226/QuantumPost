# QuantumPost - 企业级异步消息管理平台

![QuantumPost Logo](https://via.placeholder.com/150x50.png?text=QuantumPost+)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)
![License](https://img.shields.io/badge/license-Apache--2.0-blue)

**下一代智能消息中枢** | SMS/Email双通道管理 | 企业级消息中台解决方案

[English](./README_EN.md) | 简体中文

## 📌 核心特性

### 消息管理
- **双通道智能路由**  
  ▸ 基于业务场景的SMS/Email自动分流策略  
  ▸ 通道健康度实时监控与熔断机制

### 技术亮点
- **异步消息引擎**
  ```mermaid
  graph LR
  A[API请求] --> B[Gin接收]
  B --> C{Kafka}
  C --> D[邮件处理器]
  C --> E[SMS处理器]
  D --> F[(MySQL)]
  E --> F
可视化追踪
▸ 消息全生命周期状态追踪
▸ 失败重试DAG工作流

## 🛠️ 技术栈
```
层级	技术组件	版本	用途
前端	Vue3 + TypeScript	3.2+	管理界面开发
网关层	Gin	1.24+	      RESTful API接口服务
数据持久化	GORM + MySQL	2.0+	业务数据存储
异步队列	Kafka 	3.4+	      消息解耦与流量削峰
基础设施	Docker	20.10+	       容器化部署
```
## 🚀 快速开始
前置需求
Go 1.20+

Node.js 16+

MySQL 8.0+

Kafka 3.4+

### 克隆仓库
``` bash
git clone https://github.com/qujing226/quantumpost.git
```

### 前端依赖安装
``` bash
cd vue && npm install
```

### 后端依赖安装
``` bash
cd server && go mod tidy
```

### 邮件服务配置
SMTP_HOST=smtp.example.com
SMTP_PORT=587

## 🧩 项目结构
```quantumpost/
├── backend/               # Gin服务核心
│   ├── cmd/               # 启动入口
│   ├── internal/          
│   │   ├── delivery/      # HTTP接口层
│   │   ├── service/       # 业务逻辑层
│   │   └── repository/    # 数据访问层（GORM实现）
├── frontend/              # Vue3管理界面
│   ├── src/
│   │   ├── views/         # 页面组件
│   │   └── stores/        # Pinia状态管理
└── deploy/                # K8s部署文件
```
## 📡 消息处理流程
API接收层
Gin路由接收请求，进行基础验证后投递到Kafka

### 异步处理器

``` go
func ProcessEmail(msg *sarama.ConsumerMessage) {
   var task EmailTask
   json.Unmarshal(msg.Value, &task)
   // 重试逻辑实现
}
```
### 状态同步
处理结果实时推送至管理界面

## 🧑💻 开发者指南
贡献流程
Fork本项目仓库

创建特性分支 (git checkout -b feature/amazing-feature)

提交变更 (git commit -m 'Add some amazing feature')

推送分支 (git push origin feature/amazing-feature)

发起Pull Request


## 启动集成测试环境
docker-compose -f docker-compose.test.yml up
## 📜 开源协议
本项目采用 Apache License 2.0 授权

架构演进路线
2023 Q4 - 支持企业微信消息通道
2024 Q1 - 集成GPT-4智能消息生成引擎


