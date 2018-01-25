go-metrics-honeycomb
--------------------

This is a simple reporter for rcrowley's
[go-metrics](https://github.com/rcrowley/go-metrics) library, designed to work
with the [Honeycomb](https://honeycomb.io/) metrics reporting service.

Usage
=====

Basic usage:

```golang

import (
  "time"

  "github.com/getspine/go-metrics-honeycomb"
  "github.com/rcrowley/go-metrics"
)

go honeycomb.Honeycomb(
  metrics.DefaultRegistry,
  60 * time.Second, // Interval between sending metrics
  "your-write-key", // Honeycomb write key
  "your-dataset",   // Honeycomb dataset
)
```

If you wish to add extra fields or other configuration to the Honeycomb client, you
can use the following procedure:

```golang
import (
  "time"

  "github.com/getspine/go-metrics-honeycomb"
  "github.com/honeycombio/libhoney-go"
  "github.com/rcrowley/go-metrics"
)

reporter := honeycomb.NewReporter(
  metrics.DefaultRegistry,
  60 * time.Second, // Interval between sending metrics
  "your-write-key", // Honeycomb write key
  "your-dataset",   // Honeycomb dataset
)

libhoney.AddField("philcollins", "sussudio")
...

go reporter.Run()
```

Installation
============

```bash
$ go get github.com/getspine/go-metrics-honeycomb
```

Bugs
====

If you run into any bugs, please drop an issue on our GitHub and we'll be sure
to have a look.
