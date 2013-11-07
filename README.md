gographite
==========

gographite is a super simple way to put data into a running <a href="http://graphite.readthedocs.org/en/latest/">graphite</a> instance.


What?
-----

Start gographite either with Go or a binary that you have already compiled.

```
$ go run main.go -l 9999 -v
2013/11/07 10:21:58 Starting server process on port 9999
2013/11/07 10:21:58 Connecting to carbon-cache process at (tcp) 127.0.0.1:2003
```

Feed it some data.

```
T=`date +%s` curl "http://127.0.0.1:9999/localhost/random/bit/of/stats/$T/123"
```

If you started gographite in verbose mode, you will see a line similar to this:
```
2013/11/07 10:46:36 sending localhost.random.bit.of.stats 123.000000 1383849959
```

A more practical example. Log the number of packets sent by eth0 every 10 seconds:
```
$ while true; do V=`cat /sys/class/net/eth0/statistics/tx_packets`; T=`date +%s`; curl "http://127.0.0.1:9999/localhost/net/tcp/tx_packets/${T}/${V}"; sleep 10 ; done
```

Using the graphite command line client you can see the results pretty quickly:
![Screenshot](https://raw.github.com/ceberly/gographite/master/readme.screenshot.png)

Why?
----
Good question. To learn about graphite, mostly.
gographite is totally language agnostic, by virtue of having no bindings at all. Almost every language (including the unix command line) has an easy wrapper around `curl`.
If you can turn it into a URL, you can put it into graphite.

How is this different from <a href="https://github.com/etsy/statsd/">statsd</a>, etc?
---------------------------------------
gographite is not very different from other graphite interfacing tools. Unlike statsd, gographite makes no attempt to interpret the data you are sending it. This can either be a good thing or a bad thing, depending on how you are using graphite. gographite can also be easily compiled into a binary, so you can deploy it without any other dependencies.

Licensing, etc.
---------------
See the accompanying LICENSE file for details. You are more than welcome to fork gographite and make it better.
