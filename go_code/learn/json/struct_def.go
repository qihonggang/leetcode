package json

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type JobInfo struct {
	Skills []string `json:"skills"`
}
type Employee struct {
	Basicinfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
