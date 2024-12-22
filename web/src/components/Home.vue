<template>
  <div style="">
    <!--    头部  -->
    <!--    第一行信息-->
    <div class="header">
      <img
          src="data:image/png;base64,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"
          alt=" " >
    </div>
    <van-divider content-position="left" style="height:2px;margin:5px"
                 :style="{ color: '#1989fa', borderColor: '#1989fa', padding: '0 10px' }"></van-divider>

    <div class="btns">
      <!--        品牌下拉列表-->
      <el-button  class="commonElButton" @click="drawerButton = true"
                 > 品牌
      </el-button>
      <!--   品牌下拉抽屉组件     -->
      <el-drawer
          v-model="drawerButton"
          size="85%"
          direction="btt"
          :before-close="handleClose"
      >
        <div>
          <el-card style="
              width: 700px;
              display: flex;
              align-items: center;
              margin-bottom: 10px;
              padding: 0;
            "
             shadow="hover"
             v-for="brand in brandList"
             @click="drawerSeriesButton = true; getCarSeries(brand.id)"
          >
            <div style="display: flex; align-items: center;"> <!-- 保持这个div来管理图片和文本的对齐 -->
              <img :src="brand.imageUrl" alt="" style="width: 50px; margin-right: 20px;"/>
              <span style="font-size: 40px; text-align: center;">{{ brand.brand.split("-")[1] }}</span> <!-- 文本居中 -->
            </div>
          </el-card>
          <el-drawer
              v-model="drawerSeriesButton"
              size="60%"
              direction="rtl"
              :before-close="handleClose"

          >
            <div class="flex flex-wrap gap-4">
              <el-card
                  style="width: 400px;height:50px; margin-bottom: 6px;
                                display: flex; flex-direction: column; justify-content: center; align-items: center;"
                  @click="tipToHomePage(series.Id)"
                  shadow="hover"
                  v-for="series in seriesList"
              >
                <!--                            p标签居中-->

                <p style="text-align: left;">{{ series.Series }}</p>
              </el-card>
            </div>

          </el-drawer>
        </div>
      </el-drawer>

      <!--        搜索框-->
      <el-input v-model="keyword" type="text"

                placeholder="          Please input"></el-input>
      <el-button  class="commonElButton"@click="search">搜索
      </el-button>

      <!--        -->


<!--      <el-select class="custom-select"-->
<!--                 v-model="modelSearch"-->
<!--                 placeholder="Select"-->
<!--                 size="large"-->
<!--                 @change="getCarData"-->

<!--      >-->
<!--        <el-option-->
<!--            v-for="model in models"-->
<!--            :label="model.model"-->
<!--            :value="model.id">-->
<!--        </el-option>-->
<!--      </el-select>-->

      <FilterBar></FilterBar>

      <el-button
          class="four-button-group"
          type:plain
          v-for="(saleType,index) in saleType"
          color="#D5E8FF"
          :style="{color:saleType.id==saleTypeInt?'red':'black'}"
          @click="chooseCarsSaleType(saleType.id),saleTypeInt=saleType.id"
      >{{ saleType.saleType }}
      </el-button>
    </div>

    <!--    中部-->
    <ul >
      <li v-for="item in carData"
          @click="tipToOtherPage(item.ind)">
        <img :src="item.cover" alt="" >

        <span class="text-truncate">
          <span
            class = "saleTypePatten"
            :style="{backgroundColor: saleTypeColor(item.saleType)}"
        >{{ saleType[item.saleType].saleType }}</span>   {{ item.carName }}</span>
        <p class="dateAndMileage">{{ formattedDate(item.date_of_registration) }} / {{ item.mileage }}万公里</p>
        <h4>{{ item.CurrentPrice }}万</h4>
      </li>
    </ul>

    <!--    底部  -->
    <div class="empty"></div>
    <ul class="tabbar" >
      <li
          v-for="(item,index) in tabbar"
          :style="{color:index==tabbarIndex?'#2196f3':'black'}"
          @click="tabbarIndex=index,TipToChangPage(index)"
      >
        <p class="iconfont icon" v-html="item.icon"></p>
        <p class="title">{{ item.title }}</p>
      </li>
    </ul>
  </div>
</template>
<script>
import axios from "axios";
import FilterBar from "./FilterBar.vue";

