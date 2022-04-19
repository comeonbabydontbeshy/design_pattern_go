package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct{}

func (j jsonRuleConfigParser) Parse(data []byte) {
	panic("jsonRuleConfigParser implement me")
}

type yamlRuleConfigParser struct{}

func (y yamlRuleConfigParser) Parse(data []byte) {
	panic("yamlRuleConfigParser implement me")
}

func NewIRuleConfigParser(str string) IRuleConfigParser {
	switch str {
	case "json":
		return &jsonRuleConfigParser{}
	case "yaml":
		return &yamlRuleConfigParser{}
	}
	return nil
}
