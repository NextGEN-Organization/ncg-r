The attached file generates an amount of nitro codes to a text file. Usage: ./main 50000000 output.txt  << will create 50 million 16 char strings to output.txt.
Note that output.txt must be pre-made.

If you beat this generation rate using code that's feasible for the average user to do: brownie points.
We will also compete since our file-writing requires further optimization.

Conditions:
```diff
+ **50 million 16 character long strings** within the set **a-zA-Z0-9** should be generated.

+ All of the generated strings must be written to a file.
note: the file can pre-exist. Creating a file is unnecessary but brownie points if you do anyway.

+ Noise variability should be based off of int64.
Standard noise libraries are recommended. 
If you use your own custom solution, we will have to look at it to make sure it isn't cutting corners. 
(This is due to int8/16 ranges showing signs of favoritism during our own testing)

+ Users should be able to run your code regardless of OS. (Two+ separate versions are allowed for POSIX/UNIX compatibility). 
This disqualifies non-compliant languages. 
If a custom compiler works cross-OS, that is allowed so long as the binary/executable can be ran with standard available dependencies.
```

Benchmarking:
```diff
+ Include an output to console that details the time taken to generate strings AND write to the file.
If your language has a standard testing library (like go's _test), check your time manually inside your code.
- Rob informs me there is a time command. I am retarded. If you want to use that, you can go ahead and use that.

+ Send a copy of your code for us to test. 
(if you want to keep it secret, just compile. Won't bother to sift for your code unless you meet any of the special cases above.)
```

[linux download of our current generator](https://cdn.discordapp.com/attachments/810233494185574433/832002423300030484/main)

[windows download of our current generator](https://cdn.discordapp.com/attachments/810233494185574433/832011118020919296/main.exe)