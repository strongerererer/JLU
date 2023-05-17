module server

go 1.18

require (
	base v1.0.0
	github.com/sashabaranov/go-openai v1.8.0
)

require gopkg.in/yaml.v3 v3.0.1

replace base v1.0.0 => ../serverbase/base
