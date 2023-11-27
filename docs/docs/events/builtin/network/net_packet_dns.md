## DNS

The Domain Name System (DNS) is a hierarchical and distributed naming system
used on the Internet to translate human-friendly domain names (e.g.,
www.example.com) into IP addresses (e.g., 192.0.2.1) that computers use to
identify each other on the network. The DNS operates using a client-server
architecture, where DNS clients (usually resolvers) send DNS queries to DNS
servers (often called name servers) to resolve domain names. The DNS header is a
fundamental part of the DNS protocol, containing information about the DNS query
or response being exchanged.

Here's an overview of the DNS header and its significance:

```
 0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| Transaction ID (16 bits)                                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| Flags (16 bits)                                               |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| Question Count (16 bits)                                      |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| Answer Count (16 bits)                                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| Authority Record Count (16 bits)                              |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
| Additional Record Count (16 bits)                             |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

The DNS header consists of several fields, each serving a specific purpose in
the DNS query or response:

1. **Transaction ID (16 bits):** The Transaction ID is a unique identifier for each DNS query. It allows DNS clients to match responses to the corresponding queries.

2. **Flags (16 bits):** The Flags field contains various control and configuration options, including:
     - **QR** (1 bit): Indicates whether the message is a query (0) or a response (1).
     - **Opcode** (4 bits): Specifies the type of query (e.g., standard query, inverse query, status request).
     - **AA** (1 bit): Indicates if the responding name server is authoritative for the queried domain.
     - **TC** (1 bit): Signifies if the message is truncated due to its size.
     - **RD** (1 bit): Requests recursive resolution from the DNS server.
     - **RA** (1 bit): Indicates that the DNS server supports recursive queries.
     - **Z** (3 bits): Reserved for future use.
     - **RCODE** (4 bits): Represents the response code, indicating the status of the query (e.g., success, name error, server failure).

3. **Question Count (16 bits):** Specifies the number of questions (queries) in the DNS message.

4. **Answer Count (16 bits):** Indicates the number of resource records in the answer section.

5. **Authority Record Count (16 bits):** Specifies the number of authority resource records.

6. **Additional Record Count (16 bits):** Indicates the number of additional resource records.

The DNS protocol is a critical component of Internet infrastructure, enabling
users and applications to access websites and services using human-readable
domain names. When a user enters a domain name into a web browser or other
application, the DNS resolver on the user's device initiates a DNS query by
creating a DNS message with the appropriate header fields.

DNS queries are typically sent to a local DNS resolver (e.g., provided by an
ISP), which may cache previous DNS resolutions or forward the query to
authoritative name servers responsible for the queried domain. Recursive queries
are used to navigate the DNS hierarchy until the final authoritative name server
is reached, and an IP address is returned to the client.

DNS responses include the IP address associated with the queried domain name,
and optionally, additional information such as the Time-to-Live (TTL) value and
other resource records (RRs) like IPv6 addresses or mail server details.

The DNS header plays a crucial role in facilitating efficient and accurate DNS
resolution. It helps identify the type of DNS message, the number of questions
and resource records, and the status of the DNS query or response. The DNS
protocol operates transparently for most users, but it is fundamental to the
functioning of the Internet, making it possible for users to access websites and
services using user-friendly domain names rather than raw IP addresses.

### net_packet_dns

The `net_packet_dns` event provides one event for each existing DNS packet that
reaches or leaves one of the processes being traced (or even "all OS processes
for the default run"). As arguments for this event you will find: `src`, `dst`,
`src_port`, `dst_port`, `metadata` arguments and all `DNS header fields`.

Example:

```
$ tracee --output json --events net_packet_dns
```

```json
{"timestamp":1696259024822467299,"threadStartTime":1696259024820530450,"processorId":0,"processId":1053474,"cgroupId":5650,"threadId":1053476,"parentProcessId":1037836,"hostProcessId":1053474,"hostThreadId":1053476,"hostParentProcessId":1037836,"userId":1000,"mountNamespace":4026531841,"pidNamespace":4026531836,"processName":"isc-net-0000","executable":{"path":""},"hostName":"rugged","containerId":"","container":{},"kubernetes":{},"eventId":"2006","eventName":"net_packet_dns","matchedPolicies":[""],"argsNum":5,"returnValue":0,"syscall":"sendmmsg","stackAddresses":[0],"contextFlags":{"containerStarted":false,"isCompat":false},"threadEntityId":2326406626,"processEntityId":2231131033,"parentEntityId":2142180145,"args":[{"name":"src","type":"const char*","value":"127.0.0.1"},{"name":"dst","type":"const char*","value":"127.0.0.53"},{"name":"src_port","type":"u16","value":44493},{"name":"dst_port","type":"u16","value":53},{"name":"proto_dns","type":"trace.ProtoDNS","value":{"ID":60318,"QR":0,"opCode":"query","AA":0,"TC":0,"RD":1,"RA":0,"Z":0,"responseCode":"no error","QDCount":1,"ANCount":0,"NSCount":0,"ARCount":0,"questions":[{"name":"www.uol.com.br","type":"A","class":"IN"}],"answers":[],"authorities":[],"additionals":[]}}]}
{"timestamp":1696259024822806573,"threadStartTime":1695658999333342120,"processorId":4,"processId":472,"cgroupId":2626,"threadId":472,"parentProcessId":1,"hostProcessId":472,"hostThreadId":472,"hostParentProcessId":1,"userId":976,"mountNamespace":4026532555,"pidNamespace":4026531836,"processName":"systemd-resolve","executable":{"path":""},"hostName":"rugged","containerId":"","container":{},"kubernetes":{},"eventId":"2006","eventName":"net_packet_dns","matchedPolicies":[""],"argsNum":5,"returnValue":0,"syscall":"write","stackAddresses":[0],"contextFlags":{"containerStarted":false,"isCompat":false},"threadEntityId":131662446,"processEntityId":131662446,"parentEntityId":1975426032,"args":[{"name":"src","type":"const char*","value":"192.168.200.50"},{"name":"dst","type":"const char*","value":"8.8.8.8"},{"name":"src_port","type":"u16","value":47508},{"name":"dst_port","type":"u16","value":53},{"name":"proto_dns","type":"trace.ProtoDNS","value":{"ID":62897,"QR":0,"opCode":"query","AA":0,"TC":0,"RD":1,"RA":0,"Z":0,"responseCode":"no error","QDCount":1,"ANCount":0,"NSCount":0,"ARCount":1,"questions":[{"name":"www.uol.com.br","type":"A","class":"IN"}],"answers":[],"authorities":[],"additionals":[{"name":"","type":"OPT","class":"Unknown","TTL":0,"IP":"","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""}]}}]}
{"timestamp":1696259024822893266,"threadStartTime":1695658999333342120,"processorId":4,"processId":472,"cgroupId":2626,"threadId":472,"parentProcessId":1,"hostProcessId":472,"hostThreadId":472,"hostParentProcessId":1,"userId":976,"mountNamespace":4026532555,"pidNamespace":4026531836,"processName":"systemd-resolve","executable":{"path":""},"hostName":"rugged","containerId":"","container":{},"kubernetes":{},"eventId":"2006","eventName":"net_packet_dns","matchedPolicies":[""],"argsNum":5,"returnValue":0,"syscall":"write","stackAddresses":[0],"contextFlags":{"containerStarted":false,"isCompat":false},"threadEntityId":131662446,"processEntityId":131662446,"parentEntityId":1975426032,"args":[{"name":"src","type":"const char*","value":"192.168.200.50"},{"name":"dst","type":"const char*","value":"1.1.1.1"},{"name":"src_port","type":"u16","value":37385},{"name":"dst_port","type":"u16","value":53},{"name":"proto_dns","type":"trace.ProtoDNS","value":{"ID":35323,"QR":0,"opCode":"query","AA":0,"TC":0,"RD":1,"RA":0,"Z":0,"responseCode":"no error","QDCount":1,"ANCount":0,"NSCount":0,"ARCount":1,"questions":[{"name":"www.uol.com.br","type":"A","class":"IN"}],"answers":[],"authorities":[],"additionals":[{"name":"","type":"OPT","class":"Unknown","TTL":0,"IP":"","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""}]}}]}
{"timestamp":1696259024854662413,"threadStartTime":1695658999333342120,"processorId":6,"processId":472,"cgroupId":2626,"threadId":472,"parentProcessId":1,"hostProcessId":472,"hostThreadId":472,"hostParentProcessId":1,"userId":976,"mountNamespace":4026532555,"pidNamespace":4026531836,"processName":"systemd-resolve","executable":{"path":""},"hostName":"rugged","containerId":"","container":{},"kubernetes":{},"eventId":"2006","eventName":"net_packet_dns","matchedPolicies":[""],"argsNum":5,"returnValue":0,"syscall":"","stackAddresses":[0],"contextFlags":{"containerStarted":false,"isCompat":false},"threadEntityId":131662446,"processEntityId":131662446,"parentEntityId":1975426032,"args":[{"name":"src","type":"const char*","value":"8.8.8.8"},{"name":"dst","type":"const char*","value":"192.168.200.50"},{"name":"src_port","type":"u16","value":53},{"name":"dst_port","type":"u16","value":47508},{"name":"proto_dns","type":"trace.ProtoDNS","value":{"ID":62897,"QR":1,"opCode":"query","AA":0,"TC":0,"RD":1,"RA":1,"Z":0,"responseCode":"no error","QDCount":1,"ANCount":5,"NSCount":0,"ARCount":1,"questions":[{"name":"www.uol.com.br","type":"A","class":"IN"}],"answers":[{"name":"www.uol.com.br","type":"CNAME","class":"IN","TTL":49,"IP":"","NS":"","CNAME":"dftex7xfha8fh.cloudfront.net","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.81","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.15","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.88","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.16","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""}],"authorities":[],"additionals":[{"name":"","type":"OPT","class":"Unknown","TTL":0,"IP":"","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""}]}}]}
{"timestamp":1696259024855173520,"threadStartTime":1695658999333342120,"processorId":4,"processId":472,"cgroupId":2626,"threadId":472,"parentProcessId":1,"hostProcessId":472,"hostThreadId":472,"hostParentProcessId":1,"userId":976,"mountNamespace":4026532555,"pidNamespace":4026531836,"processName":"systemd-resolve","executable":{"path":""},"hostName":"rugged","containerId":"","container":{},"kubernetes":{},"eventId":"2006","eventName":"net_packet_dns","matchedPolicies":[""],"argsNum":5,"returnValue":0,"syscall":"sendmsg","stackAddresses":[0],"contextFlags":{"containerStarted":false,"isCompat":false},"threadEntityId":131662446,"processEntityId":131662446,"parentEntityId":1975426032,"args":[{"name":"src","type":"const char*","value":"127.0.0.53"},{"name":"dst","type":"const char*","value":"127.0.0.1"},{"name":"src_port","type":"u16","value":53},{"name":"dst_port","type":"u16","value":44493},{"name":"proto_dns","type":"trace.ProtoDNS","value":{"ID":60318,"QR":1,"opCode":"query","AA":0,"TC":0,"RD":1,"RA":1,"Z":0,"responseCode":"no error","QDCount":1,"ANCount":5,"NSCount":0,"ARCount":0,"questions":[{"name":"www.uol.com.br","type":"A","class":"IN"}],"answers":[{"name":"www.uol.com.br","type":"CNAME","class":"IN","TTL":49,"IP":"","NS":"","CNAME":"dftex7xfha8fh.cloudfront.net","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.81","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.15","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.88","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.16","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""}],"authorities":[],"additionals":[]}}]}
{"timestamp":1696259024855201893,"threadStartTime":1696259024820530450,"processorId":4,"processId":1053474,"cgroupId":5650,"threadId":1053476,"parentProcessId":1037836,"hostProcessId":1053474,"hostThreadId":1053476,"hostParentProcessId":1037836,"userId":1000,"mountNamespace":4026531841,"pidNamespace":4026531836,"processName":"isc-net-0000","executable":{"path":""},"hostName":"rugged","containerId":"","container":{},"kubernetes":{},"eventId":"2006","eventName":"net_packet_dns","matchedPolicies":[""],"argsNum":5,"returnValue":0,"syscall":"","stackAddresses":[0],"contextFlags":{"containerStarted":false,"isCompat":false},"threadEntityId":2326406626,"processEntityId":2231131033,"parentEntityId":2142180145,"args":[{"name":"src","type":"const char*","value":"127.0.0.53"},{"name":"dst","type":"const char*","value":"127.0.0.1"},{"name":"src_port","type":"u16","value":53},{"name":"dst_port","type":"u16","value":44493},{"name":"proto_dns","type":"trace.ProtoDNS","value":{"ID":60318,"QR":1,"opCode":"query","AA":0,"TC":0,"RD":1,"RA":1,"Z":0,"responseCode":"no error","QDCount":1,"ANCount":5,"NSCount":0,"ARCount":0,"questions":[{"name":"www.uol.com.br","type":"A","class":"IN"}],"answers":[{"name":"www.uol.com.br","type":"CNAME","class":"IN","TTL":49,"IP":"","NS":"","CNAME":"dftex7xfha8fh.cloudfront.net","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.81","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.15","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.88","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""},{"name":"dftex7xfha8fh.cloudfront.net","type":"A","class":"IN","TTL":60,"IP":"108.139.182.16","NS":"","CNAME":"","PTR":"","TXTs":null,"SOA":{"MName":"","RName":"","serial":0,"refresh":0,"retry":0,"expire":0,"minimum":0},"SRV":{"priority":0,"weight":0,"port":0,"name":""},"MX":{"preference":0,"name":""},"OPT":[],"URI":{"priority":0,"weight":0,"target":""},"TXT":""}],"authorities":[],"additionals":[]}}]}
{"timestamp":1696259024855756036,"threadStartTime":1696259024820530450,"processorId":1,"processId":1053474,"cgroupId":5650,"threadId":1053476,"parentProcessId":1037836,"hostProcessId":1053474,"hostThreadId":1053476,"hostParentProcessId":1037836,"userId":1000,"mountNamespace":4026531841,"pidNamespace":4026531836,"processName":"isc-net-0000","executable":{"path":""},"hostName":"rugged","containerId":"","container":{},"kubernetes":{},"eventId":"2006","eventName":"net_packet_dns","matchedPolicies":[""],"argsNum":5,"returnValue":0,"syscall":"sendmmsg","stackAddresses":[0],"contextFlags":{"containerStarted":false,"isCompat":false},"threadEntityId":2326406626,"processEntityId":2231131033,"parentEntityId":2142180145,"args":[{"name":"src","type":"const char*","value":"127.0.0.1"},{"name":"dst","type":"const char*","value":"127.0.0.53"},{"name":"src_port","type":"u16","value":53879},{"name":"dst_port","type":"u16","value":53},{"name":"proto_dns","type":"trace.ProtoDNS","value":{"ID":41668,"QR":0,"opCode":"query","AA":0,"TC":0,"RD":1,"RA":0,"Z":0,"responseCode":"no error","QDCount":1,"ANCount":0,"NSCount":0,"ARCount":0,"questions":[{"name":"dftex7xfha8fh.cloudfront.net","type":"AAAA","class":"IN"}],"answers":[],"authorities":[],"additionals":[]}}]}
```