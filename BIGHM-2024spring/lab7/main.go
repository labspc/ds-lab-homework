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
		// make用于创建一个切片，类型为([]*Contact)，长度和容量都是size。
		table: make([]*Contact, size),
		size:  size,
	}
}

// 哈希函数
func (ht *HashTable) HashFunc(phoneNumber string) int {
	hash := 0
	for _, char := range phoneNumber {
		hash = (31*hash + int(char)) % ht.size
	}
	return hash
}

// 添加通讯者信息
// Add 方法用于向哈希表中添加一个新的联系人。
// 它首先使用 HashFunc 函数计算联系人电话号码的哈希值，然后尝试在哈希表中找到一个空的位置来存储联系人信息。
// 如果哈希表中当前位置已存在数据，则通过线性探测法寻找下一个空位置。
// 参数:
//
//	contact - 需要添加到哈希表的联系人对象。
func (ht *HashTable) Add(contact Contact) {
	// 使用哈希函数计算联系人电话号码的 哈希值。
	index := ht.HashFunc(contact.PhoneNumber)

	// 循环查找哈希表中第一个空的位置。
	// 如果当前位置已占用，则通过线性探测法寻找下一个空位置。
	for ht.table[index] != nil {
		index = (index + 1) % ht.size
	}

	// 在找到的空位置上存储联系人信息。
	ht.table[index] = &contact
}

// 查找通讯者信息
// Find 根据电话号码查找联系人。
// 它使用哈希表的索引查找联系人，如果找到匹配的电话号码，则返回该联系人。
// 如果没有找到匹配的电话号码，则返回nil。
// 参数:
// phoneNumber - 要查找的电话号码。
// 返回值:
// *Contact - 匹配的联系人指针，如果没有找到，则为nil。
func (ht *HashTable) Find(phoneNumber string) *Contact {
	// 使用HashFunc计算电话号码的哈希索引。
	index := ht.HashFunc(phoneNumber)
	// 遍历哈希表，直到找到空的槽位或匹配的电话号码。
	for ht.table[index] != nil {
		// 如果当前槽位的电话号码与目标电话号码匹配，则找到了联系人。
		if ht.table[index].PhoneNumber == phoneNumber {
			return ht.table[index]
		}
		// 如果当前槽位不匹配，则通过取模运算寻找下一个槽位。
		index = (index + 1) % ht.size
	}
	// 如果遍历完成后没有找到匹配的电话号码，则返回nil。
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
