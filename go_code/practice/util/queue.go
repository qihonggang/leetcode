package util

func queue() {
	// 创建队列
	queue := make([]int, 0)

	// enqueue入队
	queue = append(queue, 10)

	// dequeue出队
	v := queue[0]
	queue = queue[1:]

	// 长度0为空
	isEmpty := len(queue) == 0
}
