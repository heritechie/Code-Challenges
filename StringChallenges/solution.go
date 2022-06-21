package main
import "fmt"
import "regexp"
import "strings"
import "bytes"

func StringChallenge(str string) string {
  pattern := regexp.MustCompile(`[^a-zA-Z]`)
  splittedStr := pattern.Split(str, -1)
  var buffer bytes.Buffer
  for _, word := range splittedStr {
    buffer.WriteString(strings.Title(strings.ToLower(word)))
  }
  return buffer.String();

}

func main() {

  // do not modify below here, readline is our function
  // that properly reads in the input for you
  fmt.Println(StringChallenge(readline()))

}
