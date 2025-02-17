Feature: Filter function

  Scenario: Filter positive integers
    Given an integer slice with elements
      | 1  |
      | -1 |
      | 2  |
      | -2 |
      | 3  |
      | -3 |
    When I call Filter with a function that returns true for positive integers
    Then the result should be an integer slice with elements
      | 1 |
      | 2 |
      | 3 |

  Scenario: Filter negative numbers
    Given an integer slice with elements
      | 1  |
      | -1 |
      | 2  |
      | -2 |
      | 3  |
      | -3 |
    When I call Filter with a function that returns true for negative integers
    Then the result should be an integer slice with elements
      | -1 |
      | -2 |
      | -3 |
