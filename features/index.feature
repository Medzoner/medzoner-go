Feature: Global behavior "Medzoner"
    In order to check the global behavior APP
    As a visitor
    I need to able to access

    Background:
        And I add "Authorization" header equal to ""

#------------------------------------------------------------------------------------------
# GET "Home" - Test succeeded
#------------------------------------------------------------------------------------------

    Scenario: [Medzoner - GET_ALL] "Home page"
        When    I send a GET request to ""
        Then    the response status code should be 200
