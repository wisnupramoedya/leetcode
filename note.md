## Map and Byte
It is faster to use:
````
// Runtime: 0 ms
romans := make(map[byte]int)
romans['I'] = 1
romans['V'] = 5
romans['X'] = 10
romans['L'] = 50
romans['C'] = 100
romans['D'] = 500
romans['M'] = 1000
````
instead of:
````
// Runtime: 9 ms
romans := map[uint8]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
````
- The use of variable assignment affects 9 ms -> 7 ms
- The use of byte affects 7 ms -> 0 ms