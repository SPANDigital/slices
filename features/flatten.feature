Feature: Flatten slices
  Scenario: Flatten a slice of slices with mixed elements
    Given a slice of integer slices with elements
      |   |   |   |
      | 1 |   |   |
      | 2 | 3 |   |
      | 4 | 5 | 6 |
    When I call Flatten on the slice of integer slices
    Then the result should be an integer slice with elements
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
      | 6 |

  Scenario: Flatten a slice of slices with empty sub-slices
    Given a slice of integer slices with elements
      | |
      | |
      | |
    When I call Flatten on the slice of integer slices
    Then the result should be an empty integer slice

  Scenario: Flatten an empty slice of slices
    Given an empty slice of integer slices
    When I call Flatten on the slice of integer slices
    Then the result should be an empty integer slice
