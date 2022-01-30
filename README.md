# CSVForAnki
Program to convert plain text to CSV file which can imported into Anki. The motivation of this program is to save time by automatically coverting Question and Answer into CSV file which can be imported directly into Anki.

How to use:-
makecsv <input file name>

This will create a directory called "import_ready" if not already present in the present working directory and save the CSV version of the input file.

The format of input file is:-
Question 1
Answer 1
Tag1

Question 2
Answer 2.a
Answer 2.b
Answer 2.b
Tag2

Below is one concrete example of input file:-

## Sample Input File

If the communication is not initiated but happens as a response is security rule needed in Network Security Group? Explain why.
No.

Not needed because NSG is stateful. It will remember that a request was initiated and it will remember till a response is received.
VNet

If you have 4 VMs which need to be applied the same network security rule how will you do it without adding 4 different IP addresses for the VMs?
By creating Application Security group and putting the VMs in the same Application Security Group. This application Security Group can then be given a Source or Destination to the rules of NSG
vNet

What are the points that you will check if traffic is not flowing between two resources in Azure?
0. Check if DNS Route is set.
1. Check if VNet Route is set.
2. Check if Azure Firewall is allowing the traffic.
3. Check if Subnet level NSG allows traffic.
4. Check if NIC level NSG allows traffic.
5. Check if OS firewall allows traffic.
networking
