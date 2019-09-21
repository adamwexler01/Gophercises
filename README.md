# Gophercises
Going through 20 exercises on how to use golang from https://gophercises.com 

# Exercise 1
This program creates a simulated quiz from a csv file in the format of question, answer. Pulling in information from the csv, checking whether the input matches the answer and runs a timer in the background. When the timer is up, the quiz ends with the score printed out on the bottom. You can provide your own csv file with the -csv flag and you can provide your own time limit with the -limit flag. The default for these inputs are (problems.csv & 5000ms).

# Exercise 2
This program creates a small web application that shortens urls. By providing a list of paths and their forwarded urls, you can have a "url short" that will forward you to a longer url path. This could be good for having a list of websites that you would want to go to often or save to a list. This web application can take a YAML object or a Map as input and the output will be a url to forward to. If the url doesn't exist, there is a fallback.

Future additions: 
- JSON object that reflects a list of curated urls
- being able to provide these values from a file (.txt, .json, .yml)
- Connecting this to a database and reading the urls from there
