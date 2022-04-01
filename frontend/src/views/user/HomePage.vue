<template>
  <div class="block">
    <span class="demonstration" style="font-weight:bold; color:#fcd738; font-size: 200%"> Feature Today</span>
    <el-carousel indicator-position="outside">
      <el-carousel-item v-for="item in imgs1" :key="item.id">
        <img :src="item.img" class="image" >
      </el-carousel-item>
    </el-carousel>

    <span class="demonstration" style="font-weight:bold; color:#fcd738; font-size: 200%"> Top Picks</span>
    <el-carousel :interval="4000" type="card" height="200px">
      <el-carousel-item v-for="item in imgs2" :key="item.id">
        <img :src="item.img" class="image" >
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  name: 'img_show',
  data() {
    return {
      imgs1: [
        {
          id: 1, img: require('../../assets/CODA.jpg')
        },
        {
          id: 2, img: require('../../assets/turningRed.jpg')
        },
        {
          id: 3, img: require('../../assets/topGun.jpg')
        },
        {
          id: 4, img: require('../../assets/Spider-Man3.jpg')
        },
        {
          id: 5, img: require('../../assets/500.jpg')
        },
        {
          id: 6, img: require('../../assets/inception.jpg')
        },
        {
          id: 7, img: require('../../assets/Pulp-Fiction1.jpg')
        },
        {
          id: 8, img: require('../../assets/500.jpg')
        },
        {
          id: 9, img: require('../../assets/irobot.jpg')
        },
        {
          id: 10, img: require('../../assets/Morbius.jpg')
        }
      ],
      imgs2 : [
        {
          id: 1, img: require('../../assets/theShawshankRedemption.jpg')
        },
        {
          id: 2, img: require('../../assets/TheGodfather.jpg')
        },
        {
          id: 3, img: require('../../assets/TheDarkKnight.jpg')
        },
        {
          id: 4, img: require('../../assets/12_Angry_Men_(1957_film_poster).jpg')
        },
        {
          id: 5, img: require('../../assets/SchindlerList.jpg')
        },
        {
          id: 6, img: require('../../assets/TheLord.jpg')
        },
        {
          id: 7, img: require('../../assets/Good_the_bad_and_the_ugly_poster.jpg')
        },
        {
          id: 8, img: require('../../assets/forrest.jpg')
        },
        {
          id: 9, img: require('../../assets/irobot.jpg')
        },
        {
          id: 10, img: require('../../assets/forrest.jpg')
        }
      ],
      screenWidth :0
    }
  },
  methods: {
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

<style scoped>
img_show{
  background-color: dimgrey;
  padding-top: 50px;
  width: auto;
}
img {
  align-content: center;
}
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
