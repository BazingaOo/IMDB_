<template>
  <div class="block">
    <el-image
      style="width: 315px; height: 110px"
      :src="require('@/assets/IMDB-CLone-FinalLogo.png')"
      :preview-src-list="require('@/assets/IMDB-CLone-FinalLogo.png')">
    </el-image>
    <el-dropdown>
      <el-button type="primary">
        MENU<i class="el-icon-arrow-down el-icon--right"></i>
      </el-button>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item></el-dropdown-item>
        <el-dropdown-item>Top Movies</el-dropdown-item>
        <el-dropdown-items>Top Cast</el-dropdown-items>
        <el-dropdown-item>Genre</el-dropdown-item>
        <el-dropdown-item>Awards</el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
    <el-input v-model="input" placeholder="Search movie"></el-input>
      <div class="line"></div>
      <el-menu
        :default-active="activeIndex2"
        class="el-menu-demo"
        mode="horizontal"
        @select="handleSelect"
        background-color="#545c64"
        text-color="#fff"
        active-text-color="#ffd04b">
        <el-menu-item index="1">Dashboard</el-menu-item>
        <el-submenu index="2">
          <template slot="title">Settings</template>
          <el-menu-item index="2-1">Personal Information</el-menu-item>
          <el-menu-item index="2-2">选项2</el-menu-item>
          <el-menu-item index="2-3">选项3</el-menu-item>
          <el-submenu index="2-4">
            <template slot="title">选项4</template>
            <el-menu-item index="2-4-1">选项1</el-menu-item>
            <el-menu-item index="2-4-2">选项2</el-menu-item>
            <el-menu-item index="2-4-3">选项3</el-menu-item>
          </el-submenu>
        </el-submenu>
        <el-menu-item index="3" disabled>消息中心</el-menu-item>
        <el-menu-item index="4"><a href="https://www.ele.me" target="_blank">订单管理</a></el-menu-item>
      </el-menu>
    <span class="demonstration">Click to Swtich Moive Poster</span>
    <el-carousel :interval="4000" type="card" height="300px">
      <el-caro0usel-item>
        <div class="item">
          <div class="item__content">
            Text for slide 1
          </div>
          <img class="item__image" src="https://picsum.photos/300?random=1" alt="" >
        </div>0
      </el-caro0usel-item>
      <el-carousel-item>
        <div class="item">
          <div class="item__content">
            Another text for slide 2
          </div>
          <img class="item__image" src="https://picsum.photos/300?random=2" alt="" />
        </div>
      </el-carousel-item>
      <el-carousel-item>
        <div class="item">
          <div class="item__content">
            Yet another for third image
          </div>
          <img class="item__image" src="https://picsum.photos/300?random=3" alt="" />
        </div>
      </el-carousel-item>
    </el-carousel>

    <el-popover
      placement="top-start"
      title="标题"
      width="200"
      trigger="hover"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
      <el-button slot="reference">hover 激活</el-button>
    </el-popover>

    <el-popover
      placement="bottom"
      title="标题"
      width="200"
      trigger="click"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
      <el-button slot="reference">click 激活</el-button>
    </el-popover>

    <el-popover
      ref="popover"
      placement="right"
      title="标题"
      width="200"
      trigger="focus"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。">
    </el-popover>
    <el-button v-popover:popover>focus 激活</el-button>

    <el-popover
      placement="bottom"
      title="标题"
      width="200"
      trigger="manual"
      content="这是一段内容,这是一段内容,这是一段内容,这是一段内容。"
      v-model="visible">
      <el-button slot="reference" @click="visible = !visible">手动激活</el-button>
    </el-popover>

  </div>

</template>

<script>
export default {
  data() {
    return {

      url: "require('@/assets/IMDB-CLone-FinalLogo.png')",
      srcList: [
        "require('@/assets/IMDB-CLone-FinalLogo.png')"
      ],
      input: '',
      activeIndex: '1',
      activeIndex2: '1'
    }
  },
  methods: {
    handleSelect(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClick() {
      alert('button click');
    }
  }
}

</script>

<style scoped>
.el-carousel__item h3 {
  color: #475669;
  font-size: 14px;
  opacity: 0.75;
  line-height: 200px;
  margin: 0;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}

.el-carousel__item:nth-child(2n+1) {
  background-color: #d3dce6;
}

.el-dropdown {
  vertical-align: top;
}
.el-dropdown + .el-dropdown {
  margin-left: 15px;
}
.el-icon-arrow-down {
  font-size: 12px;
}
</style>