export default {
  components: {FilterBar},
  data() {
    return {
      keyword: "", //输入框绑定数据
      carData: [],
      originalCarData: [],
      newCarData: [],
      models: [
      {
        "id": 1,
        "model": "两厢轿车",
        "abbreviation": "两厢轿车",
        "logo": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/131329075347193856"
      },
      {
        "id": 2,
        "model": "三厢轿车",
        "abbreviation": "三厢轿车",
        "logo": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/131329075875676160"
      },
      {
        "id": 3,
        "model": "SUV",
        "abbreviation": "SUV",
        "logo": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/131329075980533760"
      },
      {
        "id": 4,
        "model": "跑车",
        "abbreviation": "跑车",
        "logo": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/131329075485605889"
      },
      {
        "id": 5,
        "model": "MPV",
        "abbreviation": "MPV",
        "logo": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/131329075913424896"
      },
      {
        "id": 6,
        "model": "面包车",
        "abbreviation": "面包车",
        "logo": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/131329168909533184"
      },
      {
        "id": 7,
        "model": "皮卡",
        "abbreviation": "皮卡",
        "logo": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/131329026290614272"
      }
    ],
      modelSearch: "",
      tabbar:[
        {
          title: "首页",
          icon: "&#xe61d;",
        },
        {
          title: "二手车",
          icon: "&#xfd72;",
        },
        {
          title: "新车",
          icon: "&#xe610;",
        },
        {
          title: "卖车",
          icon: "&#xe6e9;",
        },
      ],
      tabbarIndex: 0,
      saleType: [{
        id: 0,
        saleType: "腾发自营",
      }, {
        id: 1,
        saleType: "车主直营",
      }, {
        id: 2,
        saleType: "限时特惠",
      }, {
        id: 3,
        saleType: "严选好车",
      }
      ],
      saleTypeInt: 100,
      drawerButton: false,
      drawerSeriesButton: false,
      brandList:  [
      {
        "id": 1,
        "brand": "A - 奥迪",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539288541462528.png"
      },
      {
        "id": 2,
        "brand": "A - 阿斯顿·马丁",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539179229511680.png"
      },
      {
        "id": 3,
        "brand": "A - 阿尔法·罗密欧",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539236502732800.png"
      },
      {
        "id": 4,
        "brand": "A - AC SCHNITZER",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266504955152564224.png"
      },
      {
        "id": 536,
        "brand": "A - ALPINA",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266596173756760064.png"
      },
      {
        "id": 1407,
        "brand": "A - 爱驰",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538839616716800.png"
      },
      {
        "id": 1714,
        "brand": "A - AUXUN傲旋",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539032911216640.png"
      },
      {
        "id": 1721,
        "brand": "A - ARCFOX极狐",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140539125664055296.png"
      },
      {
        "id": 2440,
        "brand": "A - AITO",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266493661376675840.png"
      },
      {
        "id": 5,
        "brand": "B - 本田",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537294732918784.png"
      },
      {
        "id": 6,
        "brand": "B - 别克",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537526032007168.png"
      },
      {
        "id": 7,
        "brand": "B - 宝马",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537473875836928.png"
      },
      {
        "id": 8,
        "brand": "B - 比亚迪",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537242106986496.png"
      },
      {
        "id": 9,
        "brand": "B - 奔驰",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537423363833856.png"
      },
      {
        "id": 10,
        "brand": "B - 标致",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537043552829440.png"
      },
      {
        "id": 11,
        "brand": "B - 保时捷",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537585096196096.png"
      },
      {
        "id": 12,
        "brand": "B - 奔腾",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538186962042880.png"
      },
      {
        "id": 13,
        "brand": "B - 北京汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538783886999552.png"
      },
      {
        "id": 14,
        "brand": "B - 北汽制造",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266592104434630656.png"
      },
      {
        "id": 15,
        "brand": "B - 宝骏",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537709771882496.png"
      },
      {
        "id": 16,
        "brand": "B - 宾利",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537372545646592.png"
      },
      {
        "id": 17,
        "brand": "B - 北汽威旺",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537645242515456.png"
      },
      {
        "id": 18,
        "brand": "B - 布加迪",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537120409255936.png"
      },
      {
        "id": 19,
        "brand": "B - 巴博斯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537186175942656.png"
      },
      {
        "id": 20,
        "brand": "B - 保斐利",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266593922334064640.png"
      },
      {
        "id": 115,
        "brand": "B - BEIJING汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538727825932288.png"
      },
      {
        "id": 156,
        "brand": "B - 北汽幻速",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538674138841088.png"
      },
      {
        "id": 167,
        "brand": "B - 北汽新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140538621508714496.png"
      },
      {
        "id": 172,
        "brand": "B - 宝沃",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140537771625283584.png"
      },
      {
        "id": 499,
        "brand": "B - 比速",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266596467148324864.png"
      },
      {
        "id": 573,
        "brand": "B - 北汽道达",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602857183051776.png"
      },
      {
        "id": 654,
        "brand": "B - 铂驰",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140536818243207168.png"
      },
      {
        "id": 853,
        "brand": "B - 比德文汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266519240440283136.png"
      },
      {
        "id": 2398,
        "brand": "B - 北汽瑞翔",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266495176497692672.png"
      },
      {
        "id": 2432,
        "brand": "B - 百智新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266494408675819520.png"
      },
      {
        "id": 21,
        "brand": "C - 长安",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535818543431680.png"
      },
      {
        "id": 22,
        "brand": "C - 长城",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140536152246452224.png"
      },
      {
        "id": 23,
        "brand": "C - 长安欧尚",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535934515937280.png"
      },
      {
        "id": 24,
        "brand": "C - 昌河",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535874256371712.png"
      },
      {
        "id": 497,
        "brand": "C - 成功",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140536078917435392.png"
      },
      {
        "id": 1716,
        "brand": "C - 车驰汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266512082277498880.png"
      },
      {
        "id": 1768,
        "brand": "C - 创维汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266507189466693632.png"
      },
      {
        "id": 25,
        "brand": "D - 大众",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535757684080640.png"
      },
      {
        "id": 26,
        "brand": "D - 东风风行",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535304288206848.png"
      },
      {
        "id": 27,
        "brand": "D - 东南",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535585084276736.png"
      },
      {
        "id": 28,
        "brand": "D - 东风小康",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535304288206848.png"
      },
      {
        "id": 29,
        "brand": "D - 道奇",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535648351158272.png"
      },
      {
        "id": 30,
        "brand": "D - 东风风神",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
      },
      {
        "id": 31,
        "brand": "D - DS",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535693653835776.png"
      },
      {
        "id": 32,
        "brand": "D - 东风风度",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
      },
      {
        "id": 33,
        "brand": "D - 东风",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
      },
      {
        "id": 170,
        "brand": "D - 东风风光",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535359460081664.png"
      },
      {
        "id": 574,
        "brand": "D - 电咖",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602661841731584.png"
      },
      {
        "id": 662,
        "brand": "D - 达契亚",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266526104389419008.png"
      },
      {
        "id": 1705,
        "brand": "D - 大运",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266514695194673152.png"
      },
      {
        "id": 2399,
        "brand": "D - 电动屋",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266494826730487808.png"
      },
      {
        "id": 35,
        "brand": "F - 福特",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534585615187968.png"
      },
      {
        "id": 36,
        "brand": "F - 丰田",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534705870077952.png"
      },
      {
        "id": 37,
        "brand": "F - 菲亚特",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534529088552960.png"
      },
      {
        "id": 38,
        "brand": "F - 法拉利",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534642573836288.png"
      },
      {
        "id": 39,
        "brand": "F - 福田",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534330265960448.png"
      },
      {
        "id": 40,
        "brand": "F - 福迪",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266594365344841728.png"
      },
      {
        "id": 162,
        "brand": "F - 福汽启腾",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534388952662016.png"
      },
      {
        "id": 543,
        "brand": "F - 飞碟",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266595905375830016.png"
      },
      {
        "id": 545,
        "brand": "F - 福田乘用车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534330265960448.png"
      },
      {
        "id": 1704,
        "brand": "F - 枫叶汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534450353078272.png"
      },
      {
        "id": 2314,
        "brand": "F - 飞凡汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/250669902166228992.jpg"
      },
      {
        "id": 44,
        "brand": "G - 广汽传祺",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534208022970368.png"
      },
      {
        "id": 45,
        "brand": "G - 广汽吉奥",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533814395928576.png"
      },
      {
        "id": 46,
        "brand": "G - 观致",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534015416336384.png"
      },
      {
        "id": 47,
        "brand": "G - GMC",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534262393733120.png"
      },
      {
        "id": 48,
        "brand": "G - 光冈",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533893991235584.png"
      },
      {
        "id": 636,
        "brand": "G - 广汽新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534151559249920.png"
      },
      {
        "id": 1075,
        "brand": "G - 国机智骏",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266516081860935680.png"
      },
      {
        "id": 1772,
        "brand": "G - 高合汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533288434401280.png"
      },
      {
        "id": 50,
        "brand": "H - 哈弗",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533232910204928.png"
      },
      {
        "id": 51,
        "brand": "H - 海马",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533018983923712.png"
      },
      {
        "id": 52,
        "brand": "H - 华泰",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266595357352263680.png"
      },
      {
        "id": 53,
        "brand": "H - 红旗",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533074604589056.png"
      },
      {
        "id": 54,
        "brand": "H - 海马商用车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533018983923712.png"
      },
      {
        "id": 55,
        "brand": "H - 黄海",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266588422129319936.png"
      },
      {
        "id": 56,
        "brand": "H - 哈飞",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266588241992351744.png"
      },
      {
        "id": 57,
        "brand": "H - 海格",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266588014572994560.png"
      },
      {
        "id": 58,
        "brand": "H - 恒天",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266587594270179328.png"
      },
      {
        "id": 59,
        "brand": "H - 汇众",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266587840559710208.png"
      },
      {
        "id": 145,
        "brand": "H - 悍马",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140533175796367360.png"
      },
      {
        "id": 146,
        "brand": "H - 华普",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266601140680916992.png"
      },
      {
        "id": 160,
        "brand": "H - 华颂",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266600187642773504.png"
      },
      {
        "id": 173,
        "brand": "H - 华泰新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266595357352263680.png"
      },
      {
        "id": 495,
        "brand": "H - 汉腾",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532961370963968.png"
      },
      {
        "id": 560,
        "brand": "H - 华骐",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602184026619904.png"
      },
      {
        "id": 668,
        "brand": "H - 霍顿",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266525668202774528.png"
      },
      {
        "id": 1703,
        "brand": "H - HYCAN合创",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140534079123619840.png"
      },
      {
        "id": 2094,
        "brand": "H - 航天凌河",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/136874226136121344.png"
      },
      {
        "id": 2279,
        "brand": "H - 黑豹",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/136874226136121344.png"
      },
      {
        "id": 2363,
        "brand": "H - 宏瑞汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266502399974178816.png"
      },
      {
        "id": 2385,
        "brand": "H - 华晨新日",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266497366624829440.png"
      },
      {
        "id": 2386,
        "brand": "H - 恒润汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266495888640180224.png"
      },
      {
        "id": 1766,
        "brand": "I - IMSA英飒",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266507880415363072.png"
      },
      {
        "id": 60,
        "brand": "J - 江淮",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532363326128128.png"
      },
      {
        "id": 61,
        "brand": "J - JEEP",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532904693334016.png"
      },
      {
        "id": 63,
        "brand": "J - 吉利帝豪",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532787043106816.png"
      },
      {
        "id": 64,
        "brand": "J - 捷豹",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532568238850048.png"
      },
      {
        "id": 66,
        "brand": "J - 江铃",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266586915904421888.png"
      },
      {
        "id": 67,
        "brand": "J - 金杯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266586807695572992.png"
      },
      {
        "id": 143,
        "brand": "J - 吉利",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532787043106816.png"
      },
      {
        "id": 542,
        "brand": "J - 江铃集团新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266596036464607232.png"
      },
      {
        "id": 572,
        "brand": "J - 君马汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532236255494144.png"
      },
      {
        "id": 634,
        "brand": "J - 捷途汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532446297849856.png"
      },
      {
        "id": 657,
        "brand": "J - 金冠",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266526916255678464.png"
      },
      {
        "id": 817,
        "brand": "J - 巨威",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266522276734697472.png"
      },
      {
        "id": 825,
        "brand": "J - 几何汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532853451522048.png"
      },
      {
        "id": 852,
        "brand": "J - 捷达",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532505315901440.png"
      },
      {
        "id": 2366,
        "brand": "J - 极氪",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266500983121510400.png"
      },
      {
        "id": 2367,
        "brand": "J - 捷尼赛思",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266500652593577984.png"
      },
      {
        "id": 73,
        "brand": "K - 凯迪拉克",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532166445498368.png"
      },
      {
        "id": 74,
        "brand": "K - 克莱斯勒",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532302533885952.png"
      },
      {
        "id": 75,
        "brand": "K - 开瑞",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530836997931008.png"
      },
      {
        "id": 76,
        "brand": "K - 科尼赛克",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532099017867264.png"
      },
      {
        "id": 77,
        "brand": "K - 卡尔森",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530899086213120.png"
      },
      {
        "id": 157,
        "brand": "K - 凯翼",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140531384727896064.png"
      },
      {
        "id": 158,
        "brand": "K - 卡威",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266600587368333312.png"
      },
      {
        "id": 552,
        "brand": "K - KTM",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140532006793510912.png"
      },
      {
        "id": 562,
        "brand": "K - 卡升",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602053579571200.png"
      },
      {
        "id": 678,
        "brand": "K - 凯佰赫",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266525227746328576.jpg"
      },
      {
        "id": 818,
        "brand": "K - 凯伦宾威",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266521276535799808.png"
      },
      {
        "id": 1720,
        "brand": "K - KARMA",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266509471105155072.png"
      },
      {
        "id": 2662,
        "brand": "K - 克蒂汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266493081543507968.png"
      },
      {
        "id": 78,
        "brand": "L - 铃木",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530238412029952.png"
      },
      {
        "id": 79,
        "brand": "L - 路虎",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530649365741568.png"
      },
      {
        "id": 80,
        "brand": "L - 雷克萨斯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530582684696576.png"
      },
      {
        "id": 82,
        "brand": "L - 兰博基尼",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530771281575936.png"
      },
      {
        "id": 83,
        "brand": "L - 陆风",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530030554906624.png"
      },
      {
        "id": 84,
        "brand": "L - 雷诺",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530098531991552.png"
      },
      {
        "id": 85,
        "brand": "L - 猎豹",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140529969456480256.png"
      },
      {
        "id": 86,
        "brand": "L - 劳斯莱斯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530709356871680.png"
      },
      {
        "id": 87,
        "brand": "L - 林肯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530517081587712.png"
      },
      {
        "id": 88,
        "brand": "L - 莲花",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266585778073632768.png"
      },
      {
        "id": 89,
        "brand": "L - 理念",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266585367342219264.png"
      },
      {
        "id": 90,
        "brand": "L - 路特斯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530389268561920.png"
      },
      {
        "id": 561,
        "brand": "L - 领克",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530160498638848.png"
      },
      {
        "id": 587,
        "brand": "L - LITE",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602489506168832.png"
      },
      {
        "id": 588,
        "brand": "L - 拉达LADA",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266602361424707584.png"
      },
      {
        "id": 619,
        "brand": "L - 陆地方舟",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266532463126249472.png"
      },
      {
        "id": 640,
        "brand": "L - 雷丁",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266530970423787520.png"
      },
      {
        "id": 661,
        "brand": "L - LORINSER",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266526371650469888.png"
      },
      {
        "id": 712,
        "brand": "L - 罗夫哈特",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266590257170874368.png"
      },
      {
        "id": 750,
        "brand": "L - 零跑汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/187582929180819456.png"
      },
      {
        "id": 815,
        "brand": "L - 理想",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140530451981795328.png"
      },
      {
        "id": 1718,
        "brand": "L - 凌宝汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266509984387301376.png"
      },
      {
        "id": 2365,
        "brand": "L - 岚图汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266501467030945792.png"
      },
      {
        "id": 2402,
        "brand": "L - LEVC",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/187582434122924032.png"
      },
      {
        "id": 2672,
        "brand": "L - LUMMA",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266492625048043520.png"
      },
      {
        "id": 92,
        "brand": "M - 马自达",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140529458212765696.png"
      },
      {
        "id": 93,
        "brand": "M - 名爵",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140529376398671872.png"
      },
      {
        "id": 94,
        "brand": "M - MINI",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140529650580324352.png"
      },
      {
        "id": 95,
        "brand": "M - 迈巴赫",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266533844620935168.png"
      },
      {
        "id": 96,
        "brand": "M - 玛莎拉蒂",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140529591553884160.png"
      },
      {
        "id": 97,
        "brand": "M - 迈凯伦",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140529532955262976.png"
      },
      {
        "id": 820,
        "brand": "M - 迈莎锐",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266520391885783040.png"
      },
      {
        "id": 2384,
        "brand": "M - 摩登汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266497803193155584.png"
      },
      {
        "id": 100,
        "brand": "N - 纳智捷",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140529221293309952.png"
      },
      {
        "id": 713,
        "brand": "N - 哪吒汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140528653564903424.png"
      },
      {
        "id": 101,
        "brand": "O - 讴歌",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140528348672557056.png"
      },
      {
        "id": 102,
        "brand": "O - 欧宝",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266584556096716800.png"
      },
      {
        "id": 103,
        "brand": "O - 欧朗",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140524701364518912.png"
      },
      {
        "id": 635,
        "brand": "O - 欧拉",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140528247157817344.png"
      },
      {
        "id": 637,
        "brand": "O - 欧尚汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140535934515937280.png"
      },
      {
        "id": 503,
        "brand": "P - 帕加尼",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527983642279936.png"
      },
      {
        "id": 679,
        "brand": "P - PGO",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266524767702482944.png"
      },
      {
        "id": 716,
        "brand": "P - POLESTAR",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527921231036416.png"
      },
      {
        "id": 104,
        "brand": "Q - 起亚",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527522449195008.png"
      },
      {
        "id": 105,
        "brand": "Q - 奇瑞",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527593261629440.png"
      },
      {
        "id": 106,
        "brand": "Q - 启辰",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527656402681856.png"
      },
      {
        "id": 107,
        "brand": "Q - 庆铃",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266583774664327168.png"
      },
      {
        "id": 624,
        "brand": "Q - 乔治·巴顿",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527808920158208.png"
      },
      {
        "id": 632,
        "brand": "Q - 前途汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527720843968512.png"
      },
      {
        "id": 2362,
        "brand": "Q - 奇瑞途居",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527593261629440.png"
      },
      {
        "id": 108,
        "brand": "R - 日产",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527376533553152.png"
      },
      {
        "id": 109,
        "brand": "R - 荣威",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527311500869632.png"
      },
      {
        "id": 110,
        "brand": "R - 瑞麒",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266583502063927296.png"
      },
      {
        "id": 570,
        "brand": "R - 如虎",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/136874226136121344.png"
      },
      {
        "id": 631,
        "brand": "R - 瑞驰新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266532164571496448.png"
      },
      {
        "id": 2360,
        "brand": "R - 瑞弗房车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266504248076795904.png"
      },
      {
        "id": 34,
        "brand": "S - 上汽MAXUS",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527009657782272.png"
      },
      {
        "id": 111,
        "brand": "S - 三菱",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527073155350528.png"
      },
      {
        "id": 112,
        "brand": "S - 斯柯达",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526860793544704.png"
      },
      {
        "id": 113,
        "brand": "S - 斯巴鲁",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526919245365248.png"
      },
      {
        "id": 114,
        "brand": "S - 双龙",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266583228406562816.png"
      },
      {
        "id": 116,
        "brand": "S - SMART",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140527129711345664.png"
      },
      {
        "id": 117,
        "brand": "S - 双环",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266582588179611648.png"
      },
      {
        "id": 118,
        "brand": "S - 世爵",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266582251150508032.png"
      },
      {
        "id": 169,
        "brand": "S - 思铭",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266599231412764672.png"
      },
      {
        "id": 498,
        "brand": "S - 斯威",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266596781352026112.png"
      },
      {
        "id": 644,
        "brand": "S - 斯达泰克",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266529283290169344.png"
      },
      {
        "id": 1301,
        "brand": "S - 思皓",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266515433337651200.png"
      },
      {
        "id": 1689,
        "brand": "S - SERES赛力斯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266515126838886400.png"
      },
      {
        "id": 1715,
        "brand": "S - 速达汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/136874226136121344.png"
      },
      {
        "id": 1717,
        "brand": "S - SHELBY",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266510555433402368.png"
      },
      {
        "id": 2368,
        "brand": "S - SS DOLPHIN",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266499251863814144.png"
      },
      {
        "id": 119,
        "brand": "T - 泰赫雅特",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266581897256108032.png"
      },
      {
        "id": 120,
        "brand": "T - 特斯拉",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526790387957760.png"
      },
      {
        "id": 166,
        "brand": "T - 腾势",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526725745344512.png"
      },
      {
        "id": 816,
        "brand": "T - 天际汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266522669883588608.png"
      },
      {
        "id": 2414,
        "brand": "T - 坦克",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/153154038492299264.png"
      },
      {
        "id": 121,
        "brand": "W - 五菱",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526047719325696.png"
      },
      {
        "id": 122,
        "brand": "W - 沃尔沃",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526482572181504.png"
      },
      {
        "id": 123,
        "brand": "W - 威麟",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266580966909149184.png"
      },
      {
        "id": 124,
        "brand": "W - 威兹曼",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526346139860992.png"
      },
      {
        "id": 163,
        "brand": "W - 五十铃",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525992501313536.png"
      },
      {
        "id": 544,
        "brand": "W - WEY",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526636029181952.png"
      },
      {
        "id": 571,
        "brand": "W - 蔚来",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526560527515648.png"
      },
      {
        "id": 617,
        "brand": "W - 威马汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526408261697536.png"
      },
      {
        "id": 125,
        "brand": "X - 现代",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525533329883136.png"
      },
      {
        "id": 126,
        "brand": "X - 雪佛兰",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525601269219328.png"
      },
      {
        "id": 127,
        "brand": "X - 雪铁龙",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525144081694720.png"
      },
      {
        "id": 128,
        "brand": "X - 西雅特",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525469287055360.png"
      },
      {
        "id": 130,
        "brand": "X - 新凯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525278530109440.png"
      },
      {
        "id": 638,
        "brand": "X - 新特汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266531413229043712.png"
      },
      {
        "id": 717,
        "brand": "X - 小鹏汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525656357208064.png"
      },
      {
        "id": 751,
        "brand": "X - 星途",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525407798558720.png"
      },
      {
        "id": 2394,
        "brand": "X - 新吉奥",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266495504173498368.png"
      },
      {
        "id": 131,
        "brand": "Y - 一汽",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140524701364518912.png"
      },
      {
        "id": 132,
        "brand": "Y - 英菲尼迪",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525000862990336.png"
      },
      {
        "id": 133,
        "brand": "Y - 野马",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140523456230850560.png"
      },
      {
        "id": 135,
        "brand": "Y - 依维柯",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140525080600903680.png"
      },
      {
        "id": 159,
        "brand": "Y - 英致",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140526275189014528.png"
      },
      {
        "id": 547,
        "brand": "Y - 御捷新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266601441215381504.png"
      },
      {
        "id": 565,
        "brand": "Y - 云度新能源",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140524555385962496.png"
      },
      {
        "id": 568,
        "brand": "Y - 云雀",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266601831029800960.png"
      },
      {
        "id": 643,
        "brand": "Y - 宇通客车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266530321560764416.png"
      },
      {
        "id": 838,
        "brand": "Y - 远程汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/187890262746333184.png"
      },
      {
        "id": 137,
        "brand": "Z - 中华",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140522645308309504.png"
      },
      {
        "id": 138,
        "brand": "Z - 众泰",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140512586436182016.png"
      },
      {
        "id": 139,
        "brand": "Z - 中兴",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140522810651967488.png"
      },
      {
        "id": 140,
        "brand": "Z - 中欧",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266580002118565888.png"
      },
      {
        "id": 141,
        "brand": "Z - 中通客车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266579594058924032.png"
      },
      {
        "id": 168,
        "brand": "Z - 知豆",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/138977218993848320.png"
      },
      {
        "id": 569,
        "brand": "Z - 之诺",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/140522312301543424.png"
      },
      {
        "id": 653,
        "brand": "Z - 宗申",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266527485376593920.png"
      },
      {
        "id": 1076,
        "brand": "Z - 中国重汽VGV",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266515774003216384.png"
      },
      {
        "id": 2671,
        "brand": "Z - 智己汽车",
        "imageUrl": "https://tfcar.oss-cn-hangzhou.aliyuncs.com/266492250869989376.png"
      }
    ],
      seriesList: [],
    };
  },

  created() {
    this.getCarData();

  },
  computed: {
    saleTypeColor() {
      return (saleType) => {
        switch (saleType) {
          case 0:
            return '#3E89E1';
          case 1:
            return '#E0E245';
          case 2:
            return '#C76B32';
          default:
            return '#9A32C7';
        }
      };
    },
  },
  methods: {
    getCarData() {
      axios
          .get("http://118.178.120.11:8080/query", {
          })
          .then((res) => {
            // 列表数据  res.data.data.content
            //请求完毕 给carData赋值 视图自动更新
            this.originalCarData = res.data.carList;
            this.carData = res.data.carList;
          });
    },

    getCarSeries(id) {
      axios
          .get(`http://118.178.120.11:8080/series?brandId=${id}`, {})
          .then((res) => {
            this.seriesList = res.data.seriesList;
          });

    },
    tipToHomePage(SeriesId) {

      // location.href = "赵雨洁-02作业.html?carSeries=" + SeriesId;
      axios
          .get(`http://118.178.120.11:8080/query?carSeries=${SeriesId}`, {})
          .then((res) => {
            this.originalCarData = res.data.carList;
            this.carData = res.data.carList;
            this.drawerButton = false;
            this.drawerSeriesButton = false;
          })

    },
    tipToOtherPage(CarId) {
      // 控制跳转
      this.$router.push(`/Detail/${CarId}`);
    },

    chooseCarsSaleType(id) {
      // 深拷贝carData，确保操作不影响原始数据
      this.carData = this.originalCarData;
      this.newCarData = [];
      for (let i = 0; i < this.carData.length; i++) {
        if (this.carData[i].saleType === id) {
          this.newCarData.push(this.carData[i]);
        }
      }
      // 更新carData为筛选后的结果，但此时原始carData未改变
      this.carData = this.newCarData;
    },
    search() {
      //根据输入框内容 发起搜索请求
      axios
          .get(
              `http://118.178.120.11:8080/search?info=${this.keyword}`,{}
          )
          .then((res) => {
            //拿到匹配的车辆列表数据
            this.originalCarData = res.data.carList;
            this.carData = res.data.carList;
          });
    },


    // 计算属性用于格式化日期
    formattedDate(timestamp) {
      const date = new Date(timestamp); // 注意：JavaScript的Date对象需要的是毫秒时间戳
      const year = date.getFullYear();
      const month = ("0" + (date.getMonth() + 1)).slice(-2); // getMonth返回的月份是从0开始的，所以需要+1
      const day = ("0" + date.getDate()).slice(-2);
      return `${year}年${month}月${day}日`;
    },

    TipToChangPage(page){
      if(page==2){
        this.$router.push(`/NewCar/`);

      }else if(page==3){
        this.$router.push(`/Assess/`)
      }
    }
  },
}

