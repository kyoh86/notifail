package internal

import (
	"fmt"
	"os/exec"
)

func Notify(level Level, timeout int, title, message, detailFilename string) error {
	command := exec.Command("osascript", "-l", "JavaScript", "-")
	pipe, err := command.StdinPipe()
	if err != nil {
		return fmt.Errorf("failed to open pipe to osascript by %w", err)
	}

	buttons := `["OK"]`
	callback := ``
	if detailFilename != "" {
		buttons = `["Details", "OK"]`
		callback = fmt.Sprintf(`
if (reply.buttonReturned === "Details") {
	var textEdit = Application("TextEdit");
	textEdit.open(%q);
}`, detailFilename)
	}
	fmt.Fprintf(pipe, `
var app = Application("System Events");
app.includeStandardAdditions = true;
var reply = app.displayDialog(%q, {
	"withTitle": %q,
	"withIcon":%q,

	"givingUpAfter": %d,

	"defaultButton": "OK",
	"buttons": %s
});%s`, message, title, level, timeout, buttons, callback)
	return command.Run()
}
