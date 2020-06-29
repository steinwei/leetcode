console.log(`当前工作目录是: ${process.cwd()}`);

module.exports={
    base: '/leetcode/',
    dest: 'docs',
    plugins: ['autobar'],
    themeConfig: {
        nav: [
            { text: 'Home', link: '/' }
          ],
        displayAllHeaders: false // 默认值：false
      }
}