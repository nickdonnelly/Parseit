package printhelper

import (
  "fmt"
  // "github.com/dotabuff/manta"
  "parseit/helpers/stringhelper"
)

// Printable structures:
var DeathData = make(map[string] int)
var HeroNames []string

type PrintableData struct{
  DataSet string // the name of the dataset
  Data map[string]string // a map containing printable data
}

func PrintSingle(toPrint PrintableData) {
  toPrint.print()
}


// This function assumes the keys are already formatted for printing. Ensure that when you make the keys in the data map that they are formatted how you want them to be printed, i.e. don't use hero internal names instead of the printable hero names.
func (pStruct *PrintableData) print(){
  alpha := stringhelper.GetAlphabetizedKeyListFromMap(pStruct.Data)
  fmt.Printf("+------------------------------------------------+\n")
  fmt.Printf("| %-46v |\n", pStruct.DataSet)
  for _, key := range alpha {
    fmt.Printf("| %-39v: %5v |\n", key, pStruct.Data[key])
  }
  fmt.Printf("+------------------------------------------------+\n")
}

func (pStruct *PrintableData) getPrintString() string{
  alpha := stringhelper.GetAlphabetizedKeyListFromMap(pStruct.Data)
  result := fmt.Sprintf("+------------------------------------------------+\n")
  result += fmt.Sprintf("| %-46v |\n", pStruct.DataSet)
  for _, key := range alpha {
    result += fmt.Sprintf("| %-39v: %5v |\n", key, pStruct.Data[key])
  }
  result += fmt.Sprintf("+------------------------------------------------+\n")
  return result
}

// func (slice []string) contains(s string) bool{
//   for _, v := range slice{
//     if v == s{
//       return true
//     }
//   }
//   return false
// }
