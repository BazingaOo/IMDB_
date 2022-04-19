<template>
  <ul class="infinite-list" v-infinite-scroll="load" style="overflow:auto">
    <el-card class="box-card">
    <div slot="header" class="clearfix">
      <span>卡片名称</span>
      <el-button style="float: right; padding: 3px 0" type="text">操作按钮</el-button>
    </div>
    <div v-for="o in 4" :key="o" class="text item">
      {{'列表内容 ' + o }}
    </div>
  </el-card>
    <li v-for="i in count" class="infinite-list-item">{{ i }}</li>
  </ul>
</template>

<script>
export default {
  name: "WishMovie",
  data () {
    return {
      count: 0
    }
  },
  methods: {
    load () {

      this.count = 250
    }
  }
}
</script>
.text {
font-size: 14px;
}

.item {
margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
display: table;
content: "";
}
.clearfix:after {
clear: both
}

.box-card {
width: 480px;
}
<style scoped>

</style>
