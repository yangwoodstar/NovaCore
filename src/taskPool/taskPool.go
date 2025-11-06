package taskPool

import (
	"sync"

	"github.com/yangwoodstar/NovaCore/src/tools"
	"go.uber.org/zap"
)

type WorkerPool struct {
	poolSize   int
	taskName   []string
	tasks      map[string]chan Task // 每个 hash 值对应一个任务通道
	wg         sync.WaitGroup
	consistent *ConsistentHash
	handler    func(task Task) error
}

func NewWorkerPool(poolSize, queueSize int, taskName []string, handler func(task Task) error) *WorkerPool {
	tasks := make(map[string]chan Task, poolSize)
	for _, value := range taskName {
		tasks[value] = make(chan Task, queueSize)
	}

	consistentIns := New(poolSize, MurmurHash)
	consistentIns.Add(taskName)
	tools.Logger.Info("taskName", zap.Any("taskName", taskName))
	return &WorkerPool{
		poolSize:   poolSize,
		tasks:      tasks,
		taskName:   taskName,
		consistent: consistentIns,
		handler:    handler,
	}
}

func (wp *WorkerPool) Start() {
	for index, value := range wp.taskName {
		wp.wg.Add(1)
		tools.Logger.Info("start", zap.Int("index", index), zap.String("value", value))
		go wp.worker(value)
	}
}

func (wp *WorkerPool) worker(workerID string) {
	tools.Logger.Debug("task", zap.String("workerID", workerID))
	defer wp.wg.Done()
	for task := range wp.tasks[workerID] {
		//tools.Logger.Debug("task", zap.String("workerID", workerID), zap.String("task", string(task.Data)))
		err := wp.handler(task)
		if err != nil {
			ackErr := task.Data.Ack()
			if ackErr != nil {
				tools.Logger.Error("ack error", zap.Error(ackErr), zap.String("msg", string(task.Data.GetBody())))
			}
			tools.Logger.Error("Failed to process record message", zap.Error(err))
			continue
		}
		ackErr := task.Data.Ack()
		if ackErr != nil {
			tools.Logger.Error("ack error", zap.Error(ackErr), zap.String("msg", string(task.Data.GetBody())))
		}
	}
}

func (wp *WorkerPool) AddTask(task Task, hashKey string) {
	key := wp.consistent.Get(hashKey)
	//tools.Logger.Info("msg", zap.String("key", key), zap.String("hashKey", hashKey))
	wp.tasks[key] <- task
}

func (wp *WorkerPool) Wait() {
	for _, ch := range wp.tasks {
		close(ch)
	}
	wp.wg.Wait()
}
