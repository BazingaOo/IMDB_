<template>
  <div class="block">
    <el-carousel indicator-position="outside">
      <el-carousel-item v-for="item in 4" :key="item">
        <el-image :src= "src" fit="cover"/>

        <el-image :src= "src1" fit="cover"/>
      </el-carousel-item>
    </el-carousel>

    <span class="demonstration">Top Picks</span>
    <el-carousel :interval="4000" type="card" height="200px">
      <el-carousel-item v-for="item in 6" :key="item">
        <h3 class="medium">
          {{<el-image :src= "src2" fit="cover"/>}}
          {{<el-image :src= "src3" fit="cover"/>}}
        </h3>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  data() {
    return {
      currentDate: new Date(),
      src : require(`@/assets/black-panther.jpg`),
      src1 : require(`@/assets/Joker.jpg`),
      src2 : require(`@/assets/uncharted.jpg`),
      src3 : require(`@/assets/Captain-Marvel.jpg`),
    }
  },

  methods: {}
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
