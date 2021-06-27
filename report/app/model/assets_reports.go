// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"assets/app/model/internal"
)

// AssetsReports is the golang structure for table assets_reports.
type AssetsReports internal.AssetsReports

// Fill with you ideas below.

// ResponseAssetsWebReport 渗透测试报告 模糊分页查询返回数据所需信息
type ResponseAssetsWebReport struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []*AssetsReports `json:"data"`
}

// ResponseAssetsReportLevelNameGroup 渗透测试报告漏洞类型Group分组
type ResponseAssetsReportLevelNameGroup struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []ResponseAssetsReportLevelNameGroupInfo `json:"data"`
}

// ResponseAssetsReportLevelNameGroupInfo 渗透测试报告漏洞类型Group分组详情
type ResponseAssetsReportLevelNameGroupInfo struct{
	ID int `json:"id"`
	LevelName string `json:"level_name"`
}

// ResponseAssetsReport 渗透测试报告 模糊分页查询返回数据所需信息
type ResponseAssetsReport struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []*ResponseAssetsReportInfo `json:"data"`
}

// ResponseAssetsReportInfo 渗透测试报告 模糊分页查询返回数据所需详细信息
type ResponseAssetsReportInfo struct{
	ID1 uint `json:"id1"`
	*AssetsReports
}

// RequestReport 漏洞报告修改
type RequestReport struct {
	LevelName string `v:"required#漏洞名称不能为空" json:"level_name"`
	Level string `v:"required|integer|min:1|max:3#漏洞等级不能空|漏洞等级只能为整数|漏洞等级参数值错误,最小为1|漏洞等级参数值错误,最大为3" json:"level"`
	LevelStatus string `v:"required|integer|min:1|max:3#漏洞整改状态不能空|漏洞整改状态只能为整数|漏洞整改状态参数值错误,最小为1|漏洞整改状态参数值错误,最大为3" json:"level_status"`
	ReportId string `v:"required|integer#漏洞ID不能空#漏洞ID只能为整数"`
}


// ResponseTongjiInfo  主页统计数据
type ResponseTongjiInfo struct{
	WebCount int `json:"web_count"` // web业务资系统数量
	LevelCount int `json:"level_count"` // 漏洞总数
	LevelYesCount int `json:"level_yes_count"` // 已整改漏洞数
	LevelNoCount int `json:"level_no_count"` // 未整改数
}

// ResponseEchartsInfo Echarts统计数据返回所需信息
type ResponseEchartsInfo struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []UResponseEchartsInfoManagerLevel `json:"data"`
	Data1 []UResponseEchartsInfoLevel `json:"data1"`
	Data2 []UResponseEchartsInfoAssetsName `json:"data2"`
}

// UResponseEchartsInfoManagerLevel Echarts统计数据返回所需data信息 安全管理人员漏洞情况
type UResponseEchartsInfoManagerLevel struct{
	ManagerName string `json:"name"`
	Number int `json:"value"`
}

// UResponseEchartsInfoLevel Echarts统计数据返回所需data信息 漏洞类型情况
type UResponseEchartsInfoLevel struct{
	LevelName string `json:"name"`
	Number int `json:"value"`
}

// UResponseEchartsInfoAssetsName Echarts统计数据返回所需data信息 漏洞类型情况
type UResponseEchartsInfoAssetsName struct{
	AssetsName string `json:"name"`
	Number int `json:"value"`
}
