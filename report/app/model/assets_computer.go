// ==========================================================================
// This is auto-generated by gf cli tool. Fill this file as you wish.
// ==========================================================================

package model

import (
	"assets/app/model/internal"
)

// AssetsComputer is the golang structure for table assets_computer.
type AssetsComputer internal.AssetsComputer

// Fill with you ideas below.

// ResponseAssetsPcDepartmentGroup 终端资产 返回部门组信息
type ResponseAssetsPcDepartmentGroup struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []ResponseAssetsPcDepartmentGroupInfo `json:"data"`
}

// ResponseAssetsPcDepartmentGroupInfo 终端资产 返回部门详细信息
type ResponseAssetsPcDepartmentGroupInfo struct{
	ID int `json:"id"`
	Department string `json:"department"`
}

// ResponseAssetsPcDepartmentSubGroup 终端资产 返回二级部门组信息
type ResponseAssetsPcDepartmentSubGroup struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []ResponseAssetsPcDepartmentSubGroupInfo `json:"data"`
}

// ResponseAssetsPcDepartmentSubGroupInfo 终端资产 返回二级部门详细信息
type ResponseAssetsPcDepartmentSubGroupInfo struct{
	ID int `json:"id"`
	DepartmentSub string `json:"department_sub"`
}

// ResponseAssetsPc 终端资产管理 模糊分页查询返回数据所需信息
type ResponseAssetsPc struct{
	Code int `json:"code"`
	Msg string `json:"msg"`
	Count int64 `json:"count"`
	Data []*ResponseAssetsPcInfo `json:"data"`
}

// ResponseAssetsPcInfo 终端资产管理 模糊分页查询返回数据所需详细信息
type ResponseAssetsPcInfo struct{
	ID1 uint `json:"id1"`
	*AssetsComputer
}

// RequestAssetsPcAdd 添加终端资产所需信息
type RequestAssetsPcAdd struct{
	FlagId string
	Department    string
	DepartmentSub string
	PersonName    string `v:"required#员工姓名不能为空"`
	WorkNumber    string
	ComputerType  string
	ComputerName  string
	Address       string
	SecretLevel   string
	InternetFlag  string
	FileCopyFlag  string
	EmailFlag     string
	VpnFlag       string
	PdmFlag       string
	Remarks       string
}