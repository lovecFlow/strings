package main

import (
	"strings"
	"unicode"
)

func latinLetters(s string) string { 
	sb := &strings.Builder{} //Создаём билдер

	for _, r := range s { //Ставим цикл с ренджом s 
		if unicode.Is(unicode.Latin, r)  { //Проверка на латиницу 
			sb.WriteRune(r) //Записываем прошедшие проверку руны в sb 
		}
    }

    return sb.String() //Возвращаем строку только с латинскими символами 
}
