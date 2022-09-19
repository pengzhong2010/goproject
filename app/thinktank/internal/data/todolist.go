package data

import "time"

type Todolist struct {
	ID        int64     `json:"id"`                    // 自增主键ID
	Title     string    `json:"title"`                 // title
	Detail    string    `json:"detail"`                // 详细信息
	Ctime     time.Time `json:"ctime" gorm:"<-:false"` // 创建时间
	Mtime     time.Time `json:"mtime" gorm:"<-:false"` // 修改时间
	IsDeleted int64     `json:"is_deleted"`            // 0-未删除,非0-Unix时间戳
}

func (*Todolist) TableName() string {
	return "todolist"
}
