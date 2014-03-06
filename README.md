middlecoin-client
=================

Simple middlecoin client written in Go. A demonstration of my API I've written, but with full releases.

####And here are those [releases!](https://github.com/norcalli/middlecoin-client/releases)

######__*Hint hint:*__ Donate to `17Nt7rWiRZKDgcNp421zZ1FHGPWSnnT1bk` if you're in a good mood. Starving college student here.

## Usage:
```
Usage: middlecoin-client <address1> <address2> ...
```

##Example:
```
 ./middlecoin-client 1DgxRTofdbau7kpf3pQeRydcoTPG2L5NUX 17Nt7rWiRZKDgcNp421zZ1FHGPWSnnT1bk
```
Output:

```
1DgxRTofdbau7kpf3pQeRydcoTPG2L5NUX
Last Hour Shares: 27
Immature Balance: 0.00085901
Last Hour Rejected Shares: 0
Paid Out: 0
Unexchanged Balance: 0.00362054
Megahashes Per Second: 0.5033
Bitcoin Balance: 0.00628384
Rejected Megahashes Per Second: 0

Total Profit: 0.010763390000000001
Total Profit(@$650): 6.996203500000001

17Nt7rWiRZKDgcNp421zZ1FHGPWSnnT1bk
Last Hour Shares: 113
Immature Balance: 0.00376599
Last Hour Rejected Shares: 1
Paid Out: 0.07056865
Unexchanged Balance: 0.01496172
Megahashes Per Second: 2.1065
Bitcoin Balance: 0.00109141
Rejected Megahashes Per Second: 0.0186

Total Profit: 0.09038776999999999
Total Profit(@$650): 58.752050499999996

Last Hour Shares: 140
Immature Balance: 0.004625
Last Hour Rejected Shares: 1
Paid Out: 0.07056865
Unexchanged Balance: 0.01858226
Megahashes Per Second: 2.6098
Bitcoin Balance: 0.00737525
Rejected Megahashes Per Second: 0.0186

Total Profit: 0.10115116
Total Profit(@$650): 65.748254
```