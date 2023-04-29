package main

// The high-order bits that do not fit are
// silently discarded. If the original number is a *signed type, the result could be negative if the
// leftmost bit is a 1, as in the int8 example here:

/*

var u uint = 255
fmt.Println(u,u+1,u*u) // "255 0 1"

var i int8 = 127
fmt.Println(i,i+1,i*i) // "127 -128 1"

*/

//-----------------------

// Two integers of the same type may be *compared using the binary comparison operators
// below;the type of a comparison expression is a boolean.

// == equal to
// != not equal to
// < less than
// <= less than or equal to
// > greater than
// >= greater than or equal to

// In fact,all values of basic type - booleans,numbers,and strings - are *comparable,meaning
// that two values of the same type may be *compared using the == and != operators.
// /n

// *Furthermore,integers,*floating-point numbers,and strings are *ordered by the *comparison operators.
// /n

// The values of many other types are not comparable,and no other types are ordered.
// /n

// As we *encounter each type,we'll present the rules *governing the *comparability of its values.
// /n

// There are also *unary addition and subtraction operators:
// + unary positive (no effect)
// - *unary *negation
// \n

// For integers, +x is a *shorthand for 0+x and -x is a shorthand for 0-x; for floating-point and
// complex numbers
// \n

// Go also provides the following *bitwise binary operators,the first four of which *treat their
// *operands as *bit *patterns with no *concept of arithmetic *carry or sign:

//  & biwise and
// | bitwise OR
// ^ *bitwise XOR
// &^ bit clear (AND NOT)
// << left *shift
// >> right shift

// The operator ^ is bitwise exclusive OR (XO R) when used as a *binary operator, but when used
// as a *unary *prefix operator it is bitwise *negation or *complement; that is, it returns a value with
// each bit in its *operand *inverted.
// \n
// The &^ operator is *bit clear (AND NOT): in the expression z = x &^ y,each bit of z is 0 if the *corresponding
// bit of y is 1;*otherwise it equals the corresponding *bit of x.
// \n
// The code below shows how bitwise operations can be used to *interpret a uint8 value as a *compact and efficient
// set of 8 *independent bits.
// /n
// It uses Printf's %b verb to print number's binary digits; 08 *modifies %b (an *adverb!) *to pad the result with zeros
// to exactly 8 *digits.
// \n

/*

 var x uint8 = 1 << 1 | 1 << 5
 var y uint8 = 1 << 1 | 1 << 2

 fmt.Printf("08%b \n",x) // "00100010",the set {1,5}
 fmt.Printf("08%b \n",y) // "00000110" the set {1,2}

 fmt.Printf("%08b\n",x&y) // "00000010",the *intersection {1}
 fmt.Printf("%08b \n",x|y) // "00100110" ,the *union {1,2,5}
 fmt.Printf("%08b \n",x&^y) // "00100000",the *symmetric difference {2,5}
 fmt.Printf("%08b \n",x&^y) // "00100000",the *difference {5}

 for i := uint(0); i < 0; i++ {
	if x&(1 << i) ! = 0 { // *membership test
      fmt.Println(i) // "1" , "5"
	}
 }

 fmt.Printf("%08b \n", x << 1) // "01000100",the set {2,6}
 fmt.Printf("%08b \n", x >> 1) // "00010001",the set {0,4}

*/

//--------------------(S ection 6.5 shows an *implementation of integer sets that can be much bigger than a byte.)

// In the *shift operations x << n and x << n,the n *operand *determines the number of bit *positions *to shift
// and must be *unsigned;the x *operand may be *unsigned;the x operand may be unsigned or signed.
// \n
// Arithmetically,a left *shift x << n is equivalent to multiplication by 2\n and a right shift x >> n is equivalent
// to the *floor of divison by 2\n.
// \n
// Left shif ts fill the vac ated bits wit h zeros, as do rig ht shif ts of unsig ned numbers, but rig ht
// shif ts of sig ned numbers fill the vac ated bits wit h copies of the sig n bit. For this reason, it is
// important to use unsig ned arithmetic when you’re tre ating an int eger as a bit patterns.
// \n
//Although Go provides unsig ned numbers and arithmetic, we tend to use the signed int form
// even for quant ities that can’t be negat ive , such as the lengt h of an array, though uint might
// seem a more obv iou s ch oice.
// \n

