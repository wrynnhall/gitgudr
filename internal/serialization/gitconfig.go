package serialization

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type GitgudrConfig struct {
	IgnoreRepos []string `gitgudr:"ignore_repo"`
	Directories []string `gitgudr:"directory"`
}

type WiprConfig struct {
	IgnoreRepos []string `gitgudr:"ignore_repo" gitgudr_wipr:"ignore_repo"`
	Directories []string `gitgudr:"directory" gitgudr_wipr:"directory"`
}

func Deserialize(data string, result interface{}) error {
	supportedTags := []string{"gitgudr", "gitgudr_wipr"}
	v := reflect.ValueOf(result).Elem()
	t := v.Type()

	dataMap := make(map[string]map[string][]string)
	parts := strings.Split(data, "\n")

	var currentSection string
	sectionRegex := regexp.MustCompile(`\[(\w+)\]`)
	keyValueRegex := regexp.MustCompile(`^\s*(\w+)\s*=\s*(.+)$`)
	for _, part := range parts {
		if matches := sectionRegex.FindStringSubmatch(part); matches != nil {
			currentSection = matches[1]
			if _, exists := dataMap[currentSection]; !exists {
				dataMap[currentSection] = make(map[string][]string)
			}
		} else if currentSection != "" {
			if kvMatches := keyValueRegex.FindStringSubmatch(part); kvMatches != nil {
				key := kvMatches[1]
				value := kvMatches[2]
				dataMap[currentSection][key] = append(dataMap[currentSection][key], value)
			}
		}
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		for _, tag := range supportedTags {
			err := handleTags(&field, &fieldType, tag, dataMap)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func handleTags(field *reflect.Value, fieldType *reflect.StructField, tag string, dataMap map[string]map[string][]string) error {
	if keyValues, ok := dataMap[tag]; ok {
		fieldTag := fieldType.Tag.Get(tag)
		if value, kvOk := keyValues[fieldTag]; kvOk && len(value) >= 1 {

			switch field.Kind() {
			case reflect.String:
				field.SetString(value[0])
			case reflect.Slice:
				if field.Type().Elem().Kind() == reflect.String {
					slice := reflect.MakeSlice(field.Type(), len(value), len(value))
					for i, elem := range value {
						slice.Index(i).SetString(elem)
					}
					field.Set(slice)
				} else {
					return fmt.Errorf("unsupported slice element type: %s", field.Type().Elem().Kind())
				}
			default:
				return fmt.Errorf("unsupported field type: %s", field.Kind())
			}
		}
	}
	return nil
}
