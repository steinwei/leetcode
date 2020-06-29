// const fs = require('fs')
// const path=require('path')
// const filepath = path.join(__dirname,'../')
// const regEx = /.vuepress/;
// const regMd = /.README/
// const files = fs.readdirSync(filepath).map(item=>{
//     if(regEx.test(item)||regMd.test(item)){
//         return null
//     }
//     return `/${item}/`
// });
const menuConfig = require("vuepress-bar")();

const autobar_options = {
  rootDir: 'xxx',
  stripNumbers : true,
  maxLevel : 2,
  navPrefix : "nav",
  skipEmptySidebar : true,
  skipEmptyNavbar : true,
  multipleSideBar : true,
  setHomepage : 'top'
};

module.exports={
    base: '/leetcode/',
    dest: 'docs',
    plugins: ['autobar', autobar_options],
    themeConfig: {
      ...menuConfig,
        nav: [
            { text: 'Home', link: '/' }
          ],
          // sidebar:files,
        displayAllHeaders: false // 默认值：false
      }
}