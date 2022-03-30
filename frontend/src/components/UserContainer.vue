<template>
  <el-container>
    <el-header>
      <el-menu class="el-menu-demo" mode="horizontal" active-text-color="#409EFF" :router=true>
        <el-menu-item index="/UserContainer/Homepage">
          <i class="el-icon-s-home"></i>
        </el-menu-item>
        <el-menu-item>
          <span>
          <el-button @click="drawer = true" size="medium" icon="el-icon-menu" round>
            Menu
          </el-button>
          <el-drawer
            title="我是标题"
            :visible.sync="drawer"
            :direction="direction"
            :before-close="handleClose">
            <span>我来啦!</span>
          </el-drawer>
          </span>
        </el-menu-item>

        <el-menu-item>
          <span>
              <el-input placeholder="请输入内容" prefix-icon="el-icon-search" v-model="searchInput"
                        @keydown.enter.native="onSubmit" style=" width: 300px;"/>
          </span>
        </el-menu-item>

        <el-submenu index="" style="float: right" v-if="username!==null||''">
          <template slot="title">{{ username }}</template>
          <el-menu-item index="Userprofile">My profile</el-menu-item>
          <el-menu-item index="2-2">Watch List</el-menu-item>
          <el-menu-item index="/UserLogIn" @click="removeToken">Sign out</el-menu-item>
        </el-submenu>
        <el-menu-item index="/UserLogIn" style="float: right" v-else>
             <span>
            <el-button type="primary" round>LogIn</el-button></span>
        </el-menu-item>
      </el-menu>
    </el-header>


    <el-main>
      <template>
        <router-view/>
      </template>
    </el-main>


    <el-footer>
      <el-divider>below is footer</el-divider>
      <el-row :gutter="20" justify="center" align="middle" type="flex">
        <el-col :span="1">ins</el-col>
        <el-col :span="1">twitter</el-col>
        <el-col :span="1">weibo</el-col>
        <el-col :span="1">wechat</el-col>
      </el-row>
      <el-divider></el-divider>
    </el-footer>
  </el-container>

</template>
<script>
export default {
  data() {
    return {
      drawer: false,
      direction: 'ttb',
      searchInput: '',
      username: localStorage.getItem('username'),
    };
  },
  methods: {
    removeToken() {
      localStorage.clear()
      console.log(localStorage.getItem('username'))
      console.log(localStorage.getItem('token'))
    },
    handleClose(done) {
      done();
    },
    onSubmit() {
      console.log("1111")
      this.$router.push({
        path: 'SearchResult',
        query: {searchInput: this.searchInput}
      });
    },
  }
}
</script>

<style>
.el-header, .el-footer {
  text-align: center;
  width: 100%;
}

.ui {
  list-style: none;
  width: 700px;
  height: 50px;
  margin: 0;
  padding: 0;
}

.el-main {
  text-align: center;
  margin-left: 20%;
  margin-right: 20%;
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
