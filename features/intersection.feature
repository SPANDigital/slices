Feature: Intersection function

  Scenario: No common elements
    Given the following string slices
      | one     | three |
      | two     | four  |
    When I call Intersection on the slices
    Then the result should be an empty string slice

  Scenario: One common element
    Given the following string slices
      | one     | common |
      | two     | three  |
      | common  | four   |
    When I call Intersection on the slices
    Then the result should be string slice
      | common  |

  Scenario: Multiple common elements
    Given the following string slices
      | one            | common        |
      | two            | second-common |
      | common         | three         |
      | second-common  | four          |
    When I call Intersection on the slices
    Then the result should be string slice
      | common         |
      | second-common  |

  Scenario: Empty slices
    Given I have 2 empty string slices
    When I call Intersection on the slices
    Then the result should be an empty string slice
