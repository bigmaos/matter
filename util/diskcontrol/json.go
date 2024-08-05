package diskcontrol

import (
	"os"
)

// 使用json方式落盘，考虑以后是否换成数据库
func LoadJSON(filepath string) ([]byte, error) {
	jsonBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

func SaveJSON(filepath string, bytes []byte) error {
	err := os.WriteFile(filepath, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
}
