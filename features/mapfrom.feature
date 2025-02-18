Feature: MapFrom function

  Scenario: Map integers to their double and triple values
    Given a slice of integers
      | 1  |
      | 2  |
      | 3  |
      | 4  |
      | 5  |
      | 6  |
      | 7  |
      | 8  |
      | 9  |
      | 10 |
    When I call MapFrom to convert the integers to their double and triple values
    Then the result should be a map of integer keys and integer values
      | 2  | 3  |
      | 4  | 6  |
      | 6  | 9  |
      | 8  | 12 |
      | 10 | 15 |
      | 12 | 18 |
      | 14 | 21 |
      | 16 | 24 |
      | 18 | 27 |
      | 20 | 30 |
