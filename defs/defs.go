package defs

// MaxRequestSize is the maximum size (in bytes) of any request that Sarama will attempt to send. Trying
// to send a request larger than this will result in an PacketEncodingError. The default of 100 MiB is aligned
// with Kafka's default `socket.request.max.bytes`, which is the largest request the broker will attempt
// to process.
var MaxRequestSize int32 = 100 * 1024 * 1024

// MaxResponseSize is the maximum size (in bytes) of any response that Sarama will attempt to parse. If
// a broker returns a response message larger than this value, Sarama will return a PacketDecodingError to
// protect the client from running out of memory. Please note that brokers do not have any natural limit on
// the size of responses they send. In particular, they can send arbitrarily large fetch responses to consumers
// (see https://issues.apache.org/jira/browse/KAFKA-2063).
var MaxResponseSize int32 = 100 * 1024 * 1024
