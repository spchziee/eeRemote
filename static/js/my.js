function isPhone(){
    return /Android|webOS|iPhone|iPod|BlackBerry/i.test(navigator.userAgent);
}


function Yesterday()
{
	todaysdate = new Date();
	newdate = new Date(todaysdate.getTime() - 1000*60*60*24);
	date  = newdate.getDate();
	day  = newdate.getDay() + 1;
	month = newdate.getMonth() + 1;
	yy = newdate.getYear();
	year = (yy < 1000) ? yy + 1900 : yy;
	return year + "-" + month + "-" + date;
}
 
function Today()
{
	todaysdate = new Date();
	date  = todaysdate.getDate();
	day  = todaysdate.getDay() + 1;
	month = todaysdate.getMonth() + 1;
	yy = todaysdate.getYear();
	year = (yy < 1000) ? yy + 1900 : yy;
	return year + "-" + month + "-" + date;
}
