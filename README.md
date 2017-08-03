# Drones

## How did you implement your solution?
The solution does a run through all packages that need to be delivered, sorted by earliest delivery. It then picks each drone, sorted by available time, to see if it can make the delivery.
If the first available drone cannot make the delivery, then none of the other can make the delivery. So, the package gets added to the unassigned list.
This continues until all packages are assigned or all drones are assigned and none remain to be assigned.

## Why did you implement it this way?
I looked at the simplest way to solve the problem on a minimal scale. Comparing deadine of delivery to (availabile time of drone + time to deliver) was the simplest way to go.

## Let's assume we need to handle dispatching thousands of jobs per second to thousands of drivers. Would the solution you've implemented still work? Why or why not? What would you modify? Feel free to describe a completely different solution than the one you've developed.

No, this will not work at scale. In the current solution, every time a new drone is available to deliver or a package is added for delivery, assignment service has to be run from start.
A better solution would be to have an event driven model - i.e. whenever a new drone is added to the pool, the next package that can be delivered is assigned (or) whenever a new package is added to the pool, the next available drone is picked and assigned for delivery.