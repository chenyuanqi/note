
### Yii 常见问题
**composer install 报 bower-asset no matching**  
需要先执行 composer global require "fxp/composer-asset-plugin:^1.1.3"，再执行 composer install。  
```bash
rm -rf ~/.composer/vendor
rm ~/.composer/composer.lock
cd ~/.composer
composer clear-cache
composer self-update
composer require "fxp/composer-asset-plugin:^1.1.3"
# 或者 composer require yidas/yii2-bower-asset
composer install
```
或者在 composer.json 添加
```json
{
    "php": ">=5.4.0",
    "yidas/yii2-composer-bower-skip": "~2.0.0",
    "yiisoft/yii2": "~2.0.5",
    "yiisoft/yii2-bootstrap": "~2.0.0"
}
```
