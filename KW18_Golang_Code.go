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

import (
	"fmt"
)

func main() {
	// boolean operators
	a := 3
	b := 4
	fmt.Println(a == b)  // false
	fmt.Println(a != b)  // true
	fmt.Println(a < b)   // true
	fmt.Println(a > b)   // false
	fmt.Println(a >= b)  // false
	fmt.Println(a <= b)  // true
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



