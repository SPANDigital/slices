Feature: Index function

  Scenario: Element present at the beginning of the slice
    Given a string slice
      | a |
      | b |
      | c |
    When I call Index on the string slice with argument "a"
    Then the result should be 0

  Scenario: Element present in the middle of the slice
    Given a string slice
      | a |
      | b |
      | c |
    When I call Index on the string slice with argument "b"
    Then the result should be 1

  Scenario: Element present at the end of the slice
    Given a string slice
      | a |
      | b |
      | c |
    When I call Index on the string slice with argument "c"
    Then the result should be 2

  Scenario: Element not present in the slice
    Given a string slice
      | a |
      | b |
      | c |
    When I call Index on the string slice with argument "z"
    Then the result should be -1
