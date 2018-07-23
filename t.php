<?php
echo "\nsystem";
$last_line = system('ls', $return_var);
echo "\nreturn_var:";
print_r($return_var);
echo "\nlast_line:";
print_r($last_line);

echo "\n\nexec";
exec('ls', $output, $return_var);
echo "\nreturn_var:";
print_r($return_var);
echo "\noutput:";
print_r($output);

echo "\n\nshell_exec";
$output = shell_exec('ls');
echo "\noutput:";
print_r($output);
?>
