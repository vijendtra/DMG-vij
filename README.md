Observations:

In this code I have noticed race condition were not handle properly.

Added two changes:

I  added sync.Mutex in Counter struct, so that each thread/request starts by locking the mutex
I have added pointer *Counter.So whenever data get modify should always be defined with pointer.
In brief mutexes must not be copied, so if this struct is passed around, it should be done by pointer.

