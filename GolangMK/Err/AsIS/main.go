// Функция errors.Is сравнивает ошибку со значением.
// Функция errors.As проверяет, совпадает ли ошибка
// по типу с указанной. Если это так,
// то устанавливает значение указанной ошибки
// в проверяемое и возвращает true, иначе — false
package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}
