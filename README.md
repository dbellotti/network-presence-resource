# network-presence-resource

This is a [concourse.ci](http://concourse.ci) resource whose `check` behavior
determines if registered IP addresses are on a network where the `geo-agent`
daemon is running. If the previous state is different than the current state,
it's `version` has changed and `in` will be triggered.  There is no `out`
behavior.

Sort of like a launch control switch a job could be triggered only if there
were multiple IP's present from the registered set.

#### Finding out your network and device IP
To find your local IP, use
```
ifconfig
```

You will most likely want to use your phone as the trigger device (because
a computer is probably always going to be on the network) so get the network's
/24 and use a tool like nmap to find your device's IP
IP's can be obtained by running
```
nmap -sP [your-local-network-24]
```

This value can be hardcoded as one of the IP's that the geo-agent daemon will
check when invoked.
