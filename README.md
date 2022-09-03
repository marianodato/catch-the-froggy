# catch-the-froggy
A simple script in Go to test the solution of the catch the froggy problem.

## Problem
There are n boxes in a row and one of them contains a frog. During the day only one box can be opened, if the frog is there the game is 
completed, if not the box is closed again and at night the frog is forced to move to an adjacent box. The next day the same process is 
repeated and so on. The goal is to find an algorithm that always allows the frog to be found with any number of steps.

## Proposed solution
The solution proposed is the sequence {n-1, n-1, n-2, n-3, ..., 2, 2, 3, 4, ..., n} with n as the number of boxes. For the special case 
of n = 1 the solution sequence is {1}.