/*
 * @Author: Sunist Chan
 * @Date: 2021-03-30 09:16:23
 * @LastEditTime: 2021-03-30 09:18:53
 * @LastEditors: Sunist Chan
 * @Description:
 * @FilePath: /SForms-BackEnd/src/models/UserInfo.go
 * <!-- i wanna play arknights -->
 */

package models

type UserInfo struct {
	UserId int64
	UserName string
	UserPass string
	UserToken string
}