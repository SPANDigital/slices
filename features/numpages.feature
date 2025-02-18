Feature: NumPages function
  As a developer
  I want to calculate the number of pages for a given slice and page size
  So that I can paginate the data correctly

  Scenario: Calculate number of pages for a slice with 6 elements and page size 2
    Given a slice of integers
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
      | 6 |
    And a page size of 2
    When I calculate the number of pages
    Then the result should be integer 3

  Scenario: Calculate number of pages for a slice with 5 elements and page size 2
    Given a slice of integers
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
    And a page size of 2
    When I calculate the number of pages
    Then the result should be integer 3

  Scenario: Calculate number of pages for an empty slice and page size 2
    Given an empty slice of integers
    And a page size of 2
    When I calculate the number of pages
    Then the result should be integer 0

  Scenario: Calculate number of pages for a slice with 7 elements and page size 3
    Given a slice of integers
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
      | 6 |
      | 7 |
    And a page size of 3
    When I calculate the number of pages
    Then the result should be integer 3
