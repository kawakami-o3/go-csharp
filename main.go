package main

import "fmt"

func main() {
	cs := `
using System;

public class Hello{
  public static void Main(){
    Console.WriteLine("hello world!");
  }
}
`
	fmt.Println(cs)
}
