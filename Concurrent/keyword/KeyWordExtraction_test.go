package keyword

import "testing"

/**
=== RUN   TestKeyWordSerial
591854
de:1444
italian:1431
best:1379
film:1372
silent:1358
documentary:1325
her:1212
festival:1209
malayalam:1188
she:1171
k:1161
john:1121
japanese:1054
external:1046
n:1046
la:1011
links:984
infobox:975
short:974
jean:951
name:924
p:907
s:904
caption:900
m:899
country:884
director:882
german:874
writer:873
image:867
by:859
george:841
producer:835
films:834
music:806
jack:801
starring:801
released:801
argentina:782
cinematography:780
the:766
kim:756
william:754
kong:752
he:750
kumar:735
distributor:733
runtime:724
of:705
tamil:698
references:689
i:688
le:682
rao:681
lee:674
robert:674
hong:670
editing:668
van:664
budget:664
david:662
v:654
charles:647
michael:646
harry:646
french:636
british:634
mary:612
tom:608
a:600
academy:594
r:591
james:588
foreign:581
peter:580
d:578
el:572
paul:564
richard:558
narrator:557
2014:555
war:552
frank:550
australia:532
language:529
in:509
australian:508
joe:503
king:498
w:494
gang:493
dr:489
khan:483
kannada:482
man:482
2013:480
is:480
award:478
b:476
von:469
Execution Time: 183306
Vocabulary Size: 591854
Keyword Size: 166321
Number of Documents: 100673
--- PASS: TestKeyWordSerial (183.39s)
PASS

*/
func TestKeyWordSerial(t *testing.T) {
	KeywordSerial()
}
