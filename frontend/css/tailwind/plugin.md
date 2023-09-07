要使用 Tailwind CSS 的插件，如 `@tailwindcss/forms`，你需要按照以下步骤进行：

### 1. 安装插件

首先，你需要使用 npm 或 yarn 安装所需的插件。对于 `@tailwindcss/forms` 插件，你可以使用以下命令来安装它：

```sh
npm install @tailwindcss/forms
```

### 2. 更新 Tailwind CSS 配置文件

接着，你需要在你的 `tailwind.config.js` 文件中引入并使用该插件。打开 `tailwind.config.js` 文件，并在 `plugins` 部分添加 `@tailwindcss/forms` 插件，如下所示：

```js
module.exports = {
  // 其他配置...

  plugins: [
    require('@tailwindcss/forms'),
    // 其他插件...
  ],
}
```

### 3. 重新生成 Tailwind CSS 样式

由于你已经更新了 Tailwind CSS 的配置，你需要重新生成 Tailwind CSS 的样式。运行以下命令来做到这一点：

```sh
npx tailwindcss build src/styles.css -o dist/styles.css
```

### 4. 使用插件提供的工具类

现在，`@tailwindcss/forms` 插件已经被正确安装和配置，你可以开始使用它提供的工具类和样式来设计你的表单。

例如，在你的 HTML 文件中创建一个表单，并使用 Tailwind CSS 的工具类来样式化它：

```html
<form class="space-y-4">
  <div>
    <label for="name" class="block text-gray-700">Name</label>
    <input id="name" name="name" class="block w-full mt-1 border-gray-300 rounded-md" type="text">
  </div>
  <div>
    <label for="email" class="block text-gray-700">Email</label>
    <input id="email" name="email" class="block w-full mt-1 border-gray-300 rounded-md" type="email">
  </div>
  <div>
    <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md">Submit</button>
  </div>
</form>
```

### 5. 查看结果

和之前一样，你可以使用一个简单的 HTTP 服务器来查看你的结果。如果你已经安装了 `live-server`，你可以运行以下命令来启动它：

```sh
live-server
```

现在，你应该可以在你的 web 浏览器中看到你的表单，并且它应该已经被正确地样式化了。

### 6. 查阅文档

为了更好地利用插件提供的功能和工具类，你应该查阅插件的文档来了解更多信息。你可以在 [Tailwind CSS Forms 插件文档](https://github.com/tailwindlabs/tailwindcss-forms) 中找到更多关于如何使用该插件的信息。

这样你就完成了 Tailwind CSS 插件的安装和使用！希望这可以帮到你！