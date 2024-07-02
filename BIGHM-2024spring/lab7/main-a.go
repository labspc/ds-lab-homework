package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact结构体，表示通讯者信息
type Contact struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	Address     string `json:"address"`
}

// HashTable结构体，表示哈希表
type HashTable struct {
	table []*Contact
	size  int // 长度（采用闭散列法，散列表的长度是确定的）
}

// 初始化哈希表
func Init(size int) *HashTable {
	return &HashTable{
		table: make([]*Contact, size),
		size:  size,
	}
}

// 哈希函数，包含中间过程的回显
func (ht *HashTable) HashFuncPrint(phoneNumber string) int {
	hash := 0
	fmt.Printf("\n 初始哈希值: %d\n", hash)
	for _, char := range phoneNumber {
		charInt := int(char)
		fmt.Printf("处理字符 '%c' (ASCII: %d)\n", char, charInt)

		nextHash := (31*hash + charInt) % ht.size
		fmt.Printf("计算步骤: 31*%d + %d = %d，取模 %d 后得：%d\n", hash, charInt, 31*hash+charInt, ht.size, nextHash)

		hash = nextHash
	}
	fmt.Printf("最终哈希值: %d\n \n", hash)
	return hash
}

// 添加通讯者信息
func (ht *HashTable) Add(contact Contact) {
	index := ht.HashFuncPrint(contact.PhoneNumber)
	for ht.table[index] != nil {
		index = (index + 1) % ht.size
	}
	ht.table[index] = &contact
}

// 查找通讯者信息
func (ht *HashTable) Find(phoneNumber string) *Contact {
	index := ht.HashFuncPrint(phoneNumber)
	for ht.table[index] != nil {
		if ht.table[index].PhoneNumber == phoneNumber {
			return ht.table[index]
		}
		index = (index + 1) % ht.size
	}
	return nil
}

// 格式化输出为JSON的方法
func Format(data any) string {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("JSON编码错误: %s", err)
	}
	return string(jsonData)
}

// 主函数
func main() {
	// 初始化哈希表
	ht := Init(10)

	// 添加通讯者信息
	ht.Add(Contact{"13800138000", "李华", "北京"})
	ht.Add(Contact{"13945678901", "Tom", "上海"})
	ht.Add(Contact{"13722223333", "王敏", "广州"})
	ht.Add(Contact{"13611112222", "Emily", "深圳"})
	ht.Add(Contact{"13500001111", "赵雷", "成都"})
	ht.Add(Contact{"13456789012", "Sophia", "杭州"})

	// 查找通讯者信息
	phoneNumber := "13611112222"
	contact := ht.Find(phoneNumber)
	if contact != nil {
		fmt.Printf("找到通讯者: %s\n", Format(contact))
	} else {
		fmt.Printf("未找到手机号码为 %s 的通讯者\n", phoneNumber)
	}

	// 查找不存在的通讯者信息
	phoneNumber = "0000000000"
	contact = ht.Find(phoneNumber)
	if contact != nil {
		fmt.Printf("找到通讯者: %s\n", Format(contact))
	} else {
		fmt.Printf("未找到手机号码为 %s 的通讯者\n", phoneNumber)
	}
}
