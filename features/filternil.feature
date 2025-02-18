Feature: FilterNil function

  Scenario: Filter nil elements from a slice
    Given a string ptr slice with elements
      | "a" |
      | nil |
      | "b" |
      | nil |
      | "c" |
    When I call FilterNil on the string ptr slice
    Then the result should be a string ptr slice with elements
      | "a" |
      | "b" |
      | "c" |

  Scenario: Filter nil elements from a slice with all nil elements
    Given a string ptr slice with elements
      | nil |
      | nil |
      | nil |
    When I call FilterNil on the string ptr slice
    Then the result should be an empty string ptr slice

  Scenario: Filter nil elements from an empty slice
    Given an empty string ptr slice
    When I call FilterNil on the string ptr slice
    Then the result should be an empty string ptr slice
