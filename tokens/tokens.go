package tokens

type Token string

const(
	Divide Token = "/"
	Multiply Token = "*"
	Plus Token = "+"
	Minus Token = "-"

	EqualTo Token = "=="
	NotEqualTo Token = "!="
	LogicalAnd Token = "&&"
	LogicalOr Token = "||"
)
