
### Im-export Link
[Excel 读写库 spout](https://github.com/box/spout)  
[Laravel excel](https://github.com/Maatwebsite/Laravel-Excel)  
[Laravel excel zip](https://github.com/cblink/laravel-excel-zip)  

### 常见导入方案


### 常见导出方案
一般简单的导出，获取数据之后，按照指定格式即可完成处理。  
```php
/**
 * 导出数据
 *
 * @param array  $data      导出数据（格式为二维数组）
 * @param array  $header    导出头部标题栏
 * @param string $fileName  导出文件名（含后缀）
 * @param bool   $needIconv 是否需要中文转换（避免乱码）
 *
 */
function doExport($data = [], $header = [], $fileName = 'demo.xlsx', $needIconv = false)
{
    set_time_limit(0);
    ini_set('memory_limit', '-1');

    header('Content-type:application/octet-stream');
    header('Accept-Ranges:bytes');
    header('Content-type:application/vnd.ms-excel');
    header('Content-Disposition:attachment;filename=' . $fileName);
    header('Pragma: no-cache');
    header('Expires: 0');

    // Excel 头部标题构造
    if ($header && is_array($header)) {
        if ($needIconv) {
            foreach ($header as $key => $val) {
                $header[$key] = iconv('UTF-8', 'GBK', $val);
            }
        }

        $header = implode("\t", $header);
        echo "{$header}\n";
    }

    // Excel 数据构造
    if ($data) {
        $tmpData = [];
        foreach ($data as $key => $val) {
            if ($needIconv) {
                foreach ($val as $k => $v) {
                    $tmpData[$key][$k] = iconv('UTF-8', 'GBK', $v);
                }
            }
            $tmpData[$key] = implode("\t", array_values($val));
        }

        echo implode("\n", $tmpData);
    }
}

// example
$time = date('YmdHis');
$name = "demo-{$time}.xlsx";
$headers = [
    'Id',
    'Mobile',
    'Name'
];

$data = [
    ['id' => 1, 'mobile' => '13800000000', 'name' => '我是中文1'],
    ['id' => 2, 'mobile' => '13800000000', 'name' => '我是中文2'],
    ['id' => 3, 'mobile' => '13800000000', 'name' => '我是中文3'],
];

doExport($data, $headers, $name, true);
```

支持百万数据导出，也常常是开发中相对常见的问题。  
通常，这样的数据库大数据量查询非常耗时，可以考虑分页或根据 id 分片读取。  
```php

/**
 * 导出数据（支持打包）
 *
 * @param        $data     数据
 * @param array  $header   数据字段及头部标题配置
 * @param string $fileName 文件名称
 * @param int    $chunk    每个分片文件限制
 */
function doExport($data, $header = [], $fileName = 'demo', $chunk = 1000000)
{
    header('Content-Type: application/csv; charset=UTF-8');
    header('Content-Disposition: attachment; filename="' . $fileName . '.csv"');
    header('Cache-Control: max-age=0');

    // 数据字段
    $columns = array_keys($header);
    // 数据标题
    $titles  = array_values($header);
    // 计数器
    $counter = 0;
    // 刷新输出 buffer 限制行
    $bufferLimit = 10000;
    // 文件列表
    $fileList = [];
    // 文件指针
    $fp = null;

    foreach ($data as $item) {
        // 达到分配数量，关闭上一个文件并打开下一个文件
        if ($counter % $chunk === 0){
            $index = ($counter / $chunk) + 1;
            if ($index !== 1){
                fclose($fp);
            }

            // 导出文件列表到本地临时目录
            $fileList[] = "/tmp/{$fileName}-{$index}.csv";
            $fp = fopen(end($fileList), 'w');
            // utf8 字符集头部声明
            fputs($fp, chr(0xEF) . chr(0xBB) . chr(0xBF));
            // 头部标题构造
            fputcsv($fp, $titles);
        }

        ++$counter;
        // 释放 buffer
        if ($counter % $bufferLimit === 0) {
            ob_flush();
            flush();
        }

        // 获取对应头部标题的列数据并写入
        $rowData = [];
        foreach ($columns as $column) {
            $rowData[] = $item[$column] ?? null;
        }
        fputcsv($fp, $rowData);
        unset($rowData);
    }
    fclose($fp);

    // 多个文件时，zip 打包并删除临时文件
    if (count($fileList) > 1) {
        $zip = new \ZipArchive();
        $zipName = "/tmp/${fileName}.zip";
        $zip->open($zipName, \ZipArchive::CREATE);

        // 添加文件到压缩文件
        foreach ($fileList as $file) {
            $zip->addFile($file, str_replace("", '', basename($file)));
        }
        $zip->close();

        // 压缩完删除临时文件
        foreach ($fileList as $file) {
            @unlink($file);
        }

        // 输出压缩文件提供下载
        header("Cache-Control: max-age=0");
        header("Content-Description: File Transfer");
        header('Content-disposition: attachment; filename=' . $fileName . '.zip');
        header("Content-Type: application/zip");
        header("Content-Transfer-Encoding: binary");
        header('Content-Length: ' . filesize($zipName));
        @readfile($zipName);
        @unlink($zipName);
        exit();
    }

    // 只有一个文件时，直接输出下载并删除临时文件
    $fileName = head($fileList);
    @readfile($fileName);
    @unlink($fileName);
}

// example
$time = date('YmdHis');
$name = "demo-{$time}";
$headers = [
    'id' => 'Id',
    'phone' => 'Mobile',
    'name' => 'Name'
];

$data = [
    ['id' => 1, 'phone' => '13800000000', 'name' => '我是中文1'],
    ['id' => 2, 'phone' => '13800000000', 'name' => '我是中文2'],
    ['id' => 3, 'phone' => '13800000000', 'name' => '我是中文3'],
];

doExport($data, $headers, $name);
```

### 常见导入导出问题

