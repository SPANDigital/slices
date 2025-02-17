Feature: Contains function

  Scenario: Element is present at the beginning of the slice
    Given a string slice with elements
      | a |
      | b |
      | c |
    When I call Contains with argument "a"
    Then the result should be boolean true

  Scenario: Element is present in the middle of the slice
    Given a string slice with elements
      | a |
      | b |
      | c |
    When I call Contains with argument "b"
    Then the result should be boolean true

  Scenario: Element is present at the end of the slice
    Given a string slice with elements
      | a |
      | b |
      | c |
    When I call Contains with argument "c"
    Then the result should be boolean true

  Scenario: Element is not present in the slice
    Given a string slice with elements
      | a |
      | b |
      | c |
    When I call Contains with argument "z"
    Then the result should be boolean false
