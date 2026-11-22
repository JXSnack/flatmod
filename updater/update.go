package updater

const URL = "https://raw.githubusercontent.com/JXSnack/flatmod/refs/heads/main/app.json"
const Current = 1
const Display = "0.1.0"

type Update struct {
	Latest  int    `json:"latest"`
	Message string `json:"message"`
}
