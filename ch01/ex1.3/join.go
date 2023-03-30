// package join Ex1.3 Поэкспериментируйте с измерением разницы времени выполнения
// потенциально неэффективных версий и версий с применением strings.Join.
// (В Разделе 1.6 демонстрируется часть пакета time,
// а в разделе 11.4 - как написать тест производительности
// для ее систематической оценки )

package join

import "strings"

// StrJoin принимает на вход слайс строк и строку разделитель
// возвращает строку склеенную из элементов слайса разделенных сепаратором
func StrJoinRange(str []string) string {
	result := ""
	for _, s := range str {
		result += s
	}
	return result
}

func StrJoinIterated(str []string) string {
	result := ""
	for i := 0; i < len(str); i++ {
		result += str[i]
	}
	return result
}

// StringsJoin использует библиотечную функцию для склейки строк
func StringsJoin(str []string, sep string) string {
	return strings.Join(str, sep)
}

func StringsJoinBuilder(str []string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()
}
