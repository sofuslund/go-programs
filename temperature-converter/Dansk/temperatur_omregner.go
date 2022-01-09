package main

import (
    "os"
    "fmt"
    "log"
    "bufio"
    "strings"
    "strconv"
    "unicode"
)

var units map[rune]string

func getRune(str string) rune {
    fmt.Print(str)
    reader := bufio.NewReader(os.Stdin)
    tmp, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    return unicode.ToUpper(rune(tmp[0]))
}

func getFloat64(str string) float64 {
    fmt.Print(str)
    reader := bufio.NewReader(os.Stdin)
    tmp, err := reader.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    var r float64
    r, err = strconv.ParseFloat(strings.ReplaceAll(tmp, "\r\n", ""), 64)
    if err != nil {
        log.Fatal(err)
    }
    return r
}

func getInput() (rune, float64, rune) {
    // Get the measurement unit the user want to convert from
    var fromUnit = getRune("Indtask hvilken temperatur-måleenhed du vil omregne fra: ")
    if fromUnit != 'C' && fromUnit != 'F' && fromUnit != 'K'{
        log.Fatal("Temperatur-måleenheden skal enten være 'C' 'F' eller 'K'")
    }
    // Get the number of degrees
    var str string
    if fromUnit != 'K' {
        str = fmt.Sprintf("Indtast antal grader af %s: ", units[fromUnit])
    } else {
        str = "Indtast antal af kelvin: "
    }
    var degrees = getFloat64(str)
    // Get the measurement unit the user want to convert to
    var toUnit = getRune("Indtast temperatur-måleenheden du vil omregne til: ")
    if toUnit != 'C' && toUnit != 'F' && toUnit != 'K'{
        log.Fatal("Temperatur-måleenheden skal enten være 'C' 'F' eller 'K'")
    }

    return fromUnit, degrees, toUnit
}

func main() {
    // Set the logging options
    log.SetPrefix("temperatur_omregner: ")
    log.SetFlags(0)
    // Set associations between the letters and words of the temperature units
    units = make(map[rune]string)
    units['C'] = "celcius"
    units['F'] = "fahrenheit"
    units['K'] = "kelvin"
    // Show the abbreviations for the temperature units
    fmt.Println("Skriv C for celcius")
    fmt.Println("Skriv F for fahrenheit")
    fmt.Println("Skriv K for kelvin")
    // Get input
    var fromUnit, degrees, toUnit = getInput()
    // Calculate the result
    fmt.Println("Beregner...")
    var result float64
    if fromUnit == toUnit {
        result = degrees
    } else if fromUnit == 'C' && toUnit == 'F' {
        result = 9./5.*degrees+32.
    } else if fromUnit == 'K' && toUnit == 'F' {
        result = 9./5.*(degrees - 273.) + 32.
    } else if fromUnit == 'F' && toUnit == 'C' {
        result = 5./9.*(degrees - 32.)
    } else if fromUnit == 'C' && toUnit == 'K' {
        result = degrees + 273.
    } else if fromUnit == 'K' && toUnit == 'C' {
        result = degrees - 273.
    } else if fromUnit == 'F' && toUnit == 'K' {
        result = 5./9.*(degrees - 32.) + 273.
    }
    // Show the result
    resString := strconv.FormatFloat(degrees, 'f', -1, 64)
    resString += " "
    if fromUnit != 'K' {
        resString += "grader "
    }
    resString += units[fromUnit]
    resString += " er lig med "
    resString += strconv.FormatFloat(result, 'f', -1, 64)
    resString += " "
    if toUnit != 'K' {
        resString += "grader "
    }
    resString += units[toUnit]
    fmt.Print(resString)
}
