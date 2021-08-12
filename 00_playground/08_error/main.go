package main

import (
	"errors"
	"fmt"
	"os"
)

type inputError struct {
	message      string
	missingField string
}

func (i *inputError) Error() string {
	return i.message
}

func (i *inputError) getMissingField() string {
	return i.missingField
}

type errorOne struct{}

func (e errorOne) Error() string {
	return "Error One happened"
}

func main() {
	// 에러 만드는 방법
	// sampleErr := errors.New("error occured")
	// sampleErr := fmt.Errorf("Err is: %s", "database connection issue")
	// fmt.Println(sampleErr)

	// 사용자 정의 에러
	// err := validate("", "")
	// if err != nil {
	//     if err, ok := err.(*inputError); ok {
	//         fmt.Println(err)
	//         fmt.Printf("Missing Field is %s\n", err.getMissingField())
	//     }
	// }

	// 에러 래핑, 언래핑 = 스택트레이스
	// e1 := errorOne{}
	// e2 := fmt.Errorf("E2: %w", e1)
	// e3 := fmt.Errorf("E3: %w", e2)
	// fmt.Println(errors.Unwrap(e3))
	// fmt.Println(errors.Unwrap(e2))
	// fmt.Println(errors.Unwrap(e1))

	// 에러타입 확인
	// err1 := errorOne{}
	// err2 := do()

	// // == 을 사용하면, 다른 객체이기 대문에 다른 에러로 확인된다
	// if err1 == err2 {
	// 	fmt.Println("Equality Operator: Both errors are equal")
	// } else {
	// 	fmt.Println("Equality Operator: Both errors are not equal")
	// }

	// // Is 메소드를 활용하면 같은 에러로 확인
	// if errors.Is(err2, err1) {
	// 	fmt.Println("Is function: Both errors are equal")
	// }

	// 에러 타입 어서션
	err := openFile("non-existing.txt")

	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("Using Assert: Error e is of type path error. Path: %v\n", e.Path)
	} else {
		fmt.Println("Using Assert: Error not of type path error")
	}

	var pathError *os.PathError
	if errors.As(err, &pathError) {
		fmt.Printf("Using As function: Error e is of type path error. Path: %v\n", pathError.Path)
	}
}

func openFile(fileName string) error {
	_, err := os.Open("non-existing.txt")
	if err != nil {
		return err
	}
	return nil
}

func do() error {
	return errorOne{}
}

// func validate(name, gender string) error {
// 	if name == "" {
// 		return &inputError{message: "Name is mandatory", missingField: "name"}
// 	}
// 	if gender == "" {
// 		return &inputError{message: "Gender is mandatory", missingField: "gender"}
// 	}
// 	return nil
// }