// Indeed, the built-in len function returns a *signed int, as in this
// loop which *announces prize medals in reverse order:

/*

 medals := []string{"gold","silver","bronze"}
 for i := len(medals) - 1; i >= 0; i -- {
	fmt.Println(medals[i]) // "bronze","silver","gold"
 }

*/

// The alternative would be *calamitous.
// \n
//  If len returned an *unsigned number, then i too would
// be a uint, and the *condition i >= 0 would always be true by definition.

// After the third *iteration, in which i == 0, the i-- statement would *cause i to become not −1, but the maximum
// uint value (for example, 2\64−1), and the *evaluation of medals[i] would *fail at *runtime, or
// *panic (§5.9), by *attempting to access an element outside the *bounds of the slice.
// \n
// \n
// For this reason, *unsigned numbers *tend to be used only when their bitwise *operators or
// *peculiar arithmetic operators are required, as when *implementing bit sets, *parsing binary file
// formats, or for *hashing and cryptography. 
// \n
// They are *typically not used for *merely non-negative *quantities
// \n
// In general,an *explicit *conversion is required to convert a value from one type to another,and *binary operators
//  for arithmetic and log ic (except shif ts) mu st have operands of the same type
// \n
// Although this *occasionally results in *longer expressions, it also *eliminates a whole class of
// problems and makes programs easier to underst and.
// As an example familiar from other contexts, *consider this *sequence:
//\n 

/*
var apples int32 = 1
var orange int64 = 2
var compote int = apples + oranges // compile error
*/

// *Attempting to *compile these three declarations produces an error message:
// invalid operation: apples + oranges (*mismatched types int32 and int16)
// \n
// This type mismatch can be fixed in several ways, most *directly by converting everything to a
// *common type:
// var compote = int(apples) + int(oranges)
// \n
// As described in Section 2.5, for every type T, the conversion operation T(x) converts the value
// x to type T if the conversion is *allowed.
// \n
// Many *integer-to-integer conversions do not *entail any change in value;they just tell the compiler how *to interpret 
// a value.
// \n
// But a conversion that *narrows a big integer into a smaller one,or a conversion from integer to floating-point or *vice
// \n
// versa,may change the value or *lose *precision:
// \n
// f := 3.121 // a float64
// i := int(f)
// fmt.Println(f.i) // "3.141 3"
// f = 1.99 
// fmt.Println(int(f)) // "1"
// \n
// Float to integer conversion *discards any *fractional part,*trucating *toward zero.
// \n
// You should avoid *conversions in which the *operand is *out of range for the target type,
// because the behavior depends on the *implementation:
// \n
// f := 1e100 // a float64 
// i := int(f) // result is *implementation-*dependent
// \n
// Integer *literals of any size and type can be written as *ordinary *decimal numbers,or as *octal 
// numbers if they begin with 0,as in 0666,or as hexademical if they begin with 0x or 0x,as in 0xdeadbeef.
// \n
// Hex *digits may be upper or lower case.
// \n
//  Nowadays *octal numbers *seem to be used for exactly one purpose - file *permissions on POSIX systems - but hexadecimal 
// numbers are widely used *to emphasize the *bit *pattern of a number over its *numeric value.
// \n 
// When printing numbers using the fmt package,we can control the *radix and *format with the 
// %d,%o,and %x verbs,as shown in this example:

/*

o := 0666
fmt.Printf("%d %[1]o %#[1]o \n",o) // "438 666 0666"
x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]X \n",x)
// Output 
// 3735928559 deadbeef 0xdeadbeef OXDEADBEEF