package template

import (
	"fmt"
	"strings"
)

type Evaluator func(string) (interface{}, error)

func templateEval(eval Evaluator) func(values ...interface{}) string {
	return func(values ...interface{}) string {
		str := ""
		for _, v := range values {
			str = fmt.Sprintf("%s %v", str, v)
		}
		result, err := eval(str)
		if err != nil {
			return err.Error()
		}
		return fmt.Sprintf("%v", result)
	}
}

func BuildDataMap(values ...interface{}) (map[string]interface{}, error) {
	ret := map[string]interface{}{}
	if len(values)%2 == 1 {
		return nil, fmt.Errorf("Must have a pair number of param : a key and a value")
	}
	for i := 0; i < len(values); i += 2 {
		ret[fmt.Sprintf("%v", values[i+0])] = values[i+1]
	}
	return ret, nil
}

func BuildDataArray(values ...interface{}) ([]interface{}, error) {
	ret := []interface{}{}
	for i := 0; i < len(values); i++ {
		ret = append(ret, values[i])
	}
	return ret, nil
}

// func Translate(tr i18n.Translation) func(string, ...interface{}) (string, error) {
// 	return func(key string, p ...interface{}) (string, error) {
// 		if len(p)%2 == 1 {
// 			return "", fmt.Errorf("Must have a pair number of param : a key and a value")
// 		}
// 		ps := []i18n.Param{}
// 		for i := 0; i < len(p); i += 2 {
// 			ps = append(ps, i18n.Param{
// 				Old: fmt.Sprintf("%v", p[i+0]),
// 				New: fmt.Sprintf("%v", p[i+1]),
// 			})
// 		}
// 		return tr.Trans(key, ps), nil
// 	}
// }

func Upper(v interface{}) string {
	return strings.ToUpper(fmt.Sprintf("%v", v))
}
