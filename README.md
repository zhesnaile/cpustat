# cpustat
### Monitor CPU clocks on a linux system.

Imitates the behaviour of the following command while keeping track of the highest (individual) clock speed reached and the current average clock speed.
```
watch -n.5 "grep \"^[c]pu MHz\" /proc/cpuinfo | awk -F':' '{print $2 }'"
```
