// Package snowflake is a network service for generating unique ID numbers at high scale with some simple guarantees.
// The first bit is unused sign bit.
// The second part consists of a 41-bit timestamp (milliseconds) whose value is the offset of the current time relative to a certain time.
// The 5 bits of the third and fourth parts represent data center and worker, and max value is 2^5 -1 = 31.
// The last part consists of 12 bits, its means the length of the serial number generated per millisecond per working node, a maximum of 2^12 -1 = 4095 IDs can be generated in the same millisecond.
// In a distributed environment, five-bit datacenter and worker mean that can deploy 31 datacenters, and each datacenter can deploy up to 31 nodes.
// The binary length of 41 bits is at most 2^41 -1 millisecond = 69 years. So the snowflake algorithm can be used for up to 69 years, In order to maximize the use of the algorithm, you should specify a start time for it.
package snowflake
