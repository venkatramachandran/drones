# Drones

The solution is done in Golang. To get the code to work, follow the steps below:

# How To Install Go
Please refer to instructions [here](https://golang.org/doc/install#install).

# Setting up GOPATH
Please refer to instructions [here](https://golang.org/doc/code.html#GOPATH)

# Getting the source code
Once Go is installed and a $GOPATH is configured, run 
```
go get -u github.com/venkatramachandran/drones
```
This will download the code and put it in `$GOPATH/src/github.com/venkatramachandran/drones`

# Building and running the code
To build, run
```
go build
```
from `$GOPATH/src/github.com/venkatramachandran/drones`

This will compile the code and create an executable in the same directory.

The executable can then be run with:
```
./drones
```
# Running Tests
To run the tests, run
```
go test
```
from `$GOPATH/src/github.com/venkatramachandran/drones`

## How did you implement your solution?
The solution does a run through all packages that need to be delivered, sorted by earliest delivery. It then picks each drone, sorted by available time, to see if it can make the delivery.
If the first available drone cannot make the delivery, then none of the other can make the delivery. So, the package gets added to the unassigned list.
This continues until all packages are assigned or all drones are assigned and none remain to be assigned.

## Why did you implement it this way?
I looked at the simplest way to solve the problem on a minimal scale. Comparing deadine of delivery to (availabile time of drone + time to deliver) was the simplest way to go.

## Let's assume we need to handle dispatching thousands of jobs per second to thousands of drivers. Would the solution you've implemented still work? Why or why not? What would you modify? Feel free to describe a completely different solution than the one you've developed.

No, this will not work at scale. In the current solution, every time a new drone is available to deliver or a package is added for delivery, assignment service has to be run from start.
A better solution would be to have an event driven model - i.e. whenever a new package is added to the pool, the next available drone is picked and assigned for delivery.