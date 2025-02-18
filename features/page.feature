Feature: Page function
  As a developer
  I want to retrieve a specific page of elements from a slice
  So that I can paginate the data correctly

  Scenario: Retrieve the first page from a slice with 6 elements and page size 2
    Given a slice of integers
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
      | 6 |
    And a page size of 2
    And a page index of 0
    When I retrieve the page
    Then the result should be a slice of integers
      | 1 |
      | 2 |

  Scenario: Retrieve the second page from a slice with 6 elements and page size 2
    Given a slice of integers
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
      | 6 |
    And a page size of 2
    And a page index of 1
    When I retrieve the page
    Then the result should be a slice of integers
      | 3 |
      | 4 |

  Scenario: Retrieve the third page from a slice with 6 elements and page size 2
    Given a slice of integers
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
      | 6 |
    And a page size of 2
    And a page index of 2
    When I retrieve the page
    Then the result should be a slice of integers
      | 5 |
      | 6 |

  Scenario: Retrieve a page from an empty slice
    Given an empty slice of integers
    And a page size of 2
    And a page index of 0
    When I retrieve the page
    Then the result should be an empty slice of integers

  Scenario: Retrieve a page with an out-of-bounds index
    Given a slice of integers
      | 1 |
      | 2 |
      | 3 |
      | 4 |
      | 5 |
      | 6 |
    And a page size of 2
    And a page index of 3
    When I retrieve the page
    Then the result should be an empty slice of integers
