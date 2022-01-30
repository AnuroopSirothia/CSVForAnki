# CSVForAnki
Program to convert plain text to CSV file which can imported into Anki. The motivation of this program is to save time by automatically coverting Question and Answer into CSV file which can be imported directly into Anki.

## How to use:-
makecsv \<input file name\>

This will create a directory called "import_ready" if not already present in the present working directory and save the CSV version of the input file.

# Input File Format
The format of input file is:-<br>
Question 1<br>
Answer 1<br>
Tag1<br>

Question 2<br>
Answer 2.a<br>
Answer 2.b<br>
Answer 2.b<br>
Tag2

Below is one concrete example of input file:-

Sample Input File

\# Question 1 <br>
If the communication is not initiated but happens as a response is security rule needed in Network Security Group? Explain why.<br>
No.

Not needed because NSG is stateful. It will remember that a request was initiated and it will remember till a response is received.<br>
VNet

\# Question 2<br>
If you have 4 VMs which need to be applied the same network security rule how will you do it without adding 4 different IP addresses for the VMs?<br>
By creating Application Security group and putting the VMs in the same Application Security Group. This application Security Group can then be given a Source or Destination to the rules of NSG<br>
vNet

\#Question 3<br>
What are the points that you will check if traffic is not flowing between two resources in Azure?<br>
1. Check if DNS Route is set.
2. Check if VNet Route is set.
3. Check if Azure Firewall is allowing the traffic.
4. Check if Subnet level NSG allows traffic.
5. Check if NIC level NSG allows traffic.
6. Check if OS firewall allows traffic.<br>
networking
