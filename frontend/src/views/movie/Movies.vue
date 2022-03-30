<template>
  <div class="block">
    <el-image
      style="width: 259px; height: 375px"
      :src="require('@/assets/Joker.jpg')">
    </el-image>
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


</style>