</script>
<style scoped>
* {
  margin: 0;
  padding: 0;
  font-family: Arial, sans-serif;
}
.header div{
  width: 750px;
}
.header img{
  height: 16px;width: 40px;margin: 8px
}
.four-button-group{
  margin-left:25px;
  width:70px;
  height: 20px;
  margin-top: 10px;
  margin-bottom: 10px;
  font-size: 11px;
  background-color: #D5E8FF;
}

hr{
  width:750px;
}
.commonElButton {
  margin-left: 8px;
  width: 50px;
  font-size: 17px;
  height: 30px;
}
.el-input{
  width: 235px;margin:16px;height: 30px;font-size: 17px
}

ul {
  margin: 2px;
  width:400px;
  display: flex;
  justify-content: space-around;
  flex-wrap: wrap;
}

ul li {
  box-sizing: border-box; /* 这里很重要，确保padding和border被包含在width和height内 */
  width: calc(50% - 2px);
  height: 240px;
  list-style: none;
  padding-left: 1.8px ;
  align-items: center; /* 垂直居中 */
}
ul li img{
  width:190px;
  border-radius: 10px;
  height:127px;
}
ul li .saleTypePatten{
  display: inline-block;
  //margin-left: 10px;
  //margin-top:8px;
  font-size:10px;
  padding-bottom:5px;
  padding: 3px;
  text-align: center;
  border-radius: 7px;
  height: 12px;
  width: 54px;
  color: white;
  letter-spacing: 2.6px;
  font-family: Times New Roman;
}
ul li .text-truncate {
  align-items: center; /* 垂直居中 */
  display: -webkit-box; /* 需要此属性才能启用line-clamp */
  -webkit-line-clamp: 2; /* 设置显示的最大行数，这里是2行 */
  -webkit-box-orient: vertical; /* 需要此属性配合line-clamp */
  overflow: hidden; /* 隐藏超出的内容 */
  font-size: 15px;
  padding: 4px;
  margin-left: 3px;
  font-family: 'Microsoft YaHei', sans-serif;
  font-weight: bold;
}
ul li .dateAndMileage{
  font-size: 12px;
  margin-top:4px;
  margin-left:3px;
  color: #606266;
}

