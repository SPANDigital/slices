Feature: GroupByLen

  Scenario: Grouping a slice of strings into pairs
    Given a slice of strings with elements
      | a |
      | b |
      | c |
      | d |
    When I call GroupByLen on the strings slice with length 2
    Then the result should be a slice of string slices with elements
      | a | b |
      | c | d |

  Scenario: Grouping a slice of strings into triples
    Given a slice of strings with elements
      | a |
      | b |
      | c |
      | d |
      | e |
      | f |
    When I call GroupByLen on the strings slice with length 3
    Then the result should be a slice of string slices with elements
      | a | b | c |
      | d | e | f |

  Scenario: Grouping a slice of strings with indivisible length
    Given a slice of strings with elements
      | a |
      | b |
      | c |
      | d |
      | e |
    When I call GroupByLen on the strings slice with length 3
    Then the result should be a slice of string slices with elements trim cells on right
      | a | b | c |
      | d | e |   |
