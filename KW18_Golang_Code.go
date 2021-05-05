//--------------------------------------------------
//Folie 7 - Syntax und Semantik

package main
 
import (
	"fmt"
)
 
func main() {
	var x int = 200
 
	if x == 200 {
		fmt.Println("USA")
	} else {
		fmt.Println("Canada")
	}
}


//--------------------------------------------------
//Folie 11 - Standard Datentypen

package main
  
import "fmt"
  
func main() {

    var x int8 = 182

    fmt.Println(x)
 
}


//--------------------------------------------------
//Folie 12 - Standard Datentypen

package main
 
import (
    "fmt"
)
 
func main() {
    var f float32 
    var f2 float64 
    f = 12.34567890123456
    f2 = 12.34567890123456
    fmt.Println(f, f2)  // prints "12.345679 12.34567890123456"
}



//--------------------------------------------------
//Folie 13 - Alphanumerische Datentypen (string, byte, rune)
package main

import (
    "fmt"
    "reflect"
    "unsafe"
)

func main() {
    //If you don't specify type here
    var b byte = 'a'
    
    fmt.Println("Priting Byte:")
    //Print Size, Type and Character
    fmt.Printf("Size: %d\nType: %s\nCharacter: %c\n", unsafe.Sizeof(b), reflect.TypeOf(b), b)
    
    r := '£'
    
    fmt.Println("\nPriting Rune:")
    //Print Size, Type, CodePoint and Character
    fmt.Printf("Size: %d\nType: %s\nUnicode CodePoint: %U\nCharacter: %c\n", unsafe.Sizeof(r), reflect.TypeOf(r), r, r)

    s := "µ" //Micro sign
    fmt.Println("\nPriting String:")
    fmt.Printf("Size: %d\nType: %s\nCharacter: %s\n", unsafe.Sizeof(s), reflect.TypeOf(s), s)
}



//--------------------------------------------------
//Folie 14 - Logische Datentypen (boolean)

 package main
 // program to perform comparison operation for the boolean data type.
 import "fmt"
 func main() {
             var num1 = 10
             var num2 = 12                       // variable declaration complete.
             // use of comparison operator on the declared variable.
             var res1 = num1 == num2
             var res2 = num1 != num2
             var res3 = num1 > num2
             var res4 = num1 < num2
             var res5 = num1 <= num2
             var res6 = num1 >= num2
             // all six results will store boolean value here.
             fmt.Println(res1, res2, res3, res4, res5, res6)
 } 


//--------------------------------------------------
//Folie 16 - Variablen


package main

import (
  "fmt"
)

func main() {
 var ErreichtePunkte float32
 const MaxPunkte float32 = 100

	fmt.Println("Bitte die erreichten Punkte eintragen: (mit ENTER bestätigen)")
	fmt.Scanln(&ErreichtePunkte)
	
	fmt.Println("Die maximal möglichen Punkte sind:", MaxPunkte)
	
	var Note float32 = (ErreichtePunkte/MaxPunkte)*5+1
	
	Nachname := "Kehrli"
	
	fmt.Println("Student", Nachname, "hat die Note:", Note)
}



//--------------------------------------------------
//Folie 19 - Konstanten


package main

import (  
    "fmt"
)

func main() {  
    const (
        Name = "Kehrli"
        Note = "5.5"
        Modul = "A06"
    )
    fmt.Println(Name)
    fmt.Println(Note)
    fmt.Println(Modul)

}



