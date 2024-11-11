<template>
  <main class="main">
    <!--  主页显示组件-->
    <div class="home">
      <!--    上面的取色按钮和状态展示 -->
      <div class="head">
        <div class="cp">
          <img :class="{ img: cpData.state == 1 }" @click="pickColor" src="@/assets/img/xiguan.png" >
        </div>
        <div class="texts">
          <div class="text1" :style="{'color':cpData.tip[cpData.state].color}">{{cpData.tip[cpData.state].text}}</div>
          <div class="text2">提取颜色值:{{cpData.colors[0]}}</div>
        </div>
      </div>

      <!--    颜色显示栏-->
      <div class="colors">
          <div v-for="(color,index) in cpData.colors" class="tc" @click="color_vary(index,'old')">
            <div class="color" :style="{'background-color':color}"></div>
            <div class="key">{{showColorKey(index)}}</div>
          </div>
      </div>

      <!--    颜色值-->
      <div class="cv">
        <div class="rgb"><div class="name">RGB:</div><div class="value">{{cpData.rgb}}</div><div class="button" @click="copyText(cpData.rgb)">复制</div></div>
        <div class="rgb"><div class="name">HTML:</div><div class="value">{{cpData.colors[0]}}</div><div class="button" @click="copyText(cpData.colors[0])">复制</div></div>
        <div class="rgb"><div class="name">HEX:</div><div class="value">{{cpData.hex}}</div><div class="button" @click="copyText(cpData.hex)">复制</div></div>
      </div>

      <!--    开源图标-->
      <div class="git">
        <img @click="open_url('www.baidu.com')" src="@/assets/img/gitee.png" />
        <img  @click="open_url('www.daodaoyouli.com')" src="@/assets/img/github.png" />
      </div>
    </div>
  </main>
</template>


<script setup lang="ts">
import {reactive} from 'vue'
import {OpenUrl} from '../../wailsjs/go/main/App'
// import HomeCom from "@/components/HomeCom.vue"
// import Index from "@/router";

const cpData = reactive({
  state:0, //1取色中 2成功  3失败
  tip:[{text:"点击吸管工具开始取色",color:"#01AAED"},{text:"取色中...",color: "#FFB800"},{text: "取色成功",color: "#5FB878"},{text: "取色失败请重试",color: "#FF5722"}],
  colors:["","#ffffff","#ffffff","#ffffff","#ffffff","#ffffff","#ffffff"],
  rgb:"",
  hex:""
})



// 颜色统一批量处理
const color_vary = (key:number,type:string) =>{
  if (!/^#[0-9A-Fa-f]{6}$/.test(cpData.colors[key])) {
    TipMsg('没有颜色','error');
    return
  }
  cpData.rgb = hexToRgb(cpData.colors[key]).join(",");
  cpData.hex = hexToUpperCaseWithoutHash(cpData.colors[key]);
  if(type != "new"){
    cpData.colors[0] = cpData.colors[key];
  }
}


// 显示颜色的下标
const showColorKey = (key:number) => {
  switch (key) {
    case 0:
      return "当前";
    case 1:
      return "上次";
    default:
      return key.toString(10)
  }
}

declare class EyeDropper {
  open(): Promise<{ sRGBHex: string }>;
  // 其他可能的方法或属性...
}
const pickColor = async () => {
  cpData.state = 1;
  try {
    const eyeDropper = new EyeDropper();
    const result = await eyeDropper.open();
    // 让颜色递归
    for (let i = 6; i > 0; i--) {
      cpData.colors[i] = cpData.colors[i-1]
    }
    cpData.colors[0] = result.sRGBHex;
    color_vary(0,'new');
    cpData.state = 2;
  } catch (error) {
    cpData.state = 3;
    TipMsg('取色失败','error');
  }
};


// 用浏览器打开外部地址
const open_url = (url:string)=>{
  OpenUrl(url).then(result => {
    console.log(result);
  })
}

