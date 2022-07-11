**Author**: George Iwu

**Go Experience**: Very limited\*

\* I was on a team "maintaining" a Go microservice, which meant very occasional changes of a few lines. I also started a failed pet project in Go a few years ago. Neither helped a lot with solving this task.

## Omitted Extra Features

There's a lot of extra things that _could_ be implemented in a project like this, depending on the business and technical context. Below are ones that came on my mind in a few minutes of thinking about this.

### Logging & Metrics

Depending on the needs and conventions, in a real world scenario, there could be a need for logging and gathering some metrics about usage frequency, performance etc.

### Tracing

If the system has support for distributed tracing, or we want to introduce it, we should implement the usage of **correlation IDs** and/or **trace IDs**. This increases system transparency and makes debugging easier.

### Resilience Mechanisms

As remote calls have a tendency for failing for various reasons, one should also consider introducing some resilience when making them e.g.

- **Circuit Breakers**, if we see a risk associated with extra requests in case of failures
- **Retries**, in case a failure is just a flake
- **Back-offs**, usually **exponential**, to avoid making all the retries within a split second
- **Timeouts**, in case the service takes too long to respond

### Caching

If there was a need for faster response times and the business context allows it (doubt it's a case for account information), one could consider caching the responses.
