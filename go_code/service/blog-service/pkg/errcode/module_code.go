package errcode
var (
	ErrorGetTagListFail = NewError(20001001, "获取标签列表失败")
	ErrorCreateTagFail = NewError(20001002, "创建标签失败")
	ErrorUpdateTagFail = NewError(20001003, "更新标签失败")
	ErrorDeleteTagFail = NewError(20001004, "删除标签失败")
	ErrorCountTagFail = NewError(20001005, "统计标签失败")
)
