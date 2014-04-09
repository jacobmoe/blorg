package template

  func Helpers() {
    "formatTime": func(args ...interface{}) string { 
      t1 := time.Unix(args[0].(int64), 0)
      return t1.Format(time.Stamp)
  }
