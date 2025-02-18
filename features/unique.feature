Feature: Unique function

  Scenario: Empty slice
    Given an empty string slice
    When the Unique function is called and result sorted alphabetically
    Then the result should be an empty string slice

  Scenario: Slice with one duplicate
    Given a string slice
      | a |
      | b |
      | b |
    When the Unique function is called and result sorted alphabetically
    Then the result should be string slice
      | a |
      | b |

  Scenario: Slice with multiple duplicates
    Given a string slice
      | a |
      | b |
      | b |
      | c |
      | c |
      | d |
      | d |
    When the Unique function is called and result sorted alphabetically
    Then the result should be string slice
      | a |
      | b |
      | c |
      | d |
