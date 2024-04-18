package model

import (
	"encoding/json"
)

type MessageDecoder struct{}

func (d *MessageDecoder) Decode(data []byte) (*Message, error) {
	var message Message
	err := json.Unmarshal(data, &message)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

func (d *MessageDecoder) WillDecode(data []byte) bool {
	return data != nil
}

func (d *MessageDecoder) Init() {
	// Không cần thực hiện bất kỳ hành động khởi tạo nào trong Go
}

func (d *MessageDecoder) Destroy() {
	// Không cần thực hiện bất kỳ hành động hủy nào trong Go
}
