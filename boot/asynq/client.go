package asynq

import (
	"github.com/hibiken/asynq"
	"github.com/small-ek/antgo/os/alog"
	"github.com/small-ek/antgo/os/config"
	"github.com/small-ek/antgo/utils/conv"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	Client *asynq.Client
	once   sync.Once
)

const (
	TypeInstitutionInfo         = "match:institution_info"               //匹配机构
	TypeLoanSupermarketProducts = "statistics:loan_supermarket_products" //统计贷超产品
	TypeChannelStatistics       = "statistics:channel_statistics"        //渠道统计
)

// Register 初始化队列客户端，确保只初始化一次
func Register() {
	once.Do(func() {
		addr := config.GetString("redis.0.address")
		password := config.GetString("redis.0.password")
		if addr == "" {
			alog.Write.Error("Redis address is not configured")
			return
		}

		// 创建 Asynq 客户端
		Client = asynq.NewClient(asynq.RedisClientOpt{
			Addr:     addr,
			Password: password,
			DB:       1,
		})

		// 检测连接
		if err := Client.Ping(); err != nil {
			alog.Write.Error("Failed to connect to Redis", zap.Error(err))
		} else {
			alog.Write.Info("Successfully connected to Redis")
		}
	})
}

// PushTask 创建新任务并加入队列
func PushTask(taskType string, taskData map[string]interface{}, delay time.Duration) (*asynq.TaskInfo, error) {
	// 将任务数据编码为 JSON
	data, err := conv.ToJSON(taskData)
	if err != nil {
		alog.Write.Error("Failed to encode payload to JSON", zap.Error(err))
		return nil, err
	}

	// 创建任务
	task := asynq.NewTask(taskType, data)
	options := []asynq.Option{
		asynq.MaxRetry(5),                      // 最大重试次数
		asynq.ProcessAt(time.Now().Add(delay)), // 延迟执行
		//asynq.Queue(queueName),                 // 指定队列名称

	}

	// 将任务加入队列
	info, err := Client.Enqueue(task, options...)
	if err != nil {
		alog.Write.Error("Failed to enqueue task", zap.String("task_type", taskType), zap.Error(err))
		return nil, err
	}

	alog.Write.Info("Task enqueued successfully", zap.String("task_id", info.ID), zap.String("queue", info.Queue))
	return info, nil
}

// CloseClient 关闭客户端连接（需要在应用退出时调用）
func CloseClient() error {
	if err := Client.Close(); err != nil {
		alog.Write.Error("Failed to close Asynq client", zap.Error(err))
		return err
	}
	alog.Write.Info("Asynq client closed successfully")
	return nil
}
