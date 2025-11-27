package tools

import "fmt"

func GetExchangeBindingQueueUrl(domain, port, exchange, vhost string) string {
	return fmt.Sprintf("http://%s:%s/api/exchanges/%s/%s/bindings/source", domain, port, vhost, exchange)
}

func GetQueueInfoUrl(domain, port, queue, vhost string) string {
	return fmt.Sprintf("http://%s:%s/api/queues/%s/%s", domain, port, vhost, queue)
}

func GetAllQueuesUrl(domain, port, vhost string) string {
	return fmt.Sprintf("http://%s:%s/api/queues", domain, port, vhost)
}
