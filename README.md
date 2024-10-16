## Introduction
# This assignment helps you learn how to interact with the Twitter API using OAuth 2.0 for user authentication. You'll set up a Twitter Developer account, generate API keys, and create a program that can post and delete tweets. You'll also handle errors and invalid inputs when making API calls.
## Setup Instruction
 # Create a Twitter Developer Account
1.Go to the Twitter Developer Portal.
2.Sign in with your Twitter account.
3. Apply for a Developer account and complete the necessary details.
Once approved, log in to the portal.
## Generate API Keys
In the Developer Portal, go to Projects & Apps.
Click Create App and fill in the required details.
Once the app is created, go to Keys and Tokens to generate your API Key, API Key Secret, Bearer Token, and Access Tokens.
## Run Your Program
Ensure the API keys and tokens.
Install necessary libraries 
Run the program by executing the script in go .


## Explain how your program posts a new tweet.
## Posting a New Tweet
When posting a tweet, the program uses the Twitter API to send a request with the content of the tweet. This involves authenticating the user with OAuth 2.0 and making a POST request to the statuses/update endpoint.

# Step 1: Authenticate the user using OAuth 2.0 by providing the API Key, API Secret Key, Access Token, and Access Token Secret.
# Step 2: Prepare the tweet content that will be posted.
# Step 3: Send the POST request to the Twitter API with the tweet content.
# Response: The API will return a JSON response that includes details about the newly posted tweet, such as the tweet ID, creation time, and tweet text.

## Deleting an Existing Tweet
The deletion of a tweet requires sending a DELETE request to the Twitter API's statuses/destroy/{id} endpoint. The program must specify the tweet ID of the post that needs to be deleted.

Step 1: Retrieve the tweet ID of the tweet you want to delete.
Step 2: Authenticate the user with OAuth 2.0, similar to posting a tweet.
Step 3: Send a DELETE request to the Twitter API with the specific tweet ID.
Response: The API will return a confirmation in JSON format that the tweet has been deleted. If the tweet cannot be found (e.g., it was already deleted or the ID is invalid), the response will contain an error message.

## error handling
# The program handles errors by:

# Invalid Authentication: If API keys or tokens are incorrect, it shows an "Authentication failed" message.
# Rate Limiting: It detects if too many requests are made and informs the user to wait until the limit resets.
# Invalid Tweet ID: If a tweet ID doesn’t exist or the tweet is already deleted, it displays "Tweet not found."
# Invalid Inputs: It checks tweet content for issues (e.g., character limits) before making the request.
# General API Errors: Any other errors are caught, and the specific issue from Twitter is shown to the user.

