#!/bin/bash
array[0]="Chau"
array[1]="Rosa"
array[2]="Tracy"
array[3]="Clifford"
array[4]="Clinton"
array[5]="Jarod"
array[6]="Cary"
array[7]="Thad"
array[8]="Earnestine"
array[9]="Paige"
array[10]="Lindsey"
array[11]="Nannie"
array[12]="Lyman"
array[13]="Olivia"
array[14]="Sue"
array[15]="Eula"
array[16]="Louisa"
array[17]="Lavonne"
array[18]="Tracey"
array[19]="Loraine"
array[20]="Shawna"
array[21]="Rosie"
array[22]="Val"
array[23]="Joesph"
array[24]="Velma"
array[25]="Kelsey"
array[26]="Alden"
array[27]="Efrain"
array[28]="Sylvia"
array[29]="Jasper"

touch big_data.csv
echo "first_name,last_name,dob" > big_data.csv

size=${#array[@]}
for i in $(seq 1 $1)
do
	fn_idx=$(($RANDOM % $size))
	ln_idx=$(($RANDOM % $size))
	year=`printf %04d $((($RANDOM % 72) + 1950))`
	month=`printf %02d $((($RANDOM % 12) + 1))`
	day=`printf %02d $((($RANDOM % 28) + 1))`
	echo "${array[$fn_idx]},${array[$ln_idx]},$year$month$day" >> big_data.csv
done
