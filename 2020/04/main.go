package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	block := []string{}
	var validCount int = 0
	var blockCount int = 0
	var validTotal int = 0

	for scanner.Scan() {
		ln := scanner.Text()
		if strings.TrimSpace(ln) != "" {
			block = append(block, ln)
			continue
		}

		if len(block) != 0 {
			blockCount++

			block = strings.Split(strings.Join(block, " "), " ")
			//fmt.Println(block)

			if validatePassport(block) {
				validCount++
				if validateFields(block) {
					validTotal++
				}
			}

			block = []string{}
		}
		// fmt.Println("---")
	}

	fmt.Println("Part 1:")
	fmt.Printf("Total: %v, Valid: %v\n\n", blockCount, validCount)

	fmt.Println("Part 2:")
	fmt.Printf("Total: %v, Valid: %v", blockCount, validTotal)

}

func validatePassport(fields []string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	sort.Strings(requiredFields)

	//fmt.Printf("block: %v count: %v\n", fields, len(fields))

	var fieldCount int
	//fields := strings.Split(p, " ")

	for _, field := range fields {
		f := strings.TrimSpace(field[:3])
		if f == "cid" {
			continue
		}
		i := sort.SearchStrings(requiredFields, f)
		//fmt.Printf("\"%s\" - %v\n", f, i)

		if i < 7 {
			fieldCount++
		}
	}

	//fmt.Printf("fields: %v\n", fieldCount)

	return fieldCount == 7
}

func validateFields(fields []string) bool {
	var results = make(map[string]bool)

	fmt.Printf("--\n%v\n", fields)
	for _, field := range fields {
		if field[:3] == "cid" {
			continue
		}

		res := validateField(field)

		results[field] = res

		if !res {
			fmt.Printf("%v - %v\n", field, res)
			return false
		}
	}

	return true
}

func validateField(field string) bool {
	parsedField := strings.Split(field, ":")
	// fmt.Println(field)
	switch parsedField[0] {
	case "byr":
		return validateByr(parsedField[1])
	case "ecl":
		return validateEcl(parsedField[1])
	case "eyr":
		return validateEyr(parsedField[1])
	case "hcl":
		return validateHcl(parsedField[1])
	case "hgt":
		return validateHgt(parsedField[1])
	case "iyr":
		return validateIyr(parsedField[1])
	case "pid":
		return validatePid(parsedField[1])
	default:
		return false
	}
}

func validateByr(value string) bool {
	s, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(fmt.Errorf("ValidateField (byr): %v", err))
		return false
	}
	return 1920 <= s && 2002 >= s
}

func validateEcl(value string) bool {
	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if value != color {
			continue
		} else {
			return true
		}
	}
	return false
}

func validateEyr(value string) bool {
	s, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(fmt.Errorf("ValidateField (eyr): %v", err))
		return false
	}
	return 2020 <= s && 2030 >= s
}

func validateHcl(value string) bool {
	matched, err := regexp.MatchString(`^#[[:xdigit:]]{6}$`, value)
	if err != nil {
		fmt.Println(fmt.Errorf("ValidateField (hcl): %v", err))
		return false
	}
	return matched
}

func validateHgt(value string) bool {
	var min, max int
	if strings.HasSuffix(value, "in") {
		min = 59
		max = 76
	} else if strings.HasSuffix(value, "cm") {
		min = 150
		max = 193
	} else {
		return false
	}

	s, err := strconv.Atoi(value[:len(value)-2])
	if err != nil {
		fmt.Println(fmt.Errorf("ValidateField (hgt): %v", err))
		return false
	}
	return min <= s && max >= s
}

func validateIyr(value string) bool {
	s, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(fmt.Errorf("ValidateField (iyr): %v", err))
		return false
	}
	return 2010 <= s && 2020 >= s
}

func validatePid(value string) bool {
	matched, err := regexp.MatchString(`^[[:digit:]]{9}`, value)
	if err != nil {
		fmt.Println(fmt.Errorf("ValidateField (pid): %v", err))
		return false
	}
	return matched
}
