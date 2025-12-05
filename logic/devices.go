package logic

import "strings"

func GetDevices(output []byte) ([]string, error) {
	lines := strings.Split(string(output), "\n")
	var devices []string
	for _, line := range lines {
		// 跳过空行和标题行
		if line == "" || strings.HasPrefix(line, "List of devices") {
			continue
		}
		// 提取设备序列号（每行的第一部分）
		parts := strings.Fields(line)
		if len(parts) >= 2 && parts[1] == "device" {
			devices = append(devices, parts[0])
		}
	}
	if len(devices) == 0 {
		err := ErrorNoDevices
		return nil, err
	}
	return devices, nil
}
