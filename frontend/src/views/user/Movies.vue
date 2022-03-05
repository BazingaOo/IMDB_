<template>
  <div class="block">
    <el-image
      style="width: 259px; height: 375px"
      :src="require('@/assets/joker.jpg')"
      :preview-src-list="require('@/assets/joker.jpg')">
    </el-image>
    <el-tag
      v-for="tag in tags"
      :key="tag.name"
      closable
      :type="tag.type">
      {{ tag.name }}
    </el-tag>
    <el-rate
      v-model="value"
      disabled
      show-score
      text-color="#ff9900"
      score-template="{value}">
    </el-rate>
    <span class="demonstration"></span>
    <el-rate v-model="value1"></el-rate>

    <span class="demonstration">区分颜色</span>
    <el-rate
      v-model="value2"
      :colors="colors">
    </el-rate>
  </div>


</template>

<script>
export default {
  data() {
    return {
      tags: [
        {name: 'Action', type: ''},
        {name: 'Comedy', type: 'success'},
        {name: 'Drama', type: 'info'},
        {name: 'Si-Fi', type: 'warning'},
        {name: 'Romance', type: 'danger'}
      ],
    };

  },
  created() {
    let movie1=this.$route.query.movie
    console.log(movie1)
  },
}
</script>

export default {
data() {
return {
value1: null,
value2: null,
colors: ['#99A9BF', '#F7BA2A', '#FF9900']  // 等同于 { 2: '#99A9BF', 4: { value: '#F7BA2A', excluded: true }, 5: '#FF9900' }
}
}
}
<style scoped>
