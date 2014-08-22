package main

import (
	"flag"
)

var (
	flag_host     = flag.String("host", "http://127.0.0.1:15672", "AMQP Host")
	flag_username = flag.String("user", "guest", "username for AMQP Host")
	flag_password = flag.String("pass", "guest", "password for AMQP Host")
	flag_check    = flag.String("check", "", "What to check: [q]ueue,[h]ost,[e]xchange")
	flag_vhost    = flag.String("vhost", "/", "vhostname for queue and exchange tests")
	flag_field    = flag.String("field", "print", "name of the field to return or print to get a list of current values")
	//Settings for Queue check
	flag_queue = flag.String("queue", "", "Name of the queue to check")
	//settings for node check
	flag_nodeIndex = flag.Int("nodeIndex", 0, "Node index (for cluster checks)")
	//settings for exchange check
	flag_exchangeName = flag.String("exchange", "", "Exchange Name")
	//Check
	flag_okLevel      = flag.String("ok", "", "OK value/level")
	flag_warningLevel = flag.String("warn", "", "Warning value/level")
	flag_errorLevel   = flag.String("error", "", "Error value/level")
	flag_operator     = flag.String("operator", "=", "Operator for check: [value] [operator] [level]")
	//Output type
	flag_output = flag.String("output", "nagios", "Format output for 'nagios', 'hostmonitor'")
)

func init() {
	flag.Parse()

}

func main() {
	output := getOutputType(*flag_output)
	value, err := runCheck(*flag_check, *flag_host, *flag_username, *flag_password, *flag_vhost, *flag_field, *flag_queue, *flag_nodeIndex, *flag_exchangeName)
	result, err := Compare(value, err, *flag_operator, *flag_okLevel, *flag_warningLevel, *flag_errorLevel)
	printResults(value, err, result, *flag_field, output)
}
