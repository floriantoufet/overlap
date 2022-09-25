# IP Overlap

IP Overlap CLI  prints to STDOUT the relation between
two CIDRs.
<br/>The relations can be:
<br/>• subset: if the network of the second address is included in the first one
<br/>• superset: if the network of the second address includes the first one
<br/>• different: if the two networks are not overlapping
<br/>• same: if both address are in the same network
<br/>The program is only intended to work with IPv4 addresses.

### Build binary

go mod download
<br/>go build -o overlap main.go

### Run

 ```shell script
 $ chmod a+x overlap
 $ ./overlap  10.0.0.0/20 10.0.2.0/24
 ```

### Tests

#### Requirements

- install [golangci](https://github.com/golangci/golangci-lint)

#### Run tests

 ```shell script
 $ ./run_test.sh
 ```
