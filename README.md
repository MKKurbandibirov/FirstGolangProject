FirstGolangProject
============================
Obscene Vacabulary Cheker
----------------------------
This application allows you to identify forbidden (from the corresponding file) words in the standard input stream and replace them with a character ```*```
### Example:
At the beginning, you need to create a file with forbidden words. The `forbidden_words.txt` contents:
*  _disgusting_
*  _unpleasant_
*  _ugly_
*  _bad_

Next, let's run our program (`go run vocabulary_checker.go`), and we will be asked to enter a file name with forbidden words:
>![image](https://user-images.githubusercontent.com/74917681/158790943-bce1a7f8-e842-48f6-a99e-6536fbaece42.png)

Let's enter the name of our file:
>![image](https://user-images.githubusercontent.com/74917681/158792052-17182d3d-65b2-44a4-9deb-fc397fd32a93.png)

Now we can enter our messages interactively and the program will output them corrected:
>![image](https://user-images.githubusercontent.com/74917681/158794730-0d96f0f1-9879-4cd1-856b-d115923e13fb.png)

CoffeeMachine
-------------------------
The app simulates a typical coffee machine. Just run it (`go run coffee_machine.go`)
