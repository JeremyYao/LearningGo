# Go Notes

## Closures

``` Go
func concatter() func(string) string {
 doc := ""
 return func(word string) string {
  doc += word + " "
  return doc
 }
}

func main() {
 harryPotterAggregator := concatter() // Initializes doc and then it stores a func(string) string
 // When we invoke harryPotterAggregator, we call the func(string) string which appends to doc
 harryPotterAggregator("Mr.") // append Mr. 
 harryPotterAggregator("and") // append and...
 harryPotterAggregator("Mrs.")
 harryPotterAggregator("Dursley")
 harryPotterAggregator("of")
 harryPotterAggregator("number")
 harryPotterAggregator("four,")
 harryPotterAggregator("Privet")

 fmt.Println(harryPotterAggregator("Drive"))
 // Mr. and Mrs. Dursley of number four, Privet Drive
}
```
