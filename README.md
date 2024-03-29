##
# Go API server with goroutines
#
#
# Level -   Advanced
# Time  -   1 - 4 Hours
#
#
# Objective
##

This task is designed to test your knowledge of Go and your ability to problem solve and troubleshoot some areas that may be new to you.

##
# Task
##

Create a GO API server that has a single route that uses goroutines to count to X. The route should run a given number of goroutines based 
on the count that was passed in, each routine should pause for 1 second and then add to the channel before finishing. Once all the routines
have finished the API should return a success message.

The API should:
    1) Reject all other endpoints gracefully
    2) Only accept POST requests
    3) The body of the request should be JSON and look like ```{"count": xx}``` where xx is an arbitrary number.
    4) Use channels
    5) Use goroutines


##
# Notes
##

Have some fun with this and make it your own.
Include lots of comments so that we can get a feel for your thought process while working through this task.