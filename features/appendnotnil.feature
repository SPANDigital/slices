Feature: AppendNotNil

Scenario: Append elements to a slice
  Given a string ptr slice with elements
    | "a"     |
    | "b"     |
    | nil     |
    | "c"     |
    | nil     |
  When I call AppendNotNil on the slice with the following variable arguments
    | "d"     |
    | "e"     |
    | nil     |
    | "f"     |
  Then the result should be a string ptr slice with elements
    | "a"     |
    | "b"     |
    | nil     |
    | "c"     |
    | nil     |
    | "d"     |
    | "e"     |
    | "f"     |

  Scenario: Append non-nil elements to an empty slice
    Given an empty string ptr slice
    When I call AppendNotNil on the slice with the following variable arguments
      | "d"     |
      | "e"     |
      | nil     |
      | "f"     |
    Then the result should be a string ptr slice with elements
      | "d"     |
      | "e"     |
      | "f"     |

  Scenario: Append non-nil elements to a slice with all nil elements
    Given a string ptr slice with elements
      | nil     |
      | nil     |
      | nil     |
    When I call AppendNotNil on the slice with the following variable arguments
      | "x" |
      | "y" |
      | nil |
      | "z" |
    Then the result should be a string ptr slice with elements
      | nil |
      | nil |
      | nil |
      | "x" |
      | "y" |
      | "z" |
