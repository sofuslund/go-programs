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
    var fromUnit = getRune("Enter the unit you want to convert from: ")
    if fromUnit != 'C' && fromUnit != 'F' && fromUnit != 'K'{
        log.Fatal("The temperature unit has to be either 'C' 'F' or 'K'")
    }
    // Get the number of degrees
    var str string
    if fromUnit != 'K' {
        str = fmt.Sprintf("Enter how many degrees of %s: ", units[fromUnit])
    } else {
        str = "Enter how many kelvin: "
    }
    var degrees = getFloat64(str)
    // Get the measurement unit the user want to convert to
    var toUnit = getRune("Enter the unit you want to convert to: ")
    if toUnit != 'C' && toUnit != 'F' && toUnit != 'K'{
        log.Fatal("The temperature unit has to be either 'C' 'F' or 'K'")
    }

    return fromUnit, degrees, toUnit
}

func main() {
    // Set the logging options
    log.SetPrefix("temperature_converter: ")
    log.SetFlags(0)
    // Set associations between the letters and words of the temperature units
    units = make(map[rune]string)
    units['C'] = "celcius"
    units['F'] = "fahrenheit"
    units['K'] = "kelvin"
    // Show the abbreviations for the temperature units
    fmt.Println("Type C for Celcius")
    fmt.Println("Type F for Fahrenheit")
    fmt.Println("Type K for Kelvin")
    // Get input
    var fromUnit, degrees, toUnit = getInput()
    // Calculate the result
    fmt.Println("Calculating...")
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
        resString += "degrees "
    }
    resString += units[fromUnit]
    resString += " is equal to "
    resString += strconv.FormatFloat(result, 'f', -1, 64)
    resString += " "
    if toUnit != 'K' {
        resString += "degrees "
    }
    resString += units[toUnit]
    fmt.Print(resString)
}
