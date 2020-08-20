<?php
    $array = file('test.txt');
    foreach($array as &$arr){
	$element = explode(";", $arr);
	if(isset($element[0], $element[1]))2
	    echo "INSERT INTO logs (flowtime, watertype) VALUES (FROM_UNIXTIME(".$element[0]. "),'".trim($element[1])."');\n";
    }
    echo mktime(0,0,0, date('m'), 1, date('Y');
    echo mktime(0,0,0, date('m'), date('d'), date('Y');
?>