# Cartographer

The Cartographer is an agent for our new “Nagios Killer” system. It is aimed at being “ohai on steroids”. For the time being, Cartographer will observe a system and report varying metrics and data. The minimum report will be on the command line, later may be a push or a pull over the network.

### Desired Functionality

- Drop-and-go - can be dropped onto a system and discover information about it.
- Pluggable - can be extended via plugins/modules
- Is smart about what it reports on
    - Does not report about a missing web server if the host simply does not run a web service.
    - Only reports database metrics if a database is installed.
