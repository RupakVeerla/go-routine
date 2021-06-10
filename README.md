# go-routine problem statement

Write a program which reads full path to the directory and return statistic about how many files this directory have grouped by file extension
-	the initial function should start with file descriptor
-	create the buffered channel as a value you can use map[string]int
-	create wait group (sync.WaitGroup)
-	go for each child item 
-	if file descriptor referring to file, increment a local counter for file extension
-	if file descriptor referring to the directory, increment wait group counter, start goroutine and pass result channel as a parameter to it, 
	the function should make the same operations go over subitems and depends on if it’s directory or file increment counter or start the daemon
-	when child daemon is over decrement wait group counter

-	read and merge results from the channel, when all child item is processed return summarized result to the caller

As a result, you will get a recursive program which can concurrently scrab statistic and print it to console