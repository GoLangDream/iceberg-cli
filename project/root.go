package project

// Root
// 返回项目的根目录，目前直接返回相对路径，之后还需要添加对目录结构的检查,看是否符合iceberg的目录结构
// 如果不是，就递归检查父目录，如果所有父目录都不是，那就返回错
func Root() string {
	return ""
}
