package constant

import "github.com/manifoldco/promptui"

var (
	IconOkay    = promptui.Styler(promptui.FGBlue)("✔")
	IconNotOkay = promptui.Styler(promptui.FGRed)("✘")
	IconWarn    = promptui.Styler(promptui.FGBold, promptui.FGYellow)("-")
	IconCaution = promptui.Styler(promptui.FGBold, promptui.FGYellow)("*")
)

const (
	LightBlue          = "\033[94m"
	Green              = "\033[92m"
	Grey               = "\033[2m"
	Magenta            = "\033[35m"
	MagentaUnderline   = "\033[35;4m"
	MagentaNoUnderline = "\033[35;24m"
	Reset              = "\033[0m"
)

const ResultBanner = ` ______    _______  _______  __   __  ___      _______
|    _ |  |       ||       ||  | |  ||   |    |       |
|   | ||  |    ___||  _____||  | |  ||   |    |_     _|
|   |_||_ |   |___ | |_____ |  |_|  ||   |      |   |
|    __  ||    ___||_____  ||       ||   |___   |   |
|   |  | ||   |___  _____| ||       ||       |  |   |
|___|  |_||_______||_______||_______||_______|  |___|  `
