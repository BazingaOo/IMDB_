<template>
  <div class="block">
    <el-carousel indicator-position="outside">
      <el-carousel-item v-for="item in assets" :key="item.id">
        <img :src="item.idView" class="image">
      </el-carousel-item>
    </el-carousel>

    <span class="demonstration">Top Picks</span>
    <el-carousel :interval="4000" type="card" height="200px">
      <el-carousel-item v-for="item in assets" :key="item.id">
        <img :src="item.idView" class="image">
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentDate: new Date(),
      assets: [{id:0,idView:require('../assets/uncharted.jpg')},
        {id:1,idView:require('../assets/Beauty-Beast.jpg')},
        {id:2,idView:require('../assets/Captain-Marvel.jpg')}
    ]
    }
      return {
      assets: [{id:0,idView:require('../assets/uncharted.jpg')},
        {id:1,idView:require('../assets/Beauty-Beast.jpg')},
        {id:2,idView:require('../assets/Captain-Marvel.jpg')}
        ]
      }
    },
  methods:{
    setSize:function () {
      // 通过浏览器宽度(图片宽度)计算高度
      this.bannerHeight = 400 / 1920 * this.screenWidth;
    },
  },
  mounted() {
    // 首次加载时,需要调用一次
    this.screenWidth =  window.innerWidth;
    this.setSize();
    // 窗口大小发生改变时,调用一次
    window.onresize = () =>{
      this.screenWidth =  window.innerWidth;
      this.setSize();
    }
  }
}


</script>

<style >
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

.time {
  font-size: 13px;
  color: #999;
}

.bottom {
  margin-top: 13px;
  line-height: 12px;
}

.button {
  padding: 0;
  float: right;
}

.image {
  width: 100%;
  display: block;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}

.clearfix:after {
  clear: both
}
</style>