h4 {
  font-size: 17px;
  color: #f4001a;
  margin-top:4px;
  margin-left:4px;
}




/*对tabbar进行样式调整*/
@font-face {
  font-family: "iconfont";
  src: url("../assets/iconfont/iconfont.woff2?t=1719390343114") format("woff2"),
  url("../assets/iconfont/iconfont.woff?t=1719390343114") format("woff"),
  url("../assets/iconfont/iconfont.ttf?t=1719390343114") format("truetype");
}

.iconfont {
  font-family: "iconfont" !important;
  font-size: 10px;
  font-style: normal;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

.tabbar {
  display: flex;
  justify-content: space-around;
  background: #f2f2f2;
  padding: 0;
  margin:0;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  /* 定位 */
  position: fixed;
  left: 0;
  bottom: 0;
  width: 400px;
  height:60px;
}

.tabbar li {
  transition: all 0.5s;
  width: 40px;
}

.icon {
  text-align: center;width: 40px;
  font-size:25px;
  margin-bottom:5px;
  margin-top:10px;
}

.title {
  text-align: center;
  width:40px;
  font-size:12px
}

.custom-select > .el-select__wrapper {
  height: 200px;
}
</style>
<style>

.custom-select > .el-select__wrapper {
  height: 30px;
  font-size: 14px;
}

.custom-select > .el-select__input {
  height: 16px;
}
</style>