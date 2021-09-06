package main

/* Remove a item of slice but not to keep item order
 */
func removeSliceItemNotKeepOrder(a []string, i int) []string {
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}
