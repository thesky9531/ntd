# NTD - Network Traffic Detection v1.0.0

By [TheSky](http://blog.xinjie.wang)

The NTD is for collects network flow in internet, It's operation depends on Pcap.

If your system in windows, you could need winpcap to run it. And system is Linux, you need libpcap.

The NTD's inspiration is originate from  Beats for elastic, Thanks. 


## inputs
NTD's input source from network's data flow packet, and depend in winpcap/libpcap.

Here is more details about pacp. 

if you will custom to set inputs of pcap, you should set BPF_Filter in the *conf.yml*,
for example: 
- "tcp" 
 - "udp" 
  - "tcp or udp"
 
 And if you not set BPF filter, we may be set the BPF filter "" (string), Then, The pcap 
 collect all flow packet in your network.


## handle
NTD's file is due to flow packet for real env

Here are more configuration guidelines for handle in the configuration file.

#### *handle*
     
     - Layer 1 LinkLayer
     - Layer 2 NetworkLayer
     - Layer 3 TransferLayer
     - Layer 4 ApplicationLayer
     
for example
## outputs
NTD's outputs is Customzable in conf.yml

Here are more configuration guidelines for output in the configuration file.

#### *outputs*

    - UDP -> syslog
    - DB -> mysql / postgres
    - TCP -> http / https / Tcp
    - I/O -> local File
    
    - stdout -> Console window

## about pcap
the winpcap/libpcap collect network flow packet is flowing your network, 

At windows system if your network is not set *promiscuous mode*, the pcap will only collect
flowing your computer. But Your may be alreay set the promiscous when you install
 the winpcap.
 
 At Linux system is simple reson to windows.
 
 especial example:
 
 The switchBoard limits your promiscuous mode, you need copy
 the image traffic into your network.
 
 ## about 
