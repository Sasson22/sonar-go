package FunctionTooBig

func foo1() {
  foo2()
  foo2()
}
func foo3() {
}
type AI interface {
    //define all methods that you want to override
    method2()
}

func foo2() { // Noncompliant {{This function has 121 lines of code, which is greater than the 120 authorized. Split it into smaller functions.}}
//   ^^^^
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
  foo1()
}

