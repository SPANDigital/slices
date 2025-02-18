Feature: Map function

  Scenario: Apply function to each element in the slice
    Given a string slice
      | a |
      | b |
      | c |
    When I call Map on the string slice with a function that converts each element to uppercase
    Then the result should be string slice
      | A |
      | B |
      | C |

  Scenario: Apply function to an empty slice
    Given an empty string slice
    When I call Map on the string slice with a function that converts each element to uppercase
    Then the result should be an empty string slice

  Scenario: Apply function that appends a suffix to each element
    Given a string slice
      | one |
      | two |
      | three |
    When I call Map on the string slice with a function that appends "-suffix" to each element
    Then the result should be string slice
      | one-suffix |
      | two-suffix |
      | three-suffix |

  Scenario: Apply function that returns the length of each element
    Given a string slice
      | one |
      | two |
      | three |
    When I call Map on the string slice with a function that returns the length of each element
    Then the result should be a slice of integers
      | 3 |
      | 3 |
      | 5 |
