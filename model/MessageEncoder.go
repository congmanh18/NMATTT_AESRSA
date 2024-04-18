package model

import (
	"encoding/json"
)

type MessageEncoder struct{}

func (e *MessageEncoder) Encode(message *Message) ([]byte, error) {
	jsonData, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (e *MessageEncoder) Init() {
	// Không cần thực hiện bất kỳ hành động khởi tạo nào trong Go
}

func (e *MessageEncoder) Destroy() {
	// Không cần thực hiện bất kỳ hành động hủy nào trong Go
}
