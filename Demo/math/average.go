package math

func Average(intSliceParam []int)int  {
	total:=0
	count:=0
	for _,x:= range intSliceParam{
		total+=x
		count+=1
	}
	return total/count

}
