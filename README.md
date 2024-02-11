# Bodega Filters

![build](https://github.com/jadidbourbaki/bodega/actions/workflows/build.yml/badge.svg)


This is a prototype implementation of the Uptown Bodega Filter and the Downtown Bodega Filter in the steady setting with a mock learning model. For real applications, the mock learning model may be replaced by a real learning model. We release this implementation with the following adversarial security challenge:

After $T$ queries to a sufficiently provisioned Uptown Bodega Filter, construct a query that generates a false positive with probability $P$ such that the difference $P - P_{FP}$ is not negligible, $P_{FP}$ being the expected false positive probability of a [Sandwiched Learning Bloom Filter](https://arxiv.org/abs/1901.00902) with the provisioned resources.
