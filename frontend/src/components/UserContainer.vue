<template>
  <el-container>
    <div :style="myStyle" id="wrapper">
    </div>
    <el-image
      style="width: 315px; height: 110px"
      :src="require('@/assets/IMDB-CLone-FinalLogo.png')"
      :preview-src-list="require('@/assets/IMDB-CLone-FinalLogo.png')">
    </el-image>
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
            title="IMDb Clone"
            :visible.sync="drawer"
            :direction="direction"
            :before-close="handleClose">
            <span>Welcome to IMDb!</span>
          </el-drawer>
          </span>
        </el-menu-item>

        <el-menu-item>
          <span>
              <el-input placeholder="Search" prefix-icon="el-icon-search" v-model="searchInput"
                        @keydown.enter.native="onSubmit" style=" width: 975px;"/>
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

      <el-row :gutter="20" justify="center" align="middle" type="flex">
        <el-col :span="5">
          <el-button type="danger">Instagram</el-button>
        </el-col>
        <el-col :span="5">
          <el-button type="primary">Twitter</el-button>
        </el-col>
        <el-col :span="5">
          <el-button type="warning">Weibo</el-button>
        </el-col>
        <el-col :span="5">
          <el-button type="success">WeChat</el-button>
        </el-col>
      </el-row>

      <ul>
        <li v-for="(item,index) in arr" :key="item.id" @click="chooseClick(index)">
          <span>
            <img v-if="item.isChoose" src="@/assets/youtube.png" alt />
            <img v-else src="@/assets/Twitter.jpg" alt />
          </span>
          <span>{{item.title}}</span>
        </li>
      </ul>
    </el-footer>
  </el-container>

</template>
<script>
export default {
  data() {
    return {
      myStyle: {
        backgroundColor: "#16a085"
      },
      drawer: false,
      direction: 'ttb',
      searchInput: '',
      username: localStorage.getItem('username'),
      n: 6,
      arr: [
        {id: "1", title: "选项1", isChoose: false},
        {id: "2", title: "选项2", isChoose: false},
        {id: "3", title: "选项3", isChoose: false},
        {id: "4", title: "选项4", isChoose: false},
        {id: "5", title: "礼", isChoose: false},
        {id: "0", title: "其他问题", isChoose: false}
      ],
      imgarr: []
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
    chooseClick(index) {
      //选择问题类型
      for (let i = 0; i < this.arr.length; i++) {
        if (index == i) {
          //选中之后可以进行取消
          // this.arr[index].isChoose = !this.arr[index].isChoose;
          //选中之后不可取消  必须有一个是选中的
          this.arr[index].isChoose = true;
          this.n = index;
        } else {
          this.arr[i].isChoose = false;
        }
      }
    },
  },
}

</script>

<style>
.el-container {
  background-color: #000000;
  color: #333;
  text-align: center;
  width: 100%;
}
.el-header {
  background-color: #000000;
  color: #333;
  text-align: center;
  width: 100%;
}

.el-menu {
  background-color: #000000;
  color: #333;
  text-align: center;
  width: 100%;
}

.el-footer {
  background-color: #000000;
  color: #333;
  text-align: center;
  width: 100%;
}

.el-row {
  background-color: #000000;
  color: #333;
  text-align: center;
  width: 100%;
}

.el-aside {
  background-color: #000000;
  color: #333;
  text-align: center;
  line-height: 200px;
}

.ui {
  list-style: none;
  width: 700px;
  height: 50px;
  margin: 0;
  padding: 0;
}

.el-main {
  background-color: #000000;
  color: #333;
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
.questionType {
  font-size: 0.14rem;
  text-align: left;
  padding: 0.1rem 0.15rem;
  padding-bottom: 0rem;
  box-sizing: border-box;
p {
  letter-spacing: 1px;
  font-size: 0.14rem;
  padding-bottom: 0.15rem;
  box-sizing: border-box;
}
ul {
  width: 100%;
  margin: 0 auto;
  box-sizing: border-box;
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: flex-start;
li {
  width: 50%;
  display: flex;
  align-items: center;
  justify-content: flex-start;
  padding: 0.1rem 0rem;
  padding-left: 0.1rem;
  box-sizing: border-box;
  border-top: 1px solid #ccc;
span:nth-child(1) {
  display: block;
  margin-right: 0.1rem;
img {
  width: 0.12rem;
  height: 0.12rem;
  vertical-align: middle;
}
}
span:nth-child(2) {
  font-size: 0.12rem;
  color: #777;
}
}
}
}



</style>
