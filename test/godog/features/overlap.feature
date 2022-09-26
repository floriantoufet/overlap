Feature: Call the overlap CLI to get network's relations

  Scenario: [Positive-01]: Call the overlap CLI success with subset relation
    When I call the overlap CLI with 10.0.0.0/20 as first CIDR and 10.0.2.0/24 as second CIDR
    Then output result should be "subset"

  Scenario: [Positive-02]: Call the overlap CLI success with superset relation
    When I call the overlap CLI with 10.0.2.0/24 as first CIDR and 10.0.0.0/20 as second CIDR
    Then output result should be "superset"

  Scenario: [Positive-03]: Call the overlap CLI success with different relation
    When I call the overlap CLI with 10.0.2.0/24 as first CIDR and 10.0.3.0/24 as second CIDR
    Then output result should be "different"

  Scenario: [Positive-04]: Call the overlap CLI success with same relation
    When I call the overlap CLI with 10.0.2.0/24 as first CIDR and 10.0.2.10/24 as second CIDR
    Then output result should be "same"

  Scenario: [Negative-01]: Call the overlap CLI failed if invalid number of arguments
    When I call the overlap CLI with too much arguments
    Then error message should be "invalid number of arguments"

  Scenario: [Negative-02]: Call the overlap CLI failed if invalid CIDR of arguments
    When I call the overlap CLI with foo as first CIDR and 10.0.2.10/24 as second CIDR and call failed
    Then error message should be "first CIDR invalid: unable to parse CIDR"

  Scenario: [Negative-03]: Call the overlap CLI failed if first CIDR is not an IPv4
    When I call the overlap CLI with 2001:db8::/32 as first CIDR and 10.0.2.10/24 as second CIDR and call failed
    Then error message should be "first CIDR invalid: not an IPv4 IP"

  Scenario: [Negative-04]: Call the overlap CLI failed if second CIDR is not an IPv4
    When I call the overlap CLI with 10.0.2.10/24 as first CIDR and 2001:db8::/32 as second CIDR and call failed
    Then error message should be "first CIDR invalid: not an IPv4 IP"