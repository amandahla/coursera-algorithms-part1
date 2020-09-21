Union-find with specific canonical element. Add a method \mathtt{find()}find() to the union-find data type so that \mathtt{find(i)}find(i) returns the largest element in the connected component containing ii. The operations, \mathtt{union()}union(), \mathtt{connected()}connected(), and \mathtt{find()}find() should all take logarithmic time or better.

For example, if one of the connected components is \{1, 2, 6, 9\}{1,2,6,9}, then the \mathtt{find()}find() method should return 99 for each of the four elements in the connected components.

