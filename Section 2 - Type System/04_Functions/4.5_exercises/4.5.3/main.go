// write a function which accepts a slice of floats and
// return three values -> slice, boolean, error
//
// 1. Multiply all the slice elements with 2 and return a new slice,
// the underlying slice should not be modified, otherwise return a blank slice.
// 2. If all the elements are below 10, then boolean value should be
// true otherwise false.
// 3. If there are any negative elements in slice, then return error saying
// "Validation Error: only positive numbers are allowed", otherwise return nil