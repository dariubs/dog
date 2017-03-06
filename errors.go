package dog

import(
  "log"
)

// soft error check
func SoftError(err error){
  if err != nil {
    log.Println(err)
  }
}

// fatal error check
func FatalError(err error){
  if err != nil {
    log.Fatalln(err)
  }
}
