Feature: Global behavior "Medzoner"
    In order to check the technos behavior APP
    As a visitor
    I need to able to access

    Background:
        And I add "Authorization" header equal to ""

#------------------------------------------------------------------------------------------
# GET "TECHNOS" - Test succeeded
#------------------------------------------------------------------------------------------

    Scenario: [Medzoner - GET_ALL] "Technos page"
        When    I send a GET request to "/technos"
        Then    the response status code should be 404
