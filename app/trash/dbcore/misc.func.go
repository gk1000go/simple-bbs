package dbcore

func xn_log(str ...string) {
	if(len(str) == 0) {
		return
	}
	/*
		if(DEBUG == 0 && strpos($file, 'error') === FALSE) return;
		$time = $_SERVER['time'];
		$ip = $_SERVER['ip'];
		$conf = _SERVER('conf');
		$uid = intval(G('uid')); // xiunophp 未定义 $uid
		$day = date('Ym', $time); // 按照月存放，否则 Ymd 目录太多。
		$mtime = date('Y-m-d H:i:s'); // 默认值为 time()
		$url = isset($_SERVER['REQUEST_URI']) ? $_SERVER['REQUEST_URI'] : '';
		$logpath = $conf['log_path'].$day;
		!is_dir($logpath) AND mkdir($logpath, 0777, true);

		$s = str_replace(array("\r\n", "\n", "\t"), ' ', $s);
		$s = "<?php exit;?>\t$mtime\t$ip\t$url\t$uid\t$s\r\n";

		@error_log($s, 3, $logpath."/$file.php");
	*/
}
