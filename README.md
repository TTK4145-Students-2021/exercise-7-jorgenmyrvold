## Disclaimer
*Diclaimer: This program is mac specific*

# The program
The main program is written in Python and can be run using
```
$ python3 main.py 1
```
Make sure the filepaths for `count.txt`, `backup_alive.txt` and `terminal_script.txt` is correct. __This has to be the absolute filepath__ and it can be set at the bottom of main.py and at line 44 for the `terminal_script.txt`.

The argrument 1 at the end is to inform that this is the first time the program is started and that it should start as primary. The program creates it's own backup and checks if the backup is alive. If the program hasn't heard from the backup within the timeout (default 3 sec) it creates another backup. If the primary dies and the backup don't hear from the primary within the timeout it takes over as primary and creates another backup.

I first tried to run the program using the `Network-Go` module handed out, but i had trouble getting it to work properly as others also pointed out on discord. Because the deadline was closing in i decides to go for a solution in python where the communication is done using files. This might not be the perfect solution but it does the job. Though the solution uses files the priciples are the same and it will probably be transferable to a network solution as well