/*
 * @Author: Sunist Chan
 * @Date: 2021-03-30 09:38:01
 * @LastEditTime: 2021-03-30 09:45:12
 * @LastEditors: Sunist Chan
 * @Description:
 * @FilePath: /SForms-BackEnd/src/models/Runtime.go
 * <!-- i wanna play arknights -->
 */

package models

type DataBaseConnection struct {
	DataBaseAddress string 
	DataBaseUsername string
	DataBasePassword string
	DataBasePort int
}

type FrontEndConnection struct {
	FrontEndAddress string
}