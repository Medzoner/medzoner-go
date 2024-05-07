Feature: Global behavior "Medzoner"
    In order to check the contact behavior APP
    As a visitor
    I need to able to access

    Background:
        And I add "Content-Type" header equal to "text/html"

#------------------------------------------------------------------------------------------
# GET "CONTACT" - Test succeeded
#------------------------------------------------------------------------------------------

    Scenario: [Medzoner - GET_ALL] "Contact page"
        When    I send a GET request to "/contact"
        Then    the response status code should be 405

#------------------------------------------------------------------------------------------
# POST "CONTACT" - Test success
#------------------------------------------------------------------------------------------

    Scenario: [Medzoner - POST] "Contact page - Test success"
        And I add "Content-Type" header equal to "text/html"
        When    I send a POST request to "/contact" with body:
          """
          {"name":"else", "email":"email@fake.com", "message":"else"}
          """
        Then    the response status code should be 303

#------------------------------------------------------------------------------------------
# POST "CONTACT" - Test failed
#------------------------------------------------------------------------------------------

    Scenario: [Medzoner - POST] "Contact page - Test failed"
        And I add "Content-Type" header equal to "application/x-www-form-urlencoded"
        When    I send a POST request to "/contact" with body:
          """
          {"foo":{"bar":{"baz":1}},"something":"else"}
          """
        Then    the response status code should be 400
