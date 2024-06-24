Channels in Golang are a way for a Go routine to receive 

Channels are for receiving data
Waitgroups are for pure synchronization

It seems that you should always use channels...
Are there some situations where WaitGroups are required???

WaitGroups aren't necessarily required, but they have a readability benefit

