#Grabbit-cli
Commandline tool for checking rabbitmq Data.
Result will be formatted for nagios or Hostmonitor test.

Binaries for windows and linux are available.

# Quickstart

Assuming you are using the tool on the rabbitmq server and there is a queue called "myQueue1"

    ./grabbit-cli -check queue -queue myQueue1

Will output:

    Name: Name                                     Value: myQueue1
    Name: Vhost                                    Value: /
    Name: Durable                                  Value: true
    Name: AutoDelete                               Value: false
    Name: Arguments                                Value: map[]
    Name: Node                                     Value: rabbit@rabbit1
    Name: Status                                   Value: 
    Name: Memory                                   Value: 13912
    Name: Consumers                                Value: 0
    Name: ExclusiveConsumerTag                     Value: 
    Name: Policy                                   Value: 
    Name: Messages                                 Value: 0
    Name: MessagesDetails                          Value: {0}
    Name: MessagesReady                            Value: 0
    Name: MessagesReadyDetails                     Value: {0}
    Name: MessagesUnacknowledged                   Value: 0
    Name: MessagesUnacknowledgedDetails            Value: {0}
    Name: MessageStats                             Value: {0 {0}}
    Name: OwnerPidDetails                          Value: { 0 }
    Name: BackingQueueStatus                       Value: {0 0 0 0 0 0 0 0 0 0 0 0 0}
    UNKNOWN - Field: print, value: , Error: Just printing values

Ok, let's create a check for the unacked messages:
    
    ./grabbit-cli -check queue -queue myQueue1 -field MessagesUnacknowledged
    UNKNOWN - Field: MessagesUnacknowledged, value: 0, Error: <nil>

Ok, we have to define some levels.
Let's say 10 or less unack'ed messages are ok, 11 - 20 are not good (warning state), more than 20 are really bad

    ./grabbit-cli -check queue -queue myQueue1 -field MessagesUnacknowledged -ok 11 -warn 21 -error 1000000000 -operator "<"
    OK - Field: MessagesUnacknowledged, value: 0

But we don't wont nagios output. Use Advanced Hostmonitor output (warnlevel removed as it is not used)

    ./grabbit-cli -check queue -queue myQueue1 -field MessagesUnacknowledged -ok 11 -error 1000000000 -operator "<" -output=h
    scriptres:Ok:0

## Parameters

    Usage of ./grabbit-cli:
      -check="": What to check: [q]ueue,[h]ost,[e]xchange
      -error="": Error value
      -exchange="": Exchange Name
      -field="print": name of the field to return or print to get a list
      -host="http://127.0.0.1:15672": AMQP Host
      -nodeIndex=0: Node index (for cluster checks)
      -ok="": OK value
      -operator="=": Operator for check: [value] [operator] [level]
      -output="nagios": Format output for 'nagios', 'hostmonitor'
      -pass="guest": password
      -queue="": Name of the queue to check
      -user="guest": username
      -vhost="/": vhostname
      -warn="": Warning value



## Example
### Hostmonitor output
    
    # check memory
    ./grabbit-cli -check q -queue myQueue1 -field Memory -operator "<" -ok 14000 -error 140000 -output=h
    scriptres:Ok:13912

    # check message count
    ./grabbit-cli -check queue -queue myQueue1 -field Messages -ok 70 -error 10000000 -operator "<" -output=hostmonitor
    scriptres:Ok:0

    # check DiskFree alarm
    ./grabbit-cli -check host -field DiskFreeAlarm -ok false -error true -output=hostmonitor
    scriptres:Ok:false

    # check isRunning
    ./grabbit-cli -check host -field isRunning -ok true -error false -output=hostmonitor
    scriptres:Ok:true

    # check Consumers 1 or more are ok.
    ./grabbit-cli -check q -queue myQueue1 -field Consumers -ok 0 -error -1 -operator ">" -output=hostmonitor
    scriptres:Bad:0

### Nagios output

    # check memory
    ./grabbit-cli -check q -queue myQueue1 -field Memory -operator "<" -ok 14000 -error 140000
    OK - Field: Memory, value: 13912

    # check message count
    ./grabbit-cli -check queue -queue myQueue1 -field Messages -ok 70 -error 10000000 -operator "<" 
    OK - Field: Messages, value: 0

    # check DiskFree alarm
    ./grabbit-cli -check host -field DiskFreeAlarm -ok false -error true 
    OK - Field: DiskFreeAlarm, value: false

    # check isRunning
    ./grabbit-cli -check host -field isRunning -ok true -error false 
    OK - Field: isRunning, value: true

    # check Consumers 1 or more are ok.
    ./grabbit-cli -check q -queue myQueue1 -field Consumers -ok 0 -error -1 -operator ">" 
    CRITICAL - Field: Consumers, value: 0


# Special thanks
grabbit-cli is using rabbit-hole for getting the data:
[http://github.com/michaelklishin/rabbit-hole](http://github.com/michaelklishin/rabbit-hole)

# Licence
BSD license.

(c) Kris Budde, 2014.