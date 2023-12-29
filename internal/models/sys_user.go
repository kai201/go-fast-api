// Package models  用户信息表，用户信息表
// author : http://www.liyang.love
// date : 2023-11-29 12:04
// desc : 用户信息表，用户信息表
package models

import "time"

// SysUser  用户信息表，用户信息表。
// 说明:用户信息表
// 表名:sys_user
// group: SysUser
// obsolete:
// appliesto:go 1.8+;
// namespace:hongmouer.his.models.SysUser
// assembly: hongmouer.his.models.go
// class:HongMouer.HIS.Models.SysUser
// version:2023-11-29 12:04
type SysUser struct {
	UserId      *int64     `gorm:"column:primaryKey;user_id" json:"userId"` //type:*int64       comment:用户ID                        version:2023-11-29 12:04
	PrvId       *int64     `gorm:"column:prv_id" json:"prvId"`              //type:*int64       comment:上级                          version:2023-11-29 12:04
	UserName    string     `gorm:"column:user_name" json:"userName"`        //type:string       comment:登录账号                      version:2023-11-29 12:04
	NickName    string     `gorm:"column:nick_name" json:"nickName"`        //type:string       comment:用户昵称                      version:2023-11-29 12:04
	Email       string     `gorm:"column:email" json:"email"`               //type:string       comment:用户邮箱                      version:2023-11-29 12:04
	PhoneNumber string     `gorm:"column:phone_number" json:"phoneNumber"`  //type:string       comment:手机号码                      version:2023-11-29 12:04
	Gender      *int       `gorm:"column:gender" json:"gender"`             //type:*int         comment:用户性别（0男;1女；2未知）    version:2023-11-29 12:04
	AvatarUrl   string     `gorm:"column:avatar_url" json:"avatarUrl"`      //type:string       comment:头像路径                      version:2023-11-29 12:04
	Password    string     `gorm:"column:password" json:"password"`         //type:string       comment:密码                          version:2023-11-29 12:04
	Salt        string     `gorm:"column:salt" json:"salt"`                 //type:string       comment:盐加密                        version:2023-11-29 12:04
	Status      *int       `gorm:"column:status" json:"status"`             //type:*int         comment:帐号状态（0正常、1停用）      version:2023-11-29 12:04
	Remark      string     `gorm:"column:remark" json:"remark"`             //type:string       comment:备注                          version:2023-11-29 12:04
	TenantId    *int64     `gorm:"column:tenant_id" json:"tenantId"`        //type:*int64       comment:租户号                        version:2023-11-29 12:04
	Revision    *int64     `gorm:"column:revision" json:"revision"`         //type:*int64       comment:乐观锁                        version:2023-11-29 12:04
	CreatedBy   *int64     `gorm:"column:created_by" json:"createdBy"`      //type:*int64       comment:创建人                        version:2023-11-29 12:04
	CreatedTime *time.Time `gorm:"column:created_time" json:"createdTime"`  //type:*time.Time   comment:创建时间                      version:2023-11-29 12:04
	UpdatedBy   *int64     `gorm:"column:updated_by" json:"updatedBy"`      //type:*int64       comment:更新人                        version:2023-11-29 12:04
	UpdatedTime *time.Time `gorm:"column:updated_time" json:"updatedTime"`  //type:*time.Time   comment:更新时间                      version:2023-11-29 12:04
}

// TableName 表名:sys_user，用户信息表。
// 说明:用户信息表
func (SysUser) TableName() string {
	return "sys_user"
}
