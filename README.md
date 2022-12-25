# wordle-helper
go tool that helps find words for wordle, quordle

sample usage:

```
$ go run main.go
Please wait while the dictionary is being loaded...
Finished loading dictionary.

******************
Enter green, orange(optional): optio
Did not find any word that match your criteria
******************

******************
Enter green, orange(optional): option~~
Found these words:  [optional optioned optionee optionor]
******************

******************
Enter green, orange(optional): h~~~~ roe
Found these words:  [hermo heron heros hoers hoker holer homer honer hoper horae horde horme horse hover]
******************
```

In cli, you input one or two words.
If you input 1 word, it is considered green, meaning of which is explained below.
If you input two words separated by space, the second word is considered orange.

Green word means that the letters in the word are green tiled in the wordle/qourdle. Order of the letter is important. the tilda (`~`) mean that any letter can take that place.

Orange word means that the letters are orange tiled. Here the order of the letters are insignificant.


