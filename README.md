### bitarray

Is an implementation of variable size bit array. Array is zero index based.

Sample usage:

```
ary := bitarray.New(10)
ary.IsSet(2) //false
ary.Set(2)
ary.IsSet(2) //true
ary.String() // 0010000000
ary.Unset(2)
ary.IsUnset(2) //true
ary.String() // 0000000000
```

For more see example in folder eratosthenes
