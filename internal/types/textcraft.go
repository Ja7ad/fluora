package types

type TextCraftToolKit string

const (
	Rewrite    TextCraftToolKit = "rewrite"
	Casualize  TextCraftToolKit = "casualize"
	Expand     TextCraftToolKit = "expand"
	Formalize  TextCraftToolKit = "formalize"
	Grammar    TextCraftToolKit = "grammar"
	Shorten    TextCraftToolKit = "shorten"
	WordChoice TextCraftToolKit = "wordchoice"
)

func (t TextCraftToolKit) String() string {
	return string(t)
}