// 将16进制颜色转为RGB
const hexToRgb = (hex:string)=>{
  // 确保输入以#开头，如果不是，则添加#
  hex = hex.replace(/^#?([a-f\d]{6})$/i, '#$1');

  // 使用正则表达式解析RGB值
  // 解析后的结果将是一个包含三个元素的数组，每个元素都是0到255之间的十进制数
  const bigint = parseInt(hex.slice(1), 16);
  const r = (bigint >> 16) & 255; // 右移16位，然后与255进行位与操作获取红色值
  const g = (bigint >> 8) & 255;  // 右移8位，然后与255进行位与操作获取绿色值
  const b = bigint & 255;         // 直接与255进行位与操作获取蓝色值

  // 返回RGB值数组
  return [r,g,b];
}

// 将16进制转为大写
const hexToUpperCaseWithoutHash = (hex:string)=> {
  // 去掉#前缀，并将字符串转换为大写
  return hex.slice(1).toUpperCase();
}


// 复制文本到剪切板的函数
const copyText = async (value:string) => {
  try {
    await navigator.clipboard.writeText(value);
    TipMsg('复制成功','success');
  } catch (err) {
    TipMsg('复制失败','error');
  }
};

// 弹出对应的消息
const TipMsg = (msg:string,type:string) => {
  // ElMessage({message: msg, type: type})
  if(type == 'success'){
    ElMessage.success(msg);
  }else{
    ElMessage.error(msg);
  }
}

</script>


<style scoped lang="scss">
.main{
  margin: 0;
  padding: 10px;
  width:260px;
  height: 300px;
  background: #FFFFFF;
  .home{
    .head{
      display: flex;
      align-items:center;
      justify-content: flex-start;
      .cp{
        width: 63px;
        height: 60px;
        background: #FFFFFF;
        border-radius: 13px 13px 13px 13px;
        border: 1px solid #2F1A75;
        display: flex;
        align-items: center;
        justify-content: center;
        cursor: pointer;
        .img{
          transform: rotate(-43deg);
        }
      }
      .texts{
        padding-left: 10px;
        .text1{
          font-weight: 400;
          font-size: 18px;
          color: #01AAED;
        }
        .text2{
          color: #707070;
          font-weight: 400;
          font-size: 16px;
        }
      }
    }
    .colors{
      margin-top: 8px;
      height: 60px;
      display: flex;
      align-items:center;
      justify-content: flex-start;
      gap: 5px;
      .tc{
        cursor: pointer;
        .color{
          width: 30px;
          height: 30px;
          border-radius: 9px;
          border: 1px solid #CFCFCF;
          background-color: #ffffff;
        }
        .key{
          text-align: center;
          font-size: 12px;
          color: #707070;
          font-weight: 200;
        }
      }
    }
    .cv{
      margin-top: 5px;
      .rgb{
        padding-top: 5px;
        display: flex;
        align-items:center;
        justify-content: flex-start;
        gap: 5px;
        .name{
          width: 50px;
          height: 22px;
          font-weight: 400;
          font-size: 16px;
          text-align: left;
        }
        .value{
          padding-left: 4px;
          width: 150px;
          height: 26px;
          line-height: 26px;
          background: #FFFFFF;
          border-radius: 5px 5px 5px 5px;
          border: 1px solid #DEDEDE;
        }
        .button{
          cursor: pointer;
          width: 30px;
          height: 16px;
          font-weight: 400;
          font-size: 14px;
          color: #01AAED;
          text-align: center;
          font-style: normal;
          text-transform: none;
        }

      }
    }
    .git{
      margin: 8px 10px 0 10px;
      border-top: 1px solid #EBEBEB;
      display: flex;
      align-items:center;
      justify-content: flex-end;
      gap: 5px;
      img{
        margin-top: 8px;
        cursor: pointer;
      }
    }
  }
}
</style>
