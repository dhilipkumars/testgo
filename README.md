# testgo
This a simple app writen in go to understand cloud foundry.  it computes factorial of a given number and returns the result.  this is very helpful as the app can be used to stress the peformance 

```
$cf apps
Getting apps in org dhilip / space gospace as admin...
OK

name     requested state   instances   memory   disk   urls
testgo   started           16/16       64M      1G     testgo.10.244.0.34.xip.io
```

then

```
$time curl http://testgo.10.244.0.34.xip.io/100
93326215443944152681699238856266700490715968264381621468592963895217599993229915608941463976156518286253697920827223758251185210916864000000000000000000000000

real    0m0.536s
user    0m0.006s
sys     0m0.006s
```
you can compute factorial of multiple numbers in one request.

```
$ time curl http://testgo.10.244.0.34.xip.io/5/10/15/20
120
3628800
1307674368000
2432902008176640000

real    0m0.036s
user    0m0.008s
sys     0m0.005s
```

or you could use the simulator to send requests in parllel 

here we are trying to send 4 parllel request to compute factorial of '100000'.  

```
$time get_request -P="http://testgo.10.244.0.34.xip.io/100000" -T=4
flags Parllel=4 Page=http://testgo.10.244.0.34.xip.io/100000
reading http://testgo.10.244.0.34.xip.io/100000
reading http://testgo.10.244.0.34.xip.io/100000
reading http://testgo.10.244.0.34.xip.io/100000
reading http://testgo.10.244.0.34.xip.io/100000

real    0m1.453s
user    0m0.009s
sys     0m0.009s
```
