package tag

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var CopyToClipboard = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Copy the tag to the clipboard",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
		// Include delimiters around the text so that we can assert on the entire content
		config.GetUserConfig().OS.CopyToClipboardCmd = "echo _{{text}}_ > clipboard"
	},
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("one")
		shell.CreateLightweightTag("tag1", "HEAD")
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Tags().
			Focus().
			Lines(
				Contains("tag").IsSelected(),
			).
			Press(keys.Universal.CopyToClipboard)

		t.ExpectToast(Equals("'tag1' copied to clipboard"))

		t.Views().Files().
			Focus().
			Press(keys.Files.RefreshFiles).
			Lines(
				Contains("clipboard").IsSelected(),
			)

		t.Views().Main().Content(Contains("_tag1_"))
	},
})
