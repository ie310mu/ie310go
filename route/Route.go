/*

注意，route模块依赖于session模块，且session.GlobalSessions必须要初始化

*/

package route

//此文件是为了配合代码生成工具，不需要在上层应用上另外定义

//Services 静态注册的服务，还需要根据配置文件来判断是否需要导出
var Services = []IService{}

//Add ..
func Add(s IService) {
	Services = append(Services, s)
}

//AdminServices ..
var AdminServices = []IService{}

//AddAdmin ..
func AddAdmin(s IService) {
	AdminServices = append(AdminServices, s)
}
