# Composite Types
1. Arrays and structs are aggregate types; their values are concatenations of other values in memory. Arrays are homogeneous and structs are heterogeneous.
2. Both arrays and structs are fixed size. In contrast, slices and maps are dynamic data structures that grow as values are added.

## Arrays
1. We can use an array literal to initialize an array with a list of values. if an ellipsis `...` appears in place of the length, the array length is deter- mined by the number of initializers.

    ```
    var q [3]int = [3]int{1, 2, 3}
    q := [...]int{1, 2, 3}  // same as above
    q := []int{1, 2, 3} // q is a slice in this case!
    ```
  
2. It is possible to specify a list of index and value pairs in array literal. In this form, indices can appear in any order and some may be omitted; as before, unspecified values take on the zero value for the element type. 

    ```
    type Currency int
    const (
      USD Currency = iota
      EUR
      GBP
      RMB
    )
    symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: """}
    
    r := [...]int{99: -1}      // this defines an array of 100 int
    ```
    
3. If an array’s element type is comparable then the array type is comparable too, so we may directly compare two arrays of that type using the `==` operator, which reports whether all corresponding elements are equal. 

4. When a function is called, a copy of each argument value is assigned to the corresponding parameter variable, so the function receives a copy, not the original. Passing large arrays in this way can be inefficient, and any changes that the function makes to array elements affect only the copy, not the original. This behavior is different from languages that implicitly pass arrays by reference. So use slices for function paramerters.

## Slices
1. Since a slice contains a pointer to an element of an array, passing a slice to a function permits the function to modify the underlying array elements.
2. Left shift a slice by `n` elements can be achieved by reverse the leading `n` elements, reverse the remaining elements, and then reverse the whole slice. Right shift can be achieved by doing the whole reverse first. This is not most time efficient, but is done in place without extra space.
3. Unlike arrays, slices are not comparable, so we cannot use == to test whether two slices contain the same elements. There are two reasons for the language not providing a deep comparison: 1. slices may refer to itself. 2. slices used as map key. Questions: 1. can maps compare? 2. Can slices be used as map keys?
4. The zero value of a slice type is nil. A nil slice has no underlying array. The only legal slice comparison is against nil
5. The built-in function `make` can be used to create a slice of a specified element type, length, and capacity.
6. The build-in funciton `append`can be used to append items to slices. Internally, it checks the capacity, if greater than length, extend the slice by defining a larger slice and add the new item. They share the same underlying array. If not enough capacity, allocate a new array enough to hold the new element, copy all existing elements and the new element. 
7. The strategy for allocation of new array in `append` is to double the existing length, which can achieve amortized linear complexity. (why?)
8. the built-in function `copy` can be used to copy elements from one slice to another of the same type. The total number of elements that are copied is returned, which is the smaller of the length of destination and source slices.
9. The `append` can accept any number of final arguments. This is defined using `y ...int`, and called with `...y`.

