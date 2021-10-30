You were asked to build a simple encryption program. This program should make it possible to send coded messages without anyone being able to read them. The process is very simple. Three passes are made throughout the text.

In the first pass, only characters that are lowercase and uppercase letters must be shifted 3 positions to the right, according to the ASCII table: letter 'a' must become letter 'd', letter 'e' must become character '|' and so on. On the second pass, the line must be reversed. On the third and last pass, any and all characters from the halfway onwards (truncated) must be shifted one position to the left in the ASCII table. In this case, 'b' becomes 'a' and 'a' becomes '`'.

For example, if the input is “Text #3”, the first processing on this input would produce “Wh {wr #3”. The result of the second processing reverses the characters and produces “3 # rw {hW”. Finally, with the characters shifting from halfway onwards, the final result should be “3 # rvzgV”.

Input
The input contains several test cases. The first line of each test case contains an integer N (1 ≤ N ≤ 1 * 104), indicating the number of lines that the problem must handle. The N lines each contain M (1 ≤ M ≤ 1 * 103) characters.
Exit
For each entry, an encrypted message must be presented.

**Input**	
4
Texto #3
abcABC1
vxpdylY .ph
vv.xwfxo.fd

**Outpt**
3# rvzgV
1FECedc
ks. \n{frzx
gi.r{hyz-xx
